/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/workload/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkloadLister helps list Workloads.
type WorkloadLister interface {
	// List lists all Workloads in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Workload, err error)
	// Workloads returns an object that can list and get Workloads.
	Workloads(namespace string) WorkloadNamespaceLister
	WorkloadListerExpansion
}

// workloadLister implements the WorkloadLister interface.
type workloadLister struct {
	indexer cache.Indexer
}

// NewWorkloadLister returns a new WorkloadLister.
func NewWorkloadLister(indexer cache.Indexer) WorkloadLister {
	return &workloadLister{indexer: indexer}
}

// List lists all Workloads in the indexer.
func (s *workloadLister) List(selector labels.Selector) (ret []*v1alpha1.Workload, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Workload))
	})
	return ret, err
}

// Workloads returns an object that can list and get Workloads.
func (s *workloadLister) Workloads(namespace string) WorkloadNamespaceLister {
	return workloadNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkloadNamespaceLister helps list and get Workloads.
type WorkloadNamespaceLister interface {
	// List lists all Workloads in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Workload, err error)
	// Get retrieves the Workload from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Workload, error)
	WorkloadNamespaceListerExpansion
}

// workloadNamespaceLister implements the WorkloadNamespaceLister
// interface.
type workloadNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Workloads in the indexer for a given namespace.
func (s workloadNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Workload, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Workload))
	})
	return ret, err
}

// Get retrieves the Workload from the indexer for a given namespace and name.
func (s workloadNamespaceLister) Get(name string) (*v1alpha1.Workload, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("workload"), name)
	}
	return obj.(*v1alpha1.Workload), nil
}
