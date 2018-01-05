// +build !ignore_autogenerated

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

// This ***REMOVED***le was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthInfo).DeepCopyInto(out.(*AuthInfo))
			return nil
		}, InType: reflect.TypeOf(&AuthInfo{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AuthProviderCon***REMOVED***g).DeepCopyInto(out.(*AuthProviderCon***REMOVED***g))
			return nil
		}, InType: reflect.TypeOf(&AuthProviderCon***REMOVED***g{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Cluster).DeepCopyInto(out.(*Cluster))
			return nil
		}, InType: reflect.TypeOf(&Cluster{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Con***REMOVED***g).DeepCopyInto(out.(*Con***REMOVED***g))
			return nil
		}, InType: reflect.TypeOf(&Con***REMOVED***g{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Context).DeepCopyInto(out.(*Context))
			return nil
		}, InType: reflect.TypeOf(&Context{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Preferences).DeepCopyInto(out.(*Preferences))
			return nil
		}, InType: reflect.TypeOf(&Preferences{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthInfo) DeepCopyInto(out *AuthInfo) {
	*out = *in
	if in.ClientCerti***REMOVED***cateData != nil {
		in, out := &in.ClientCerti***REMOVED***cateData, &out.ClientCerti***REMOVED***cateData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ClientKeyData != nil {
		in, out := &in.ClientKeyData, &out.ClientKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateGroups != nil {
		in, out := &in.ImpersonateGroups, &out.ImpersonateGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateUserExtra != nil {
		in, out := &in.ImpersonateUserExtra, &out.ImpersonateUserExtra
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	if in.AuthProvider != nil {
		in, out := &in.AuthProvider, &out.AuthProvider
		if *in == nil {
			*out = nil
		} ***REMOVED*** {
			*out = new(AuthProviderCon***REMOVED***g)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]runtime.Object, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = val.DeepCopyObject()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthInfo.
func (in *AuthInfo) DeepCopy() *AuthInfo {
	if in == nil {
		return nil
	}
	out := new(AuthInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProviderCon***REMOVED***g) DeepCopyInto(out *AuthProviderCon***REMOVED***g) {
	*out = *in
	if in.Con***REMOVED***g != nil {
		in, out := &in.Con***REMOVED***g, &out.Con***REMOVED***g
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProviderCon***REMOVED***g.
func (in *AuthProviderCon***REMOVED***g) DeepCopy() *AuthProviderCon***REMOVED***g {
	if in == nil {
		return nil
	}
	out := new(AuthProviderCon***REMOVED***g)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	if in.Certi***REMOVED***cateAuthorityData != nil {
		in, out := &in.Certi***REMOVED***cateAuthorityData, &out.Certi***REMOVED***cateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]runtime.Object, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = val.DeepCopyObject()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Con***REMOVED***g) DeepCopyInto(out *Con***REMOVED***g) {
	*out = *in
	in.Preferences.DeepCopyInto(&out.Preferences)
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make(map[string]*Cluster, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = new(Cluster)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	if in.AuthInfos != nil {
		in, out := &in.AuthInfos, &out.AuthInfos
		*out = make(map[string]*AuthInfo, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = new(AuthInfo)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	if in.Contexts != nil {
		in, out := &in.Contexts, &out.Contexts
		*out = make(map[string]*Context, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = new(Context)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]runtime.Object, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = val.DeepCopyObject()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Con***REMOVED***g.
func (in *Con***REMOVED***g) DeepCopy() *Con***REMOVED***g {
	if in == nil {
		return nil
	}
	out := new(Con***REMOVED***g)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Con***REMOVED***g) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} ***REMOVED*** {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]runtime.Object, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = val.DeepCopyObject()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preferences) DeepCopyInto(out *Preferences) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make(map[string]runtime.Object, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} ***REMOVED*** {
				(*out)[key] = val.DeepCopyObject()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preferences.
func (in *Preferences) DeepCopy() *Preferences {
	if in == nil {
		return nil
	}
	out := new(Preferences)
	in.DeepCopyInto(out)
	return out
}
