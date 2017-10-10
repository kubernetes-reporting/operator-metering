/*
Copyright 2016 The Kubernetes Authors.

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

package v1

// This ***REMOVED***le contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_LocalSubjectAccessReview = map[string]string{
	"":       "LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace. Having a namespace scoped resource makes it much easier to grant namespace scoped policy that includes permissions checking.",
	"spec":   "Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.",
	"status": "Status is ***REMOVED***lled in by the server and indicates whether the request is allowed or not",
}

func (LocalSubjectAccessReview) SwaggerDoc() map[string]string {
	return map_LocalSubjectAccessReview
}

var map_NonResourceAttributes = map[string]string{
	"":     "NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface",
	"path": "Path is the URL path of the request",
	"verb": "Verb is the standard HTTP verb",
}

func (NonResourceAttributes) SwaggerDoc() map[string]string {
	return map_NonResourceAttributes
}

var map_NonResourceRule = map[string]string{
	"":                "NonResourceRule holds information that describes a rule for the non-resource",
	"verbs":           "Verb is a list of kubernetes non-resource API verbs, like: get, post, put, delete, patch, head, options.  \"*\" means all.",
	"nonResourceURLs": "NonResourceURLs is a set of partial urls that a user should have access to.  *s are allowed, but only as the full, ***REMOVED***nal step in the path.  \"*\" means all.",
}

func (NonResourceRule) SwaggerDoc() map[string]string {
	return map_NonResourceRule
}

var map_ResourceAttributes = map[string]string{
	"":            "ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface",
	"namespace":   "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces \"\" (empty) is defaulted for LocalSubjectAccessReviews \"\" (empty) is empty for cluster-scoped resources \"\" (empty) means \"all\" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
	"verb":        "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  \"*\" means all.",
	"group":       "Group is the API Group of the Resource.  \"*\" means all.",
	"version":     "Version is the API Version of the Resource.  \"*\" means all.",
	"resource":    "Resource is one of the existing resource types.  \"*\" means all.",
	"subresource": "Subresource is one of the existing resource types.  \"\" means none.",
	"name":        "Name is the name of the resource being requested for a \"get\" or deleted for a \"delete\". \"\" (empty) means all.",
}

func (ResourceAttributes) SwaggerDoc() map[string]string {
	return map_ResourceAttributes
}

var map_ResourceRule = map[string]string{
	"":              "ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn't signi***REMOVED***cant, may contain duplicates, and possibly be incomplete.",
	"verbs":         "Verb is a list of kubernetes resource API verbs, like: get, list, watch, create, update, delete, proxy.  \"*\" means all.",
	"apiGroups":     "APIGroups is the name of the APIGroup that contains the resources.  If multiple API groups are speci***REMOVED***ed, any action requested against one of the enumerated resources in any API group will be allowed.  \"*\" means all.",
	"resources":     "Resources is a list of resources this rule applies to.  ResourceAll represents all resources.  \"*\" means all.",
	"resourceNames": "ResourceNames is an optional white list of names that the rule applies to.  An empty set means that everything is allowed.  \"*\" means all.",
}

func (ResourceRule) SwaggerDoc() map[string]string {
	return map_ResourceRule
}

var map_SelfSubjectAccessReview = map[string]string{
	"":       "SelfSubjectAccessReview checks whether or the current user can perform an action.  Not ***REMOVED***lling in a spec.namespace means \"in all namespaces\".  Self is a special case, because users should always be able to check whether they can perform an action",
	"spec":   "Spec holds information about the request being evaluated.  user and groups must be empty",
	"status": "Status is ***REMOVED***lled in by the server and indicates whether the request is allowed or not",
}

func (SelfSubjectAccessReview) SwaggerDoc() map[string]string {
	return map_SelfSubjectAccessReview
}

var map_SelfSubjectAccessReviewSpec = map[string]string{
	"":                      "SelfSubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set",
	"resourceAttributes":    "ResourceAuthorizationAttributes describes information for a resource access request",
	"nonResourceAttributes": "NonResourceAttributes describes information for a non-resource access request",
}

func (SelfSubjectAccessReviewSpec) SwaggerDoc() map[string]string {
	return map_SelfSubjectAccessReviewSpec
}

var map_SelfSubjectRulesReview = map[string]string{
	"":       "SelfSubjectRulesReview enumerates the set of actions the current user can perform within a namespace. The returned list of actions may be incomplete depending on the server's authorization mode, and any errors experienced during the evaluation. SelfSubjectRulesReview should be used by UIs to show/hide actions, or to quickly let an end user reason about their permissions. It should NOT Be used by external systems to drive authorization decisions as this raises confused deputy, cache lifetime/revocation, and correctness concerns. SubjectAccessReview, and LocalAccessReview are the correct way to defer authorization decisions to the API server.",
	"spec":   "Spec holds information about the request being evaluated.",
	"status": "Status is ***REMOVED***lled in by the server and indicates the set of actions a user can perform.",
}

func (SelfSubjectRulesReview) SwaggerDoc() map[string]string {
	return map_SelfSubjectRulesReview
}

var map_SelfSubjectRulesReviewSpec = map[string]string{
	"namespace": "Namespace to evaluate rules for. Required.",
}

func (SelfSubjectRulesReviewSpec) SwaggerDoc() map[string]string {
	return map_SelfSubjectRulesReviewSpec
}

var map_SubjectAccessReview = map[string]string{
	"":       "SubjectAccessReview checks whether or not a user or group can perform an action.",
	"spec":   "Spec holds information about the request being evaluated",
	"status": "Status is ***REMOVED***lled in by the server and indicates whether the request is allowed or not",
}

func (SubjectAccessReview) SwaggerDoc() map[string]string {
	return map_SubjectAccessReview
}

var map_SubjectAccessReviewSpec = map[string]string{
	"":                      "SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAuthorizationAttributes and NonResourceAuthorizationAttributes must be set",
	"resourceAttributes":    "ResourceAuthorizationAttributes describes information for a resource access request",
	"nonResourceAttributes": "NonResourceAttributes describes information for a non-resource access request",
	"user":                  "User is the user you're testing for. If you specify \"User\" but not \"Groups\", then is it interpreted as \"What if User were not a member of any groups",
	"groups":                "Groups is the groups you're testing for.",
	"extra":                 "Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.",
	"uid":                   "UID information about the requesting user.",
}

func (SubjectAccessReviewSpec) SwaggerDoc() map[string]string {
	return map_SubjectAccessReviewSpec
}

var map_SubjectAccessReviewStatus = map[string]string{
	"":                "SubjectAccessReviewStatus",
	"allowed":         "Allowed is required.  True if the action would be allowed, false otherwise.",
	"reason":          "Reason is optional.  It indicates why a request was allowed or denied.",
	"evaluationError": "EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.",
}

func (SubjectAccessReviewStatus) SwaggerDoc() map[string]string {
	return map_SubjectAccessReviewStatus
}

var map_SubjectRulesReviewStatus = map[string]string{
	"":                 "SubjectRulesReviewStatus contains the result of a rules check. This check can be incomplete depending on the set of authorizers the server is con***REMOVED***gured with and any errors experienced during evaluation. Because authorization rules are additive, if a rule appears in a list it's safe to assume the subject has that permission, even if that list is incomplete.",
	"resourceRules":    "ResourceRules is the list of actions the subject is allowed to perform on resources. The list ordering isn't signi***REMOVED***cant, may contain duplicates, and possibly be incomplete.",
	"nonResourceRules": "NonResourceRules is the list of actions the subject is allowed to perform on non-resources. The list ordering isn't signi***REMOVED***cant, may contain duplicates, and possibly be incomplete.",
	"incomplete":       "Incomplete is true when the rules returned by this call are incomplete. This is most commonly encountered when an authorizer, such as an external authorizer, doesn't support rules evaluation.",
	"evaluationError":  "EvaluationError can appear in combination with Rules. It indicates an error occurred during rule evaluation, such as an authorizer that doesn't support rule evaluation, and that ResourceRules and/or NonResourceRules may be incomplete.",
}

func (SubjectRulesReviewStatus) SwaggerDoc() map[string]string {
	return map_SubjectRulesReviewStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
