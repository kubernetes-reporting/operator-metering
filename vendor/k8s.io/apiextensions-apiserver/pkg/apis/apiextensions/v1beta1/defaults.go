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

package v1beta1

import (
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&CustomResourceDe***REMOVED***nition{}, func(obj interface{}) { SetDefaults_CustomResourceDe***REMOVED***nition(obj.(*CustomResourceDe***REMOVED***nition)) })
	// TODO ***REMOVED***gure out why I can't seem to get my defaulter generated
	// return RegisterDefaults(scheme)
	return nil
}

func SetDefaults_CustomResourceDe***REMOVED***nition(obj *CustomResourceDe***REMOVED***nition) {
	SetDefaults_CustomResourceDe***REMOVED***nitionSpec(&obj.Spec)
}

func SetDefaults_CustomResourceDe***REMOVED***nitionSpec(obj *CustomResourceDe***REMOVED***nitionSpec) {
	if len(obj.Scope) == 0 {
		obj.Scope = NamespaceScoped
	}
	if len(obj.Names.Singular) == 0 {
		obj.Names.Singular = strings.ToLower(obj.Names.Kind)
	}
	if len(obj.Names.ListKind) == 0 && len(obj.Names.Kind) > 0 {
		obj.Names.ListKind = obj.Names.Kind + "List"
	}
}
