/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/evaluation/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EvalJobLister helps list EvalJobs.
type EvalJobLister interface {
	// List lists all EvalJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.EvalJob, err error)
	// EvalJobs returns an object that can list and get EvalJobs.
	EvalJobs(namespace string) EvalJobNamespaceLister
	EvalJobListerExpansion
}

// evalJobLister implements the EvalJobLister interface.
type evalJobLister struct {
	indexer cache.Indexer
}

// NewEvalJobLister returns a new EvalJobLister.
func NewEvalJobLister(indexer cache.Indexer) EvalJobLister {
	return &evalJobLister{indexer: indexer}
}

// List lists all EvalJobs in the indexer.
func (s *evalJobLister) List(selector labels.Selector) (ret []*v1alpha1.EvalJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EvalJob))
	})
	return ret, err
}

// EvalJobs returns an object that can list and get EvalJobs.
func (s *evalJobLister) EvalJobs(namespace string) EvalJobNamespaceLister {
	return evalJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EvalJobNamespaceLister helps list and get EvalJobs.
type EvalJobNamespaceLister interface {
	// List lists all EvalJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.EvalJob, err error)
	// Get retrieves the EvalJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.EvalJob, error)
	EvalJobNamespaceListerExpansion
}

// evalJobNamespaceLister implements the EvalJobNamespaceLister
// interface.
type evalJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EvalJobs in the indexer for a given namespace.
func (s evalJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EvalJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EvalJob))
	})
	return ret, err
}

// Get retrieves the EvalJob from the indexer for a given namespace and name.
func (s evalJobNamespaceLister) Get(name string) (*v1alpha1.EvalJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("evaljob"), name)
	}
	return obj.(*v1alpha1.EvalJob), nil
}
