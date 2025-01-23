/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package configuration

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	appsv1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	appsv1beta1 "github.com/apecloud/kubeblocks/apis/apps/v1beta1"
	cfgcore "github.com/apecloud/kubeblocks/pkg/configuration/core"
	"github.com/apecloud/kubeblocks/pkg/controller/component"
	"github.com/apecloud/kubeblocks/pkg/controller/render"
	testapps "github.com/apecloud/kubeblocks/pkg/testutil/apps"
	testutil "github.com/apecloud/kubeblocks/pkg/testutil/k8s"
)

var _ = Describe("TemplateMergerTest", func() {

	baseConfig := `
[mysqld]
log-bin=master-bin
gtid_mode=OFF
consensus_auto_leader_transfer=ON
max_connections=1000
`
	extendConfig := `
[mysqld]
default_storage_engine=xengine
log_error_verbosity=3
xengine_compression_per_level=kNoCompression:KZSTD:kZSTD
xengine_db_total_write_buffer_size=307
xengine_block_cache_size=307
xengine_row_cache_size=102
max_connections=666
`

	testString := "// this is a test string"

	const (
		baseCMName    = "base-cm"
		updatedCMName = "updated-cm"

		testConfigSpecName = "test-config"
		testClusterName    = "test-cluster"
		testConfigName     = "my.cnf"
		testConfig2Name    = "test.txt"
	)

	var (
		mockClient          *testutil.K8sClientMockHelper
		templateBuilder     render.TemplateRender
		configSpec          appsv1.ComponentConfigSpec
		configConstraintObj *appsv1beta1.ConfigConstraint

		baseCMObject    *corev1.ConfigMap
		updatedCMObject *corev1.ConfigMap
	)

	BeforeEach(func() {
		mockClient = testutil.NewK8sMockClient()
		configConstraintObj = testapps.CheckedCreateCustomizedObj(&testCtx,
			"resources/mysql-config-constraint.yaml",
			&appsv1beta1.ConfigConstraint{})
		baseCMObject = &corev1.ConfigMap{
			Data: map[string]string{
				testConfigName:  baseConfig,
				testConfig2Name: testString,
			},
		}
		updatedCMObject = &corev1.ConfigMap{
			Data: map[string]string{
				testConfigName:  extendConfig,
				testConfig2Name: testString,
			},
		}
		baseCMObject.SetName(baseCMName)
		baseCMObject.SetNamespace("default")
		updatedCMObject.SetName(updatedCMName)
		updatedCMObject.SetNamespace("default")

		configSpec = appsv1.ComponentConfigSpec{
			ComponentTemplateSpec: appsv1.ComponentTemplateSpec{
				Name:        testConfigSpecName,
				TemplateRef: baseCMObject.GetName(),
				Namespace:   "default",
			},
			Keys:                []string{"my.cnf"},
			ConfigConstraintRef: configConstraintObj.GetName(),
		}

		templateBuilder = render.NewTemplateBuilder(&render.ReconcileCtx{
			ResourceCtx: &render.ResourceCtx{
				Context:     ctx,
				Client:      mockClient.Client(),
				ClusterName: testClusterName,
				Namespace:   "default",
			},
			Cluster: &appsv1.Cluster{
				ObjectMeta: metav1.ObjectMeta{
					Name:      testClusterName,
					Namespace: "default",
				},
			},
			SynthesizedComponent: &component.SynthesizedComponent{},
			PodSpec:              nil,
		})

		mockClient.MockGetMethod(testutil.WithGetReturned(testutil.WithConstructSimpleGetResult([]client.Object{
			baseCMObject,
			updatedCMObject,
			configConstraintObj,
		}), testutil.WithAnyTimes()))
	})

	AfterEach(func() {
		mockClient.Finish()
	})

	Context("with patch Merge", func() {
		It("mergerConfigTemplate patch policy", func() {
			importedTemplate := appsv1.ConfigTemplateExtension{
				Namespace: "default",
				// Name:        configSpec.Name,
				TemplateRef: updatedCMObject.GetName(),
				Policy:      appsv1.PatchPolicy,
			}

			tmpCM := baseCMObject.DeepCopy()
			mergedData, err := mergerConfigTemplate(importedTemplate, templateBuilder, configSpec, tmpCM.Data, ctx, mockClient.Client())
			Expect(err).To(Succeed())
			Expect(mergedData).Should(HaveLen(2))

			configReaders, err := cfgcore.LoadRawConfigObject(mergedData, configConstraintObj.Spec.FileFormatConfig, configSpec.Keys)
			Expect(err).Should(Succeed())
			Expect(configReaders).Should(HaveLen(1))
			configObject := configReaders[testConfigName]
			Expect(configObject.Get("gtid_mode")).Should(BeEquivalentTo("OFF"))
			Expect(configObject.Get("consensus_auto_leader_transfer")).Should(BeEquivalentTo("ON"))
			Expect(configObject.Get("default_storage_engine")).Should(BeEquivalentTo("xengine"))
			Expect(configObject.Get("log_error_verbosity")).Should(BeEquivalentTo("3"))
			Expect(configObject.Get("max_connections")).Should(BeEquivalentTo("666"))
			Expect(configObject.Get("xengine_compression_per_level")).Should(BeEquivalentTo("kNoCompression:KZSTD:kZSTD"))
			Expect(configObject.Get("xengine_block_cache_size")).Should(BeEquivalentTo("307"))
		})
	})

	Context("with replace Merge", func() {
		It("test mergerConfigTemplate replace policy", func() {
			importedTemplate := appsv1.ConfigTemplateExtension{
				Namespace:   "default",
				TemplateRef: updatedCMObject.GetName(),
				Policy:      appsv1.ReplacePolicy,
			}

			tmpCM := baseCMObject.DeepCopy()
			mergedData, err := mergerConfigTemplate(importedTemplate, templateBuilder, configSpec, tmpCM.Data, ctx, mockClient.Client())
			Expect(err).Should(Succeed())
			Expect(mergedData).Should(HaveLen(2))
			Expect(reflect.DeepEqual(mergedData, updatedCMObject.Data)).Should(BeTrue())

			configReaders, err := cfgcore.LoadRawConfigObject(mergedData, configConstraintObj.Spec.FileFormatConfig, configSpec.Keys)
			Expect(err).Should(Succeed())
			Expect(configReaders).Should(HaveLen(1))
			configObject := configReaders[testConfigName]
			Expect(configObject.Get("gtid_mode")).Should(BeNil())
			Expect(configObject.Get("consensus_auto_leader_transfer")).Should(BeNil())
			Expect(configObject.Get("default_storage_engine")).Should(BeEquivalentTo("xengine"))
			Expect(configObject.Get("max_connections")).Should(BeEquivalentTo("666"))
		})
	})

	Context("with only add Merge", func() {
		It("test mergerConfigTemplate add policy", func() {
			importedTemplate := appsv1.ConfigTemplateExtension{
				Namespace:   "default",
				TemplateRef: updatedCMObject.GetName(),
				Policy:      appsv1.OnlyAddPolicy,
			}

			tmpCM := baseCMObject.DeepCopy()
			_, err := mergerConfigTemplate(importedTemplate, templateBuilder, configSpec, tmpCM.Data, ctx, mockClient.Client())
			Expect(err).ShouldNot(Succeed())
		})
	})

	Context("with none Merge", func() {
		It("test mergerConfigTemplate none policy", func() {
			importedTemplate := appsv1.ConfigTemplateExtension{
				Namespace:   "default",
				TemplateRef: updatedCMObject.GetName(),
				Policy:      appsv1.NoneMergePolicy,
			}

			tmpCM := baseCMObject.DeepCopy()
			mergedData, err := mergerConfigTemplate(importedTemplate, templateBuilder, configSpec, tmpCM.Data, ctx, mockClient.Client())
			Expect(err).Should(Succeed())
			Expect(reflect.DeepEqual(mergedData, updatedCMObject.Data)).Should(BeTrue())
		})
	})

	Context("failed test", func() {
		It("test mergerConfigTemplate function", func() {
			importedTemplate := appsv1.ConfigTemplateExtension{
				Namespace:   "default",
				TemplateRef: updatedCMObject.GetName(),
				Policy:      "",
			}

			tmpCM := baseCMObject.DeepCopy()
			_, err := mergerConfigTemplate(importedTemplate, templateBuilder, configSpec, tmpCM.Data, ctx, mockClient.Client())
			Expect(err).ShouldNot(Succeed())
		})

		It("not configconstraint", func() {
			importedTemplate := appsv1.ConfigTemplateExtension{
				Namespace:   "default",
				TemplateRef: updatedCMObject.GetName(),
				Policy:      "none",
			}

			tmpCM := baseCMObject.DeepCopy()
			tmpConfigSpec := configSpec.DeepCopy()
			tmpConfigSpec.ConfigConstraintRef = ""
			_, err := mergerConfigTemplate(importedTemplate, templateBuilder, *tmpConfigSpec, tmpCM.Data, ctx, mockClient.Client())
			Expect(err).ShouldNot(Succeed())
		})

		It("not formatter", func() {
			importedTemplate := appsv1.ConfigTemplateExtension{
				Namespace:   "default",
				TemplateRef: updatedCMObject.GetName(),
				Policy:      "none",
			}

			tmpCM := baseCMObject.DeepCopy()
			tmpConfigSpec := configSpec.DeepCopy()
			tmpConfigSpec.ConfigConstraintRef = "not_exist"
			_, err := mergerConfigTemplate(importedTemplate, templateBuilder, *tmpConfigSpec, tmpCM.Data, ctx, mockClient.Client())
			Expect(err).ShouldNot(Succeed())
		})
	})
})
