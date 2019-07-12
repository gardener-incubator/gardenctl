// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/gardener/gardener/pkg/apis/garden/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudProfiles implements CloudProfileInterface
type FakeCloudProfiles struct {
	Fake *FakeGardenV1beta1
}

var cloudprofilesResource = schema.GroupVersionResource{Group: "garden.sapcloud.io", Version: "v1beta1", Resource: "cloudprofiles"}

var cloudprofilesKind = schema.GroupVersionKind{Group: "garden.sapcloud.io", Version: "v1beta1", Kind: "CloudProfile"}

// Get takes name of the cloudProfile, and returns the corresponding cloudProfile object, and an error if there is any.
func (c *FakeCloudProfiles) Get(name string, options v1.GetOptions) (result *v1beta1.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cloudprofilesResource, name), &v1beta1.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudProfile), err
}

// List takes label and field selectors, and returns the list of CloudProfiles that match those selectors.
func (c *FakeCloudProfiles) List(opts v1.ListOptions) (result *v1beta1.CloudProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cloudprofilesResource, cloudprofilesKind, opts), &v1beta1.CloudProfileList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CloudProfileList{ListMeta: obj.(*v1beta1.CloudProfileList).ListMeta}
	for _, item := range obj.(*v1beta1.CloudProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudProfiles.
func (c *FakeCloudProfiles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cloudprofilesResource, opts))
}

// Create takes the representation of a cloudProfile and creates it.  Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *FakeCloudProfiles) Create(cloudProfile *v1beta1.CloudProfile) (result *v1beta1.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cloudprofilesResource, cloudProfile), &v1beta1.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudProfile), err
}

// Update takes the representation of a cloudProfile and updates it. Returns the server's representation of the cloudProfile, and an error, if there is any.
func (c *FakeCloudProfiles) Update(cloudProfile *v1beta1.CloudProfile) (result *v1beta1.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cloudprofilesResource, cloudProfile), &v1beta1.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudProfile), err
}

// Delete takes name of the cloudProfile and deletes it. Returns an error if one occurs.
func (c *FakeCloudProfiles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cloudprofilesResource, name), &v1beta1.CloudProfile{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudProfiles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cloudprofilesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.CloudProfileList{})
	return err
}

// Patch applies the patch and returns the patched cloudProfile.
func (c *FakeCloudProfiles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.CloudProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cloudprofilesResource, name, pt, data, subresources...), &v1beta1.CloudProfile{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudProfile), err
}
