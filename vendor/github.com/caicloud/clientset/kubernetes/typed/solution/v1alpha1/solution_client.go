/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/solution/v1alpha1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type SolutionV1alpha1Interface interface {
	RESTClient() rest.Interface
	SolutionsGetter
}

// SolutionV1alpha1Client is used to interact with features provided by the solution.caicloud.io group.
type SolutionV1alpha1Client struct {
	restClient rest.Interface
}

func (c *SolutionV1alpha1Client) Solutions(namespace string) SolutionInterface {
	return newSolutions(c, namespace)
}

// NewForConfig creates a new SolutionV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*SolutionV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SolutionV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new SolutionV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SolutionV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SolutionV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SolutionV1alpha1Client {
	return &SolutionV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SolutionV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}