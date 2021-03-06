/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	apiregistration "k8s.io/kubernetes/cmd/kube-aggregator/pkg/apis/apiregistration"
	core "k8s.io/kubernetes/pkg/client/testing/core"
)

// FakeAPIServices implements APIServiceInterface
type FakeAPIServices struct {
	Fake *FakeApiregistration
}

var apiservicesResource = schema.GroupVersionResource{Group: "apiregistration.k8s.io", Version: "", Resource: "apiservices"}

func (c *FakeAPIServices) Create(aPIService *apiregistration.APIService) (result *apiregistration.APIService, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootCreateAction(apiservicesResource, aPIService), &apiregistration.APIService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiregistration.APIService), err
}

func (c *FakeAPIServices) Update(aPIService *apiregistration.APIService) (result *apiregistration.APIService, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateAction(apiservicesResource, aPIService), &apiregistration.APIService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiregistration.APIService), err
}

func (c *FakeAPIServices) UpdateStatus(aPIService *apiregistration.APIService) (*apiregistration.APIService, error) {
	obj, err := c.Fake.
		Invokes(core.NewRootUpdateSubresourceAction(apiservicesResource, "status", aPIService), &apiregistration.APIService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiregistration.APIService), err
}

func (c *FakeAPIServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(core.NewRootDeleteAction(apiservicesResource, name), &apiregistration.APIService{})
	return err
}

func (c *FakeAPIServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := core.NewRootDeleteCollectionAction(apiservicesResource, listOptions)

	_, err := c.Fake.Invokes(action, &apiregistration.APIServiceList{})
	return err
}

func (c *FakeAPIServices) Get(name string, options v1.GetOptions) (result *apiregistration.APIService, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootGetAction(apiservicesResource, name), &apiregistration.APIService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiregistration.APIService), err
}

func (c *FakeAPIServices) List(opts v1.ListOptions) (result *apiregistration.APIServiceList, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootListAction(apiservicesResource, opts), &apiregistration.APIServiceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := core.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apiregistration.APIServiceList{}
	for _, item := range obj.(*apiregistration.APIServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIServices.
func (c *FakeAPIServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(core.NewRootWatchAction(apiservicesResource, opts))
}

// Patch applies the patch and returns the patched aPIService.
func (c *FakeAPIServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *apiregistration.APIService, err error) {
	obj, err := c.Fake.
		Invokes(core.NewRootPatchSubresourceAction(apiservicesResource, name, data, subresources...), &apiregistration.APIService{})
	if obj == nil {
		return nil, err
	}
	return obj.(*apiregistration.APIService), err
}
