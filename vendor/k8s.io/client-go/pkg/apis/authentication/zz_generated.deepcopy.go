// +build !ignore_autogenerated

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

// This ***REMOVED***le was autogenerated by deepcopy-gen. Do not edit it manually!

package authentication

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authentication_TokenReview, InType: reflect.TypeOf(&TokenReview{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authentication_TokenReviewSpec, InType: reflect.TypeOf(&TokenReviewSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authentication_TokenReviewStatus, InType: reflect.TypeOf(&TokenReviewStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_authentication_UserInfo, InType: reflect.TypeOf(&UserInfo{})},
	)
}

// DeepCopy_authentication_TokenReview is an autogenerated deepcopy function.
func DeepCopy_authentication_TokenReview(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TokenReview)
		out := out.(*TokenReview)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} ***REMOVED*** {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_authentication_TokenReviewStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_authentication_TokenReviewSpec is an autogenerated deepcopy function.
func DeepCopy_authentication_TokenReviewSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TokenReviewSpec)
		out := out.(*TokenReviewSpec)
		*out = *in
		return nil
	}
}

// DeepCopy_authentication_TokenReviewStatus is an autogenerated deepcopy function.
func DeepCopy_authentication_TokenReviewStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TokenReviewStatus)
		out := out.(*TokenReviewStatus)
		*out = *in
		if err := DeepCopy_authentication_UserInfo(&in.User, &out.User, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_authentication_UserInfo is an autogenerated deepcopy function.
func DeepCopy_authentication_UserInfo(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*UserInfo)
		out := out.(*UserInfo)
		*out = *in
		if in.Groups != nil {
			in, out := &in.Groups, &out.Groups
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		if in.Extra != nil {
			in, out := &in.Extra, &out.Extra
			*out = make(map[string]ExtraValue)
			for key, val := range *in {
				if newVal, err := c.DeepCopy(&val); err != nil {
					return err
				} ***REMOVED*** {
					(*out)[key] = *newVal.(*ExtraValue)
				}
			}
		}
		return nil
	}
}
