/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/servicemesh/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IstioLister helps list Istios.
type IstioLister interface {
	// List lists all Istios in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Istio, err error)
	// Get retrieves the Istio from the index for a given name.
	Get(name string) (*v1alpha1.Istio, error)
	IstioListerExpansion
}

// istioLister implements the IstioLister interface.
type istioLister struct {
	indexer cache.Indexer
}

// NewIstioLister returns a new IstioLister.
func NewIstioLister(indexer cache.Indexer) IstioLister {
	return &istioLister{indexer: indexer}
}

// List lists all Istios in the indexer.
func (s *istioLister) List(selector labels.Selector) (ret []*v1alpha1.Istio, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Istio))
	})
	return ret, err
}

// Get retrieves the Istio from the index for a given name.
func (s *istioLister) Get(name string) (*v1alpha1.Istio, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("istio"), name)
	}
	return obj.(*v1alpha1.Istio), nil
}