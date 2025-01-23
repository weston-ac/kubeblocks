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

	v1 "github.com/apecloud/kubeblocks/apis/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComponentVersions implements ComponentVersionInterface
type FakeComponentVersions struct {
	Fake *FakeAppsV1
}

var componentversionsResource = v1.SchemeGroupVersion.WithResource("componentversions")

var componentversionsKind = v1.SchemeGroupVersion.WithKind("ComponentVersion")

// Get takes name of the componentVersion, and returns the corresponding componentVersion object, and an error if there is any.
func (c *FakeComponentVersions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ComponentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(componentversionsResource, name), &v1.ComponentVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentVersion), err
}

// List takes label and field selectors, and returns the list of ComponentVersions that match those selectors.
func (c *FakeComponentVersions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ComponentVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(componentversionsResource, componentversionsKind, opts), &v1.ComponentVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ComponentVersionList{ListMeta: obj.(*v1.ComponentVersionList).ListMeta}
	for _, item := range obj.(*v1.ComponentVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested componentVersions.
func (c *FakeComponentVersions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(componentversionsResource, opts))
}

// Create takes the representation of a componentVersion and creates it.  Returns the server's representation of the componentVersion, and an error, if there is any.
func (c *FakeComponentVersions) Create(ctx context.Context, componentVersion *v1.ComponentVersion, opts metav1.CreateOptions) (result *v1.ComponentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(componentversionsResource, componentVersion), &v1.ComponentVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentVersion), err
}

// Update takes the representation of a componentVersion and updates it. Returns the server's representation of the componentVersion, and an error, if there is any.
func (c *FakeComponentVersions) Update(ctx context.Context, componentVersion *v1.ComponentVersion, opts metav1.UpdateOptions) (result *v1.ComponentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(componentversionsResource, componentVersion), &v1.ComponentVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComponentVersions) UpdateStatus(ctx context.Context, componentVersion *v1.ComponentVersion, opts metav1.UpdateOptions) (*v1.ComponentVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(componentversionsResource, "status", componentVersion), &v1.ComponentVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentVersion), err
}

// Delete takes name of the componentVersion and deletes it. Returns an error if one occurs.
func (c *FakeComponentVersions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(componentversionsResource, name, opts), &v1.ComponentVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComponentVersions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(componentversionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ComponentVersionList{})
	return err
}

// Patch applies the patch and returns the patched componentVersion.
func (c *FakeComponentVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ComponentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(componentversionsResource, name, pt, data, subresources...), &v1.ComponentVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ComponentVersion), err
}
