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

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/schema/generated.proto
// DO NOT EDIT!

/*
	Package schema is a generated protocol buffer package.

	It is generated from these ***REMOVED***les:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/schema/generated.proto

	It has these top-level messages:
*/
package schema

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated ***REMOVED***le
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/schema/generated.proto", ***REMOVED***leDescriptorGenerated)
}

var ***REMOVED***leDescriptorGenerated = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xce, 0xaf, 0x4e, 0x04, 0x31,
	0x10, 0xc7, 0xf1, 0xd6, 0x20, 0x90, 0xc8, 0x13, 0x23, 0x51, 0xd0, 0x11, 0x18, 0x34, 0x2f, 0x80,
	0xc7, 0x75, 0xf7, 0x86, 0x6e, 0x53, 0xfa, 0x27, 0xed, 0x94, 0x04, 0xc7, 0x23, 0xf0, 0x58, 0x27,
	0x4f, 0xae, 0x64, 0xcb, 0x8b, 0x90, 0xb4, 0x2b, 0x08, 0xc9, 0xb9, 0xfe, 0xd2, 0x7c, 0x26, 0xdf,
	0xeb, 0x67, 0xf7, 0x58, 0x94, 0x8d, 0xe8, 0xea, 0x44, 0x39, 0x10, 0x53, 0xc1, 0x77, 0x0a, 0xc7,
	0x98, 0x71, 0xff, 0xd0, 0xc9, 0x7a, 0x3d, 0x2f, 0x36, 0x50, 0xfe, 0xc0, 0xe4, 0x0c, 0xe6, 0x1a,
	0xd8, 0x7a, 0xc2, 0x32, 0x2f, 0xe4, 0x35, 0x1a, 0x0a, 0x94, 0x35, 0xd3, 0x51, 0xa5, 0x1c, 0x39,
	0xde, 0xdc, 0x0e, 0xa7, 0xfe, 0x3a, 0x95, 0x9c, 0x51, 0xbb, 0x53, 0xc3, 0x1d, 0xee, 0x8d, 0xe5,
	0xa5, 0x4e, 0x6a, 0x8e, 0x1e, 0x4d, 0x34, 0x11, 0x3b, 0x9f, 0xea, 0x6b, 0x5f, 0x7d, 0xf4, 0xd7,
	0x38, 0x7b, 0x78, 0xb8, 0x94, 0x53, 0xd9, 0xbe, 0xa1, 0x0d, 0x5c, 0x38, 0xff, 0x6f, 0x79, 0xba,
	0x3b, 0x6d, 0x20, 0xce, 0x1b, 0x88, 0x75, 0x03, 0xf1, 0xd9, 0x40, 0x9e, 0x1a, 0xc8, 0x73, 0x03,
	0xb9, 0x36, 0x90, 0xdf, 0x0d, 0xe4, 0xd7, 0x0f, 0x88, 0x97, 0xab, 0x51, 0xf4, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xfd, 0x59, 0x57, 0x93, 0x0b, 0x01, 0x00, 0x00,
}
