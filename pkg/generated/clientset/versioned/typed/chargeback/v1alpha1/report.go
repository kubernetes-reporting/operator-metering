package v1alpha1

import (
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	scheme "github.com/coreos-inc/kube-chargeback/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReportsGetter has a method to return a ReportInterface.
// A group's client should implement this interface.
type ReportsGetter interface {
	Reports(namespace string) ReportInterface
}

// ReportInterface has methods to work with Report resources.
type ReportInterface interface {
	Create(*v1alpha1.Report) (*v1alpha1.Report, error)
	Update(*v1alpha1.Report) (*v1alpha1.Report, error)
	UpdateStatus(*v1alpha1.Report) (*v1alpha1.Report, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Report, error)
	List(opts v1.ListOptions) (*v1alpha1.ReportList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Report, err error)
	ReportExpansion
}

// reports implements ReportInterface
type reports struct {
	client rest.Interface
	ns     string
}

// newReports returns a Reports
func newReports(c *ChargebackV1alpha1Client, namespace string) *reports {
	return &reports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the report, and returns the corresponding report object, and an error if there is any.
func (c *reports) Get(name string, options v1.GetOptions) (result *v1alpha1.Report, err error) {
	result = &v1alpha1.Report{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and ***REMOVED***eld selectors, and returns the list of Reports that match those selectors.
func (c *reports) List(opts v1.ListOptions) (result *v1alpha1.ReportList, err error) {
	result = &v1alpha1.ReportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested reports.
func (c *reports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("reports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a report and creates it.  Returns the server's representation of the report, and an error, if there is any.
func (c *reports) Create(report *v1alpha1.Report) (result *v1alpha1.Report, err error) {
	result = &v1alpha1.Report{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("reports").
		Body(report).
		Do().
		Into(result)
	return
}

// Update takes the representation of a report and updates it. Returns the server's representation of the report, and an error, if there is any.
func (c *reports) Update(report *v1alpha1.Report) (result *v1alpha1.Report, err error) {
	result = &v1alpha1.Report{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reports").
		Name(report.Name).
		Body(report).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *reports) UpdateStatus(report *v1alpha1.Report) (result *v1alpha1.Report, err error) {
	result = &v1alpha1.Report{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reports").
		Name(report.Name).
		SubResource("status").
		Body(report).
		Do().
		Into(result)
	return
}

// Delete takes name of the report and deletes it. Returns an error if one occurs.
func (c *reports) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reports").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *reports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reports").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched report.
func (c *reports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Report, err error) {
	result = &v1alpha1.Report{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("reports").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
