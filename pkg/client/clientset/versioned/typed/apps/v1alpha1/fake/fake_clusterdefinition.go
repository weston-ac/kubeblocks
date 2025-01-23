/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterDefinitions implements ClusterDefinitionInterface
type FakeClusterDefinitions struct {
	Fake *FakeAppsV1alpha1
}

var clusterdefinitionsResource = v1alpha1.SchemeGroupVersion.WithResource("clusterdefinitions")

var clusterdefinitionsKind = v1alpha1.SchemeGroupVersion.WithKind("ClusterDefinition")

// Get takes name of the clusterDefinition, and returns the corresponding clusterDefinition object, and an error if there is any.
func (c *FakeClusterDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterdefinitionsResource, name), &v1alpha1.ClusterDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDefinition), err
}

// List takes label and field selectors, and returns the list of ClusterDefinitions that match those selectors.
func (c *FakeClusterDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterdefinitionsResource, clusterdefinitionsKind, opts), &v1alpha1.ClusterDefinitionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterDefinitionList{ListMeta: obj.(*v1alpha1.ClusterDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterDefinitions.
func (c *FakeClusterDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterdefinitionsResource, opts))
}

// Create takes the representation of a clusterDefinition and creates it.  Returns the server's representation of the clusterDefinition, and an error, if there is any.
func (c *FakeClusterDefinitions) Create(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.CreateOptions) (result *v1alpha1.ClusterDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterdefinitionsResource, clusterDefinition), &v1alpha1.ClusterDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDefinition), err
}

// Update takes the representation of a clusterDefinition and updates it. Returns the server's representation of the clusterDefinition, and an error, if there is any.
func (c *FakeClusterDefinitions) Update(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.UpdateOptions) (result *v1alpha1.ClusterDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterdefinitionsResource, clusterDefinition), &v1alpha1.ClusterDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterDefinitions) UpdateStatus(ctx context.Context, clusterDefinition *v1alpha1.ClusterDefinition, opts v1.UpdateOptions) (*v1alpha1.ClusterDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterdefinitionsResource, "status", clusterDefinition), &v1alpha1.ClusterDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDefinition), err
}

// Delete takes name of the clusterDefinition and deletes it. Returns an error if one occurs.
func (c *FakeClusterDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterdefinitionsResource, name, opts), &v1alpha1.ClusterDefinition{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterdefinitionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched clusterDefinition.
func (c *FakeClusterDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterdefinitionsResource, name, pt, data, subresources...), &v1alpha1.ClusterDefinition{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDefinition), err
}
