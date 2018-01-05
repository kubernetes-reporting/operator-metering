/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this ***REMOVED***le except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the speci***REMOVED***c language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	v1beta1 "k8s.io/api/extensions/v1beta1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type ExtensionsV1beta1Interface interface {
	RESTClient() rest.Interface
	DaemonSetsGetter
	DeploymentsGetter
	IngressesGetter
	PodSecurityPoliciesGetter
	ReplicaSetsGetter
	ScalesGetter
	ThirdPartyResourcesGetter
}

// ExtensionsV1beta1Client is used to interact with features provided by the extensions group.
type ExtensionsV1beta1Client struct {
	restClient rest.Interface
}

func (c *ExtensionsV1beta1Client) DaemonSets(namespace string) DaemonSetInterface {
	return newDaemonSets(c, namespace)
}

func (c *ExtensionsV1beta1Client) Deployments(namespace string) DeploymentInterface {
	return newDeployments(c, namespace)
}

func (c *ExtensionsV1beta1Client) Ingresses(namespace string) IngressInterface {
	return newIngresses(c, namespace)
}

func (c *ExtensionsV1beta1Client) PodSecurityPolicies() PodSecurityPolicyInterface {
	return newPodSecurityPolicies(c)
}

func (c *ExtensionsV1beta1Client) ReplicaSets(namespace string) ReplicaSetInterface {
	return newReplicaSets(c, namespace)
}

func (c *ExtensionsV1beta1Client) Scales(namespace string) ScaleInterface {
	return newScales(c, namespace)
}

func (c *ExtensionsV1beta1Client) ThirdPartyResources() ThirdPartyResourceInterface {
	return newThirdPartyResources(c)
}

// NewForCon***REMOVED***g creates a new ExtensionsV1beta1Client for the given con***REMOVED***g.
func NewForCon***REMOVED***g(c *rest.Con***REMOVED***g) (*ExtensionsV1beta1Client, error) {
	con***REMOVED***g := *c
	if err := setCon***REMOVED***gDefaults(&con***REMOVED***g); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&con***REMOVED***g)
	if err != nil {
		return nil, err
	}
	return &ExtensionsV1beta1Client{client}, nil
}

// NewForCon***REMOVED***gOrDie creates a new ExtensionsV1beta1Client for the given con***REMOVED***g and
// panics if there is an error in the con***REMOVED***g.
func NewForCon***REMOVED***gOrDie(c *rest.Con***REMOVED***g) *ExtensionsV1beta1Client {
	client, err := NewForCon***REMOVED***g(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ExtensionsV1beta1Client for the given RESTClient.
func New(c rest.Interface) *ExtensionsV1beta1Client {
	return &ExtensionsV1beta1Client{c}
}

func setCon***REMOVED***gDefaults(con***REMOVED***g *rest.Con***REMOVED***g) error {
	gv := v1beta1.SchemeGroupVersion
	con***REMOVED***g.GroupVersion = &gv
	con***REMOVED***g.APIPath = "/apis"
	con***REMOVED***g.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if con***REMOVED***g.UserAgent == "" {
		con***REMOVED***g.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ExtensionsV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
