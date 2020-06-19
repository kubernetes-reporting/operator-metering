// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/kube-reporting/metering-operator/pkg/apis/metering/v1"
	scheme "github.com/kube-reporting/metering-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HiveTablesGetter has a method to return a HiveTableInterface.
// A group's client should implement this interface.
type HiveTablesGetter interface {
	HiveTables(namespace string) HiveTableInterface
}

// HiveTableInterface has methods to work with HiveTable resources.
type HiveTableInterface interface {
	Create(ctx context.Context, hiveTable *v1.HiveTable, opts metav1.CreateOptions) (*v1.HiveTable, error)
	Update(ctx context.Context, hiveTable *v1.HiveTable, opts metav1.UpdateOptions) (*v1.HiveTable, error)
	UpdateStatus(ctx context.Context, hiveTable *v1.HiveTable, opts metav1.UpdateOptions) (*v1.HiveTable, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.HiveTable, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HiveTableList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HiveTable, err error)
	HiveTableExpansion
}

// hiveTables implements HiveTableInterface
type hiveTables struct {
	client rest.Interface
	ns     string
}

// newHiveTables returns a HiveTables
func newHiveTables(c *MeteringV1Client, namespace string) *hiveTables {
	return &hiveTables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hiveTable, and returns the corresponding hiveTable object, and an error if there is any.
func (c *hiveTables) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.HiveTable, err error) {
	result = &v1.HiveTable{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hivetables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HiveTables that match those selectors.
func (c *hiveTables) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HiveTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.HiveTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hivetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hiveTables.
func (c *hiveTables) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hivetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a hiveTable and creates it.  Returns the server's representation of the hiveTable, and an error, if there is any.
func (c *hiveTables) Create(ctx context.Context, hiveTable *v1.HiveTable, opts metav1.CreateOptions) (result *v1.HiveTable, err error) {
	result = &v1.HiveTable{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hivetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hiveTable).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a hiveTable and updates it. Returns the server's representation of the hiveTable, and an error, if there is any.
func (c *hiveTables) Update(ctx context.Context, hiveTable *v1.HiveTable, opts metav1.UpdateOptions) (result *v1.HiveTable, err error) {
	result = &v1.HiveTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hivetables").
		Name(hiveTable.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hiveTable).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hiveTables) UpdateStatus(ctx context.Context, hiveTable *v1.HiveTable, opts metav1.UpdateOptions) (result *v1.HiveTable, err error) {
	result = &v1.HiveTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hivetables").
		Name(hiveTable.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hiveTable).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the hiveTable and deletes it. Returns an error if one occurs.
func (c *hiveTables) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hivetables").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hiveTables) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hivetables").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched hiveTable.
func (c *hiveTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HiveTable, err error) {
	result = &v1.HiveTable{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hivetables").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
