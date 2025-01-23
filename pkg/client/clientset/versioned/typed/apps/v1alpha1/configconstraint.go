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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/apecloud/kubeblocks/apis/apps/v1alpha1"
	scheme "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigConstraintsGetter has a method to return a ConfigConstraintInterface.
// A group's client should implement this interface.
type ConfigConstraintsGetter interface {
	ConfigConstraints() ConfigConstraintInterface
}

// ConfigConstraintInterface has methods to work with ConfigConstraint resources.
type ConfigConstraintInterface interface {
	Create(ctx context.Context, configConstraint *v1alpha1.ConfigConstraint, opts v1.CreateOptions) (*v1alpha1.ConfigConstraint, error)
	Update(ctx context.Context, configConstraint *v1alpha1.ConfigConstraint, opts v1.UpdateOptions) (*v1alpha1.ConfigConstraint, error)
	UpdateStatus(ctx context.Context, configConstraint *v1alpha1.ConfigConstraint, opts v1.UpdateOptions) (*v1alpha1.ConfigConstraint, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ConfigConstraint, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ConfigConstraintList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigConstraint, err error)
	ConfigConstraintExpansion
}

// configConstraints implements ConfigConstraintInterface
type configConstraints struct {
	client rest.Interface
}

// newConfigConstraints returns a ConfigConstraints
func newConfigConstraints(c *AppsV1alpha1Client) *configConstraints {
	return &configConstraints{
		client: c.RESTClient(),
	}
}

// Get takes name of the configConstraint, and returns the corresponding configConstraint object, and an error if there is any.
func (c *configConstraints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigConstraint, err error) {
	result = &v1alpha1.ConfigConstraint{}
	err = c.client.Get().
		Resource("configconstraints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigConstraints that match those selectors.
func (c *configConstraints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigConstraintList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigConstraintList{}
	err = c.client.Get().
		Resource("configconstraints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configConstraints.
func (c *configConstraints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("configconstraints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a configConstraint and creates it.  Returns the server's representation of the configConstraint, and an error, if there is any.
func (c *configConstraints) Create(ctx context.Context, configConstraint *v1alpha1.ConfigConstraint, opts v1.CreateOptions) (result *v1alpha1.ConfigConstraint, err error) {
	result = &v1alpha1.ConfigConstraint{}
	err = c.client.Post().
		Resource("configconstraints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configConstraint).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a configConstraint and updates it. Returns the server's representation of the configConstraint, and an error, if there is any.
func (c *configConstraints) Update(ctx context.Context, configConstraint *v1alpha1.ConfigConstraint, opts v1.UpdateOptions) (result *v1alpha1.ConfigConstraint, err error) {
	result = &v1alpha1.ConfigConstraint{}
	err = c.client.Put().
		Resource("configconstraints").
		Name(configConstraint.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configConstraint).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *configConstraints) UpdateStatus(ctx context.Context, configConstraint *v1alpha1.ConfigConstraint, opts v1.UpdateOptions) (result *v1alpha1.ConfigConstraint, err error) {
	result = &v1alpha1.ConfigConstraint{}
	err = c.client.Put().
		Resource("configconstraints").
		Name(configConstraint.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(configConstraint).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the configConstraint and deletes it. Returns an error if one occurs.
func (c *configConstraints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("configconstraints").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configConstraints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("configconstraints").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched configConstraint.
func (c *configConstraints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigConstraint, err error) {
	result = &v1alpha1.ConfigConstraint{}
	err = c.client.Patch(pt).
		Resource("configconstraints").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
