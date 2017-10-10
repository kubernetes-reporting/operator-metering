package v1alpha1

import (
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	scheme "github.com/coreos-inc/kube-chargeback/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReportPrometheusQueriesGetter has a method to return a ReportPrometheusQueryInterface.
// A group's client should implement this interface.
type ReportPrometheusQueriesGetter interface {
	ReportPrometheusQueries(namespace string) ReportPrometheusQueryInterface
}

// ReportPrometheusQueryInterface has methods to work with ReportPrometheusQuery resources.
type ReportPrometheusQueryInterface interface {
	Create(*v1alpha1.ReportPrometheusQuery) (*v1alpha1.ReportPrometheusQuery, error)
	Update(*v1alpha1.ReportPrometheusQuery) (*v1alpha1.ReportPrometheusQuery, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ReportPrometheusQuery, error)
	List(opts v1.ListOptions) (*v1alpha1.ReportPrometheusQueryList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReportPrometheusQuery, err error)
	ReportPrometheusQueryExpansion
}

// reportPrometheusQueries implements ReportPrometheusQueryInterface
type reportPrometheusQueries struct {
	client rest.Interface
	ns     string
}

// newReportPrometheusQueries returns a ReportPrometheusQueries
func newReportPrometheusQueries(c *ChargebackV1alpha1Client, namespace string) *reportPrometheusQueries {
	return &reportPrometheusQueries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the reportPrometheusQuery, and returns the corresponding reportPrometheusQuery object, and an error if there is any.
func (c *reportPrometheusQueries) Get(name string, options v1.GetOptions) (result *v1alpha1.ReportPrometheusQuery, err error) {
	result = &v1alpha1.ReportPrometheusQuery{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and ***REMOVED***eld selectors, and returns the list of ReportPrometheusQueries that match those selectors.
func (c *reportPrometheusQueries) List(opts v1.ListOptions) (result *v1alpha1.ReportPrometheusQueryList, err error) {
	result = &v1alpha1.ReportPrometheusQueryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested reportPrometheusQueries.
func (c *reportPrometheusQueries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a reportPrometheusQuery and creates it.  Returns the server's representation of the reportPrometheusQuery, and an error, if there is any.
func (c *reportPrometheusQueries) Create(reportPrometheusQuery *v1alpha1.ReportPrometheusQuery) (result *v1alpha1.ReportPrometheusQuery, err error) {
	result = &v1alpha1.ReportPrometheusQuery{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		Body(reportPrometheusQuery).
		Do().
		Into(result)
	return
}

// Update takes the representation of a reportPrometheusQuery and updates it. Returns the server's representation of the reportPrometheusQuery, and an error, if there is any.
func (c *reportPrometheusQueries) Update(reportPrometheusQuery *v1alpha1.ReportPrometheusQuery) (result *v1alpha1.ReportPrometheusQuery, err error) {
	result = &v1alpha1.ReportPrometheusQuery{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		Name(reportPrometheusQuery.Name).
		Body(reportPrometheusQuery).
		Do().
		Into(result)
	return
}

// Delete takes name of the reportPrometheusQuery and deletes it. Returns an error if one occurs.
func (c *reportPrometheusQueries) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *reportPrometheusQueries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched reportPrometheusQuery.
func (c *reportPrometheusQueries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReportPrometheusQuery, err error) {
	result = &v1alpha1.ReportPrometheusQuery{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("reportprometheusqueries").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
