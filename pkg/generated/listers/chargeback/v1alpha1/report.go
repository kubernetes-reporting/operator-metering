// Code generated by lister-gen. DO NOT EDIT.

// This ***REMOVED***le was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/operator-framework/operator-metering/pkg/apis/chargeback/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReportLister helps list Reports.
type ReportLister interface {
	// List lists all Reports in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Report, err error)
	// Reports returns an object that can list and get Reports.
	Reports(namespace string) ReportNamespaceLister
	ReportListerExpansion
}

// reportLister implements the ReportLister interface.
type reportLister struct {
	indexer cache.Indexer
}

// NewReportLister returns a new ReportLister.
func NewReportLister(indexer cache.Indexer) ReportLister {
	return &reportLister{indexer: indexer}
}

// List lists all Reports in the indexer.
func (s *reportLister) List(selector labels.Selector) (ret []*v1alpha1.Report, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Report))
	})
	return ret, err
}

// Reports returns an object that can list and get Reports.
func (s *reportLister) Reports(namespace string) ReportNamespaceLister {
	return reportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReportNamespaceLister helps list and get Reports.
type ReportNamespaceLister interface {
	// List lists all Reports in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Report, err error)
	// Get retrieves the Report from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Report, error)
	ReportNamespaceListerExpansion
}

// reportNamespaceLister implements the ReportNamespaceLister
// interface.
type reportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Reports in the indexer for a given namespace.
func (s reportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Report, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Report))
	})
	return ret, err
}

// Get retrieves the Report from the indexer for a given namespace and name.
func (s reportNamespaceLister) Get(name string) (*v1alpha1.Report, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("report"), name)
	}
	return obj.(*v1alpha1.Report), nil
}
