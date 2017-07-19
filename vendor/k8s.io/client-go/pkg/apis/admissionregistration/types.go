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

package admissionregistration

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient=true
// +nonNamespaced=true

// InitializerCon***REMOVED***guration describes the con***REMOVED***guration of initializers.
type InitializerCon***REMOVED***guration struct {
	metav1.TypeMeta
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta

	// Initializers is a list of resources and their default initializers
	// Order-sensitive.
	// When merging multiple InitializerCon***REMOVED***gurations, we sort the initializers
	// from different InitializerCon***REMOVED***gurations by the name of the
	// InitializerCon***REMOVED***gurations; the order of the initializers from the same
	// InitializerCon***REMOVED***guration is preserved.
	// +optional
	Initializers []Initializer
}

// InitializerCon***REMOVED***gurationList is a list of InitializerCon***REMOVED***guration.
type InitializerCon***REMOVED***gurationList struct {
	metav1.TypeMeta
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta

	// List of InitializerCon***REMOVED***guration.
	Items []InitializerCon***REMOVED***guration
}

// Initializer describes the name and the failure policy of an initializer, and
// what resources it applies to.
type Initializer struct {
	// Name is the identi***REMOVED***er of the initializer. It will be added to the
	// object that needs to be initialized.
	// Name should be fully quali***REMOVED***ed, e.g., alwayspullimages.kubernetes.io, where
	// "alwayspullimages" is the name of the webhook, and kubernetes.io is the name
	// of the organization.
	// Required
	Name string

	// Rules describes what resources/subresources the initializer cares about.
	// The initializer cares about an operation if it matches _any_ Rule.
	// Rule.Resources must not include subresources.
	Rules []Rule

	// FailurePolicy de***REMOVED***nes what happens if the responsible initializer controller
	// fails to takes action. Allowed values are Ignore, or Fail. If "Ignore" is
	// set, initializer is removed from the initializers list of an object if
	// the timeout is reached; If "Fail" is set, admissionregistration returns timeout error
	// if the timeout is reached.
	FailurePolicy *FailurePolicyType
}

// Rule is a tuple of APIGroups, APIVersion, and Resources.It is recommended
// to make sure that all the tuple expansions are valid.
type Rule struct {
	// APIGroups is the API groups the resources belong to. '*' is all groups.
	// If '*' is present, the length of the slice must be one.
	// Required.
	APIGroups []string

	// APIVersions is the API versions the resources belong to. '*' is all versions.
	// If '*' is present, the length of the slice must be one.
	// Required.
	APIVersions []string

	// Resources is a list of resources this rule applies to.
	//
	// For example:
	// 'pods' means pods.
	// 'pods/log' means the log subresource of pods.
	// '*' means all resources, but not subresources.
	// 'pods/*' means all subresources of pods.
	// '*/scale' means all scale subresources.
	// '*/*' means all resources and their subresources.
	//
	// If wildcard is present, the validation rule will ensure resources do not
	// overlap with each other.
	//
	// Depending on the enclosing object, subresources might not be allowed.
	// Required.
	Resources []string
}

type FailurePolicyType string

const (
	// Ignore means the initilizer is removed from the initializers list of an
	// object if the initializer is timed out.
	Ignore FailurePolicyType = "Ignore"
	// For 1.7, only "Ignore" is allowed. "Fail" will be allowed when the
	// extensible admission feature is beta.
	Fail FailurePolicyType = "Fail"
)

// +genclient=true
// +nonNamespaced=true

// ExternalAdmissionHookCon***REMOVED***guration describes the con***REMOVED***guration of initializers.
type ExternalAdmissionHookCon***REMOVED***guration struct {
	metav1.TypeMeta
	// Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata.
	// +optional
	metav1.ObjectMeta
	// ExternalAdmissionHooks is a list of external admission webhooks and the
	// affected resources and operations.
	// +optional
	ExternalAdmissionHooks []ExternalAdmissionHook
}

// ExternalAdmissionHookCon***REMOVED***gurationList is a list of ExternalAdmissionHookCon***REMOVED***guration.
type ExternalAdmissionHookCon***REMOVED***gurationList struct {
	metav1.TypeMeta
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	// +optional
	metav1.ListMeta
	// List of ExternalAdmissionHookCon***REMOVED***guration.
	Items []ExternalAdmissionHookCon***REMOVED***guration
}

// ExternalAdmissionHook describes an external admission webhook and the
// resources and operations it applies to.
type ExternalAdmissionHook struct {
	// The name of the external admission webhook.
	// Name should be fully quali***REMOVED***ed, e.g., imagepolicy.kubernetes.io, where
	// "imagepolicy" is the name of the webhook, and kubernetes.io is the name
	// of the organization.
	// Required.
	Name string

	// ClientCon***REMOVED***g de***REMOVED***nes how to communicate with the hook.
	// Required
	ClientCon***REMOVED***g AdmissionHookClientCon***REMOVED***g

	// Rules describes what operations on what resources/subresources the webhook cares about.
	// The webhook cares about an operation if it matches _any_ Rule.
	Rules []RuleWithOperations

	// FailurePolicy de***REMOVED***nes how unrecognized errors from the admission endpoint are handled -
	// allowed values are Ignore or Fail. Defaults to Ignore.
	// +optional
	FailurePolicy *FailurePolicyType
}

// RuleWithOperations is a tuple of Operations and Resources. It is recommended to make
// sure that all the tuple expansions are valid.
type RuleWithOperations struct {
	// Operations is the operations the admission hook cares about - CREATE, UPDATE, or *
	// for all operations.
	// If '*' is present, the length of the slice must be one.
	// Required.
	Operations []OperationType
	// Rule is embedded, it describes other criteria of the rule, like
	// APIGroups, APIVersions, Resources, etc.
	Rule
}

type OperationType string

// The constants should be kept in sync with those de***REMOVED***ned in k8s.io/kubernetes/pkg/admission/interface.go.
const (
	OperationAll OperationType = "*"
	Create       OperationType = "CREATE"
	Update       OperationType = "UPDATE"
	Delete       OperationType = "DELETE"
	Connect      OperationType = "CONNECT"
)

// AdmissionHookClientCon***REMOVED***g contains the information to make a TLS
// connection with the webhook
type AdmissionHookClientCon***REMOVED***g struct {
	// Service is a reference to the service for this webhook. If there is only
	// one port open for the service, that port will be used. If there are multiple
	// ports open, port 443 will be used if it is open, otherwise it is an error.
	// Required
	Service ServiceReference
	// CABundle is a PEM encoded CA bundle which will be used to validate webhook's server certi***REMOVED***cate.
	// Required
	CABundle []byte
}

// ServiceReference holds a reference to Service.legacy.k8s.io
type ServiceReference struct {
	// Namespace is the namespace of the service
	// Required
	Namespace string
	// Name is the name of the service
	// Required
	Name string
}
