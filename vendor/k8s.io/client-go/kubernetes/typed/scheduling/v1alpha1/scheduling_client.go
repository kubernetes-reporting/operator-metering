/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "k8s.io/api/scheduling/v1alpha1"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

type SchedulingV1alpha1Interface interface {
	RESTClient() rest.Interface
	PriorityClassesGetter
}

// SchedulingV1alpha1Client is used to interact with features provided by the scheduling.k8s.io group.
type SchedulingV1alpha1Client struct {
	restClient rest.Interface
}

func (c *SchedulingV1alpha1Client) PriorityClasses() PriorityClassInterface {
	return newPriorityClasses(c)
}

// NewForCon***REMOVED***g creates a new SchedulingV1alpha1Client for the given con***REMOVED***g.
func NewForCon***REMOVED***g(c *rest.Con***REMOVED***g) (*SchedulingV1alpha1Client, error) {
	con***REMOVED***g := *c
	if err := setCon***REMOVED***gDefaults(&con***REMOVED***g); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&con***REMOVED***g)
	if err != nil {
		return nil, err
	}
	return &SchedulingV1alpha1Client{client}, nil
}

// NewForCon***REMOVED***gOrDie creates a new SchedulingV1alpha1Client for the given con***REMOVED***g and
// panics if there is an error in the con***REMOVED***g.
func NewForCon***REMOVED***gOrDie(c *rest.Con***REMOVED***g) *SchedulingV1alpha1Client {
	client, err := NewForCon***REMOVED***g(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SchedulingV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SchedulingV1alpha1Client {
	return &SchedulingV1alpha1Client{c}
}

func setCon***REMOVED***gDefaults(con***REMOVED***g *rest.Con***REMOVED***g) error {
	gv := v1alpha1.SchemeGroupVersion
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
func (c *SchedulingV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
