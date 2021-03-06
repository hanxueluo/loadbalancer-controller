/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/dataset/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DatasetLister helps list Datasets.
type DatasetLister interface {
	// List lists all Datasets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Dataset, err error)
	// Get retrieves the Dataset from the index for a given name.
	Get(name string) (*v1alpha1.Dataset, error)
	DatasetListerExpansion
}

// datasetLister implements the DatasetLister interface.
type datasetLister struct {
	indexer cache.Indexer
}

// NewDatasetLister returns a new DatasetLister.
func NewDatasetLister(indexer cache.Indexer) DatasetLister {
	return &datasetLister{indexer: indexer}
}

// List lists all Datasets in the indexer.
func (s *datasetLister) List(selector labels.Selector) (ret []*v1alpha1.Dataset, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Dataset))
	})
	return ret, err
}

// Get retrieves the Dataset from the index for a given name.
func (s *datasetLister) Get(name string) (*v1alpha1.Dataset, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dataset"), name)
	}
	return obj.(*v1alpha1.Dataset), nil
}
