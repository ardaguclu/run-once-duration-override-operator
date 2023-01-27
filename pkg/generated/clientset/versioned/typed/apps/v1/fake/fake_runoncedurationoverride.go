/*
Copyright 2023 Red Hat, Inc.

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

	appsv1 "github.com/openshift/run-once-duration-override-operator/pkg/apis/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRunOnceDurationOverrides implements RunOnceDurationOverrideInterface
type FakeRunOnceDurationOverrides struct {
	Fake *FakeAppsV1
}

var runoncedurationoverridesResource = schema.GroupVersionResource{Group: "apps.openshift.io", Version: "v1", Resource: "runoncedurationoverrides"}

var runoncedurationoverridesKind = schema.GroupVersionKind{Group: "apps.openshift.io", Version: "v1", Kind: "RunOnceDurationOverride"}

// Get takes name of the runOnceDurationOverride, and returns the corresponding runOnceDurationOverride object, and an error if there is any.
func (c *FakeRunOnceDurationOverrides) Get(ctx context.Context, name string, options v1.GetOptions) (result *appsv1.RunOnceDurationOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(runoncedurationoverridesResource, name), &appsv1.RunOnceDurationOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.RunOnceDurationOverride), err
}

// List takes label and field selectors, and returns the list of RunOnceDurationOverrides that match those selectors.
func (c *FakeRunOnceDurationOverrides) List(ctx context.Context, opts v1.ListOptions) (result *appsv1.RunOnceDurationOverrideList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(runoncedurationoverridesResource, runoncedurationoverridesKind, opts), &appsv1.RunOnceDurationOverrideList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appsv1.RunOnceDurationOverrideList{ListMeta: obj.(*appsv1.RunOnceDurationOverrideList).ListMeta}
	for _, item := range obj.(*appsv1.RunOnceDurationOverrideList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested runOnceDurationOverrides.
func (c *FakeRunOnceDurationOverrides) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(runoncedurationoverridesResource, opts))
}

// Create takes the representation of a runOnceDurationOverride and creates it.  Returns the server's representation of the runOnceDurationOverride, and an error, if there is any.
func (c *FakeRunOnceDurationOverrides) Create(ctx context.Context, runOnceDurationOverride *appsv1.RunOnceDurationOverride, opts v1.CreateOptions) (result *appsv1.RunOnceDurationOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(runoncedurationoverridesResource, runOnceDurationOverride), &appsv1.RunOnceDurationOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.RunOnceDurationOverride), err
}

// Update takes the representation of a runOnceDurationOverride and updates it. Returns the server's representation of the runOnceDurationOverride, and an error, if there is any.
func (c *FakeRunOnceDurationOverrides) Update(ctx context.Context, runOnceDurationOverride *appsv1.RunOnceDurationOverride, opts v1.UpdateOptions) (result *appsv1.RunOnceDurationOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(runoncedurationoverridesResource, runOnceDurationOverride), &appsv1.RunOnceDurationOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.RunOnceDurationOverride), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRunOnceDurationOverrides) UpdateStatus(ctx context.Context, runOnceDurationOverride *appsv1.RunOnceDurationOverride, opts v1.UpdateOptions) (*appsv1.RunOnceDurationOverride, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(runoncedurationoverridesResource, "status", runOnceDurationOverride), &appsv1.RunOnceDurationOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.RunOnceDurationOverride), err
}

// Delete takes name of the runOnceDurationOverride and deletes it. Returns an error if one occurs.
func (c *FakeRunOnceDurationOverrides) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(runoncedurationoverridesResource, name), &appsv1.RunOnceDurationOverride{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRunOnceDurationOverrides) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(runoncedurationoverridesResource, listOpts)

	_, err := c.Fake.Invokes(action, &appsv1.RunOnceDurationOverrideList{})
	return err
}

// Patch applies the patch and returns the patched runOnceDurationOverride.
func (c *FakeRunOnceDurationOverrides) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appsv1.RunOnceDurationOverride, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(runoncedurationoverridesResource, name, pt, data, subresources...), &appsv1.RunOnceDurationOverride{})
	if obj == nil {
		return nil, err
	}
	return obj.(*appsv1.RunOnceDurationOverride), err
}