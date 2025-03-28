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

package instanceset

import (
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/golang/mock/gomock"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	workloads "github.com/apecloud/kubeblocks/apis/workloads/v1"
	"github.com/apecloud/kubeblocks/pkg/constant"
	"github.com/apecloud/kubeblocks/pkg/controller/builder"
	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
)

var _ = Describe("pod role label event handler test", func() {
	Context("Handle function", func() {
		It("should work well", func() {
			cli := k8sMock
			reqCtx := intctrlutil.RequestCtx{
				Ctx: ctx,
				Log: logger,
			}
			pod := builder.NewPodBuilder(namespace, getPodName(name, 0)).SetUID(uid).GetObject()
			pod.ResourceVersion = "1"
			objectRef := corev1.ObjectReference{
				APIVersion: "v1",
				Kind:       "Pod",
				Namespace:  pod.Namespace,
				Name:       pod.Name,
				UID:        pod.UID,
				FieldPath:  lorryEventFieldPath,
			}
			role := workloads.ReplicaRole{
				Name:                 "leader",
				ParticipatesInQuorum: true,
				UpdatePriority:       5,
			}

			By("build an expected message")
			message := fmt.Sprintf("Readiness probe failed: error: health rpc failed: rpc error: code = Unknown desc = {\"event\":\"Success\",\"originalRole\":\"\",\"role\":\"%s\"}", role.Name)
			event := builder.NewEventBuilder(namespace, "foo").
				SetInvolvedObject(objectRef).
				SetReason(checkRoleOperation).
				SetMessage(message).
				GetObject()

			handler := &PodRoleEventHandler{}
			k8sMock.EXPECT().
				Get(gomock.Any(), gomock.Any(), &corev1.Pod{}, gomock.Any()).
				DoAndReturn(func(_ context.Context, objKey client.ObjectKey, p *corev1.Pod, _ ...client.GetOption) error {
					p.Namespace = objKey.Namespace
					p.Name = objKey.Name
					p.UID = pod.UID
					p.Labels = map[string]string{
						constant.AppInstanceLabelKey: name,
						WorkloadsInstanceLabelKey:    name,
					}
					return nil
				}).Times(1)
			k8sMock.EXPECT().
				Get(gomock.Any(), gomock.Any(), &workloads.InstanceSet{}, gomock.Any()).
				DoAndReturn(func(_ context.Context, objKey client.ObjectKey, its *workloads.InstanceSet, _ ...client.GetOption) error {
					its.Namespace = objKey.Namespace
					its.Name = objKey.Name
					its.Spec.Roles = []workloads.ReplicaRole{role}
					return nil
				}).Times(1)
			k8sMock.EXPECT().
				Update(gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, pd *corev1.Pod, _ ...client.UpdateOption) error {
					Expect(pd).ShouldNot(BeNil())
					Expect(pd.Labels).ShouldNot(BeNil())
					Expect(pd.Labels[RoleLabelKey]).Should(Equal(role.Name))
					return nil
				}).Times(1)
			k8sMock.EXPECT().
				Patch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, evt *corev1.Event, patch client.Patch, _ ...client.PatchOption) error {
					Expect(evt).ShouldNot(BeNil())
					Expect(evt.Annotations).ShouldNot(BeNil())
					Expect(evt.Annotations[roleChangedAnnotKey]).Should(Equal(fmt.Sprintf("count-%d", evt.Count)))
					return nil
				}).Times(1)
			Expect(handler.Handle(cli, reqCtx, nil, event)).Should(Succeed())

			By("build an unexpected message")
			message = "unexpected message"
			event = builder.NewEventBuilder(namespace, "foo").
				SetInvolvedObject(objectRef).
				SetMessage(message).
				SetReason(checkRoleOperation).
				GetObject()
			k8sMock.EXPECT().
				Patch(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, evt *corev1.Event, patch client.Patch, _ ...client.PatchOption) error {
					Expect(evt).ShouldNot(BeNil())
					Expect(evt.Annotations).ShouldNot(BeNil())
					Expect(evt.Annotations[roleChangedAnnotKey]).Should(Equal(fmt.Sprintf("count-%d", evt.Count)))
					return nil
				}).Times(1)
			Expect(handler.Handle(cli, reqCtx, nil, event)).Should(Succeed())

			By("read a stale pod")
			message = fmt.Sprintf("Readiness probe failed: error: health rpc failed: rpc error: code = Unknown desc = {\"event\":\"Success\",\"originalRole\":\"\",\"role\":\"%s\"}", role.Name)
			event = builder.NewEventBuilder(namespace, "foo").
				SetInvolvedObject(objectRef).
				SetReason(checkRoleOperation).
				SetMessage(message).
				GetObject()

			k8sMock.EXPECT().
				Get(gomock.Any(), gomock.Any(), &corev1.Pod{}, gomock.Any()).
				DoAndReturn(func(_ context.Context, objKey client.ObjectKey, p *corev1.Pod, _ ...client.GetOption) error {
					p.Namespace = objKey.Namespace
					p.ResourceVersion = "0"
					p.Name = objKey.Name
					p.UID = pod.UID
					p.Labels = map[string]string{
						constant.AppInstanceLabelKey: name,
						WorkloadsInstanceLabelKey:    name,
					}
					return nil
				}).Times(1)
			k8sMock.EXPECT().
				Get(gomock.Any(), gomock.Any(), &workloads.InstanceSet{}, gomock.Any()).
				DoAndReturn(func(_ context.Context, objKey client.ObjectKey, its *workloads.InstanceSet, _ ...client.GetOption) error {
					its.Namespace = objKey.Namespace
					its.Name = objKey.Name
					its.Spec.Roles = []workloads.ReplicaRole{role}
					return nil
				}).Times(1)
			updateErr := fmt.Errorf("the object has been modified; please apply your changes to the latest version and try again")
			k8sMock.EXPECT().
				Update(gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, pd *corev1.Pod, _ ...client.UpdateOption) error {
					Expect(pd).ShouldNot(BeNil())
					Expect(pd.Labels).ShouldNot(BeNil())
					Expect(pd.Labels[RoleLabelKey]).Should(Equal(role.Name))
					if pd.ResourceVersion <= pod.ResourceVersion {
						return updateErr
					}
					return nil
				}).Times(1)
			Expect(handler.Handle(cli, reqCtx, nil, event)).Should(Equal(updateErr))
		})
	})

	Context("parseProbeEventMessage function", func() {
		It("should work well", func() {
			reqCtx := intctrlutil.RequestCtx{
				Ctx: ctx,
				Log: logf.FromContext(ctx).WithValues("pod-role-event-handler", namespace),
			}

			By("build an well formatted message")
			roleName := "leader"
			message := fmt.Sprintf("Readiness probe failed: error: health rpc failed: rpc error: code = Unknown desc = {\"event\":\"Success\",\"originalRole\":\"\",\"role\":\"%s\"}", roleName)
			event := builder.NewEventBuilder(namespace, "foo").
				SetMessage(message).
				GetObject()
			msg := parseProbeEventMessage(reqCtx, event)
			Expect(msg).ShouldNot(BeNil())
			Expect(msg.Role).Should(Equal(roleName))

			By("build an error formatted message")
			message = "Readiness probe failed: error: health rpc failed: rpc error: code = Unknown desc = {\"event\":}"
			event.Message = message
			Expect(parseProbeEventMessage(reqCtx, event)).Should(BeNil())
		})
	})
})
