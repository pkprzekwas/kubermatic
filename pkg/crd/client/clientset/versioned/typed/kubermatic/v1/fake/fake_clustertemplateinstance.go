// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/crd/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterTemplateInstances implements ClusterTemplateInstanceInterface
type FakeClusterTemplateInstances struct {
	Fake *FakeKubermaticV1
}

var clustertemplateinstancesResource = schema.GroupVersionResource{Group: "kubermatic.k8s.io", Version: "v1", Resource: "clustertemplateinstances"}

var clustertemplateinstancesKind = schema.GroupVersionKind{Group: "kubermatic.k8s.io", Version: "v1", Kind: "ClusterTemplateInstance"}

// Get takes name of the clusterTemplateInstance, and returns the corresponding clusterTemplateInstance object, and an error if there is any.
func (c *FakeClusterTemplateInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.ClusterTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustertemplateinstancesResource, name), &kubermaticv1.ClusterTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplateInstance), err
}

// List takes label and field selectors, and returns the list of ClusterTemplateInstances that match those selectors.
func (c *FakeClusterTemplateInstances) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.ClusterTemplateInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustertemplateinstancesResource, clustertemplateinstancesKind, opts), &kubermaticv1.ClusterTemplateInstanceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.ClusterTemplateInstanceList{ListMeta: obj.(*kubermaticv1.ClusterTemplateInstanceList).ListMeta}
	for _, item := range obj.(*kubermaticv1.ClusterTemplateInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTemplateInstances.
func (c *FakeClusterTemplateInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustertemplateinstancesResource, opts))
}

// Create takes the representation of a clusterTemplateInstance and creates it.  Returns the server's representation of the clusterTemplateInstance, and an error, if there is any.
func (c *FakeClusterTemplateInstances) Create(ctx context.Context, clusterTemplateInstance *kubermaticv1.ClusterTemplateInstance, opts v1.CreateOptions) (result *kubermaticv1.ClusterTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustertemplateinstancesResource, clusterTemplateInstance), &kubermaticv1.ClusterTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplateInstance), err
}

// Update takes the representation of a clusterTemplateInstance and updates it. Returns the server's representation of the clusterTemplateInstance, and an error, if there is any.
func (c *FakeClusterTemplateInstances) Update(ctx context.Context, clusterTemplateInstance *kubermaticv1.ClusterTemplateInstance, opts v1.UpdateOptions) (result *kubermaticv1.ClusterTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustertemplateinstancesResource, clusterTemplateInstance), &kubermaticv1.ClusterTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplateInstance), err
}

// Delete takes name of the clusterTemplateInstance and deletes it. Returns an error if one occurs.
func (c *FakeClusterTemplateInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustertemplateinstancesResource, name, opts), &kubermaticv1.ClusterTemplateInstance{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTemplateInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustertemplateinstancesResource, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.ClusterTemplateInstanceList{})
	return err
}

// Patch applies the patch and returns the patched clusterTemplateInstance.
func (c *FakeClusterTemplateInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.ClusterTemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustertemplateinstancesResource, name, pt, data, subresources...), &kubermaticv1.ClusterTemplateInstance{})
	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplateInstance), err
}
