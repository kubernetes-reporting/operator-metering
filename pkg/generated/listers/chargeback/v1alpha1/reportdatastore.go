// This ***REMOVED***le was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReportDataStoreLister helps list ReportDataStores.
type ReportDataStoreLister interface {
	// List lists all ReportDataStores in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ReportDataStore, err error)
	// ReportDataStores returns an object that can list and get ReportDataStores.
	ReportDataStores(namespace string) ReportDataStoreNamespaceLister
	ReportDataStoreListerExpansion
}

// reportDataStoreLister implements the ReportDataStoreLister interface.
type reportDataStoreLister struct {
	indexer cache.Indexer
}

// NewReportDataStoreLister returns a new ReportDataStoreLister.
func NewReportDataStoreLister(indexer cache.Indexer) ReportDataStoreLister {
	return &reportDataStoreLister{indexer: indexer}
}

// List lists all ReportDataStores in the indexer.
func (s *reportDataStoreLister) List(selector labels.Selector) (ret []*v1alpha1.ReportDataStore, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReportDataStore))
	})
	return ret, err
}

// ReportDataStores returns an object that can list and get ReportDataStores.
func (s *reportDataStoreLister) ReportDataStores(namespace string) ReportDataStoreNamespaceLister {
	return reportDataStoreNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReportDataStoreNamespaceLister helps list and get ReportDataStores.
type ReportDataStoreNamespaceLister interface {
	// List lists all ReportDataStores in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ReportDataStore, err error)
	// Get retrieves the ReportDataStore from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ReportDataStore, error)
	ReportDataStoreNamespaceListerExpansion
}

// reportDataStoreNamespaceLister implements the ReportDataStoreNamespaceLister
// interface.
type reportDataStoreNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReportDataStores in the indexer for a given namespace.
func (s reportDataStoreNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ReportDataStore, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReportDataStore))
	})
	return ret, err
}

// Get retrieves the ReportDataStore from the indexer for a given namespace and name.
func (s reportDataStoreNamespaceLister) Get(name string) (*v1alpha1.ReportDataStore, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("reportdatastore"), name)
	}
	return obj.(*v1alpha1.ReportDataStore), nil
}
