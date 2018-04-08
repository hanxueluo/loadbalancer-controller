/*
Copyright 2018 caicloud authors. All rights reserved.
*/

package v1beta1

import (
	"github.com/caicloud/clientset/kubernetes/scheme"
	v1beta1 "github.com/caicloud/clientset/pkg/apis/apiextensions/v1beta1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type ApiextensionsV1beta1Interface interface {
	RESTClient() rest.Interface
	CustomResourceDefinitionsGetter
}

// ApiextensionsV1beta1Client is used to interact with features provided by the apiextensions.k8s.io group.
type ApiextensionsV1beta1Client struct {
	restClient rest.Interface
}

func (c *ApiextensionsV1beta1Client) CustomResourceDefinitions() CustomResourceDefinitionInterface {
	return newCustomResourceDefinitions(c)
}

// NewForConfig creates a new ApiextensionsV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*ApiextensionsV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ApiextensionsV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new ApiextensionsV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ApiextensionsV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ApiextensionsV1beta1Client for the given RESTClient.
func New(c rest.Interface) *ApiextensionsV1beta1Client {
	return &ApiextensionsV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
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
func (c *ApiextensionsV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
