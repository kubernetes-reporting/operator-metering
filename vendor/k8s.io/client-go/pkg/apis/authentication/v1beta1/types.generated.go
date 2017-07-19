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

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package v1beta1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkg2_types "k8s.io/apimachinery/pkg/types"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, ***REMOVED***le, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate ***REMOVED***le: %v",
			5, codec1978.GenVersion, ***REMOVED***le)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg1_v1.TypeMeta
		var v1 pkg2_types.UID
		var v2 time.Time
		_, _, _ = v0, v1, v2
	}
}

func (x *TokenReview) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} ***REMOVED*** {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} ***REMOVED*** if z.HasExtensions() && z.EncExt(x) {
		} ***REMOVED*** {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [5]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			yyq2[4] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(5)
			} ***REMOVED*** {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} ***REMOVED*** {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} ***REMOVED*** {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} ***REMOVED*** {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} ***REMOVED*** {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ObjectMeta
					yym11 := z.EncBinary()
					_ = yym11
					if false {
					} ***REMOVED*** if z.HasExtensions() && z.EncExt(yy10) {
					} ***REMOVED*** {
						z.EncFallback(yy10)
					}
				} ***REMOVED*** {
					r.EncodeNil()
				}
			} ***REMOVED*** {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy12 := &x.ObjectMeta
					yym13 := z.EncBinary()
					_ = yym13
					if false {
					} ***REMOVED*** if z.HasExtensions() && z.EncExt(yy12) {
					} ***REMOVED*** {
						z.EncFallback(yy12)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy15 := &x.Spec
				yy15.CodecEncodeSelf(e)
			} ***REMOVED*** {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("spec"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy17 := &x.Spec
				yy17.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[4] {
					yy20 := &x.Status
					yy20.CodecEncodeSelf(e)
				} ***REMOVED*** {
					r.EncodeNil()
				}
			} ***REMOVED*** {
				if yyq2[4] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("status"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy22 := &x.Status
					yy22.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TokenReview) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} ***REMOVED*** if z.HasExtensions() && z.DecExt(x) {
	} ***REMOVED*** {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} ***REMOVED*** if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} ***REMOVED*** {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReview) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} ***REMOVED*** {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} ***REMOVED*** {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} ***REMOVED*** {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} ***REMOVED*** {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} ***REMOVED*** {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ObjectMeta = pkg1_v1.ObjectMeta{}
			} ***REMOVED*** {
				yyv8 := &x.ObjectMeta
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} ***REMOVED*** if z.HasExtensions() && z.DecExt(yyv8) {
				} ***REMOVED*** {
					z.DecFallback(yyv8, false)
				}
			}
		case "spec":
			if r.TryDecodeAsNil() {
				x.Spec = TokenReviewSpec{}
			} ***REMOVED*** {
				yyv10 := &x.Spec
				yyv10.CodecDecodeSelf(d)
			}
		case "status":
			if r.TryDecodeAsNil() {
				x.Status = TokenReviewStatus{}
			} ***REMOVED*** {
				yyv11 := &x.Status
				yyv11.CodecDecodeSelf(d)
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReview) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} ***REMOVED*** {
		yyv13 := &x.Kind
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} ***REMOVED*** {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} ***REMOVED*** {
		yyv15 := &x.APIVersion
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} ***REMOVED*** {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ObjectMeta = pkg1_v1.ObjectMeta{}
	} ***REMOVED*** {
		yyv17 := &x.ObjectMeta
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} ***REMOVED*** if z.HasExtensions() && z.DecExt(yyv17) {
		} ***REMOVED*** {
			z.DecFallback(yyv17, false)
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Spec = TokenReviewSpec{}
	} ***REMOVED*** {
		yyv19 := &x.Spec
		yyv19.CodecDecodeSelf(d)
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Status = TokenReviewStatus{}
	} ***REMOVED*** {
		yyv20 := &x.Status
		yyv20.CodecDecodeSelf(d)
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} ***REMOVED*** {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TokenReviewSpec) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} ***REMOVED*** {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} ***REMOVED*** if z.HasExtensions() && z.EncExt(x) {
		} ***REMOVED*** {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [1]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Token != ""
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(1)
			} ***REMOVED*** {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Token))
					}
				} ***REMOVED*** {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} ***REMOVED*** {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("token"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Token))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TokenReviewSpec) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} ***REMOVED*** if z.HasExtensions() && z.DecExt(x) {
	} ***REMOVED*** {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} ***REMOVED*** if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} ***REMOVED*** {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReviewSpec) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} ***REMOVED*** {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "token":
			if r.TryDecodeAsNil() {
				x.Token = ""
			} ***REMOVED*** {
				yyv4 := &x.Token
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} ***REMOVED*** {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReviewSpec) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj6 int
	var yyb6 bool
	var yyhl6 bool = l >= 0
	yyj6++
	if yyhl6 {
		yyb6 = yyj6 > l
	} ***REMOVED*** {
		yyb6 = r.CheckBreak()
	}
	if yyb6 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Token = ""
	} ***REMOVED*** {
		yyv7 := &x.Token
		yym8 := z.DecBinary()
		_ = yym8
		if false {
		} ***REMOVED*** {
			*((*string)(yyv7)) = r.DecodeString()
		}
	}
	for {
		yyj6++
		if yyhl6 {
			yyb6 = yyj6 > l
		} ***REMOVED*** {
			yyb6 = r.CheckBreak()
		}
		if yyb6 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj6-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *TokenReviewStatus) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} ***REMOVED*** {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} ***REMOVED*** if z.HasExtensions() && z.EncExt(x) {
		} ***REMOVED*** {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [3]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Authenticated != false
			yyq2[1] = true
			yyq2[2] = x.Error != ""
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(3)
			} ***REMOVED*** {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} ***REMOVED*** {
						r.EncodeBool(bool(x.Authenticated))
					}
				} ***REMOVED*** {
					r.EncodeBool(false)
				}
			} ***REMOVED*** {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("authenticated"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} ***REMOVED*** {
						r.EncodeBool(bool(x.Authenticated))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yy7 := &x.User
					yy7.CodecEncodeSelf(e)
				} ***REMOVED*** {
					r.EncodeNil()
				}
			} ***REMOVED*** {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("user"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy9 := &x.User
					yy9.CodecEncodeSelf(e)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yym12 := z.EncBinary()
					_ = yym12
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Error))
					}
				} ***REMOVED*** {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} ***REMOVED*** {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("error"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym13 := z.EncBinary()
					_ = yym13
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Error))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *TokenReviewStatus) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} ***REMOVED*** if z.HasExtensions() && z.DecExt(x) {
	} ***REMOVED*** {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} ***REMOVED*** if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} ***REMOVED*** {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *TokenReviewStatus) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} ***REMOVED*** {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "authenticated":
			if r.TryDecodeAsNil() {
				x.Authenticated = false
			} ***REMOVED*** {
				yyv4 := &x.Authenticated
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} ***REMOVED*** {
					*((*bool)(yyv4)) = r.DecodeBool()
				}
			}
		case "user":
			if r.TryDecodeAsNil() {
				x.User = UserInfo{}
			} ***REMOVED*** {
				yyv6 := &x.User
				yyv6.CodecDecodeSelf(d)
			}
		case "error":
			if r.TryDecodeAsNil() {
				x.Error = ""
			} ***REMOVED*** {
				yyv7 := &x.Error
				yym8 := z.DecBinary()
				_ = yym8
				if false {
				} ***REMOVED*** {
					*((*string)(yyv7)) = r.DecodeString()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *TokenReviewStatus) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj9 int
	var yyb9 bool
	var yyhl9 bool = l >= 0
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} ***REMOVED*** {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Authenticated = false
	} ***REMOVED*** {
		yyv10 := &x.Authenticated
		yym11 := z.DecBinary()
		_ = yym11
		if false {
		} ***REMOVED*** {
			*((*bool)(yyv10)) = r.DecodeBool()
		}
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} ***REMOVED*** {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.User = UserInfo{}
	} ***REMOVED*** {
		yyv12 := &x.User
		yyv12.CodecDecodeSelf(d)
	}
	yyj9++
	if yyhl9 {
		yyb9 = yyj9 > l
	} ***REMOVED*** {
		yyb9 = r.CheckBreak()
	}
	if yyb9 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Error = ""
	} ***REMOVED*** {
		yyv13 := &x.Error
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} ***REMOVED*** {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	for {
		yyj9++
		if yyhl9 {
			yyb9 = yyj9 > l
		} ***REMOVED*** {
			yyb9 = r.CheckBreak()
		}
		if yyb9 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj9-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *UserInfo) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} ***REMOVED*** {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} ***REMOVED*** if z.HasExtensions() && z.EncExt(x) {
		} ***REMOVED*** {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [4]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Username != ""
			yyq2[1] = x.UID != ""
			yyq2[2] = len(x.Groups) != 0
			yyq2[3] = len(x.Extra) != 0
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(4)
			} ***REMOVED*** {
				yynn2 = 0
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Username))
					}
				} ***REMOVED*** {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} ***REMOVED*** {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("username"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.Username))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.UID))
					}
				} ***REMOVED*** {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} ***REMOVED*** {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("uid"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} ***REMOVED*** {
						r.EncodeString(codecSelferC_UTF81234, string(x.UID))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					if x.Groups == nil {
						r.EncodeNil()
					} ***REMOVED*** {
						yym10 := z.EncBinary()
						_ = yym10
						if false {
						} ***REMOVED*** {
							z.F.EncSliceStringV(x.Groups, false, e)
						}
					}
				} ***REMOVED*** {
					r.EncodeNil()
				}
			} ***REMOVED*** {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("groups"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Groups == nil {
						r.EncodeNil()
					} ***REMOVED*** {
						yym11 := z.EncBinary()
						_ = yym11
						if false {
						} ***REMOVED*** {
							z.F.EncSliceStringV(x.Groups, false, e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[3] {
					if x.Extra == nil {
						r.EncodeNil()
					} ***REMOVED*** {
						yym13 := z.EncBinary()
						_ = yym13
						if false {
						} ***REMOVED*** {
							h.encMapstringExtraValue((map[string]ExtraValue)(x.Extra), e)
						}
					}
				} ***REMOVED*** {
					r.EncodeNil()
				}
			} ***REMOVED*** {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("extra"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.Extra == nil {
						r.EncodeNil()
					} ***REMOVED*** {
						yym14 := z.EncBinary()
						_ = yym14
						if false {
						} ***REMOVED*** {
							h.encMapstringExtraValue((map[string]ExtraValue)(x.Extra), e)
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *UserInfo) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} ***REMOVED*** if z.HasExtensions() && z.DecExt(x) {
	} ***REMOVED*** {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} ***REMOVED*** if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} ***REMOVED*** {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} ***REMOVED*** {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *UserInfo) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} ***REMOVED*** {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "username":
			if r.TryDecodeAsNil() {
				x.Username = ""
			} ***REMOVED*** {
				yyv4 := &x.Username
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} ***REMOVED*** {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "uid":
			if r.TryDecodeAsNil() {
				x.UID = ""
			} ***REMOVED*** {
				yyv6 := &x.UID
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} ***REMOVED*** {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "groups":
			if r.TryDecodeAsNil() {
				x.Groups = nil
			} ***REMOVED*** {
				yyv8 := &x.Groups
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} ***REMOVED*** {
					z.F.DecSliceStringX(yyv8, false, d)
				}
			}
		case "extra":
			if r.TryDecodeAsNil() {
				x.Extra = nil
			} ***REMOVED*** {
				yyv10 := &x.Extra
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} ***REMOVED*** {
					h.decMapstringExtraValue((*map[string]ExtraValue)(yyv10), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *UserInfo) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Username = ""
	} ***REMOVED*** {
		yyv13 := &x.Username
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} ***REMOVED*** {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.UID = ""
	} ***REMOVED*** {
		yyv15 := &x.UID
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} ***REMOVED*** {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Groups = nil
	} ***REMOVED*** {
		yyv17 := &x.Groups
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} ***REMOVED*** {
			z.F.DecSliceStringX(yyv17, false, d)
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} ***REMOVED*** {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Extra = nil
	} ***REMOVED*** {
		yyv19 := &x.Extra
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} ***REMOVED*** {
			h.decMapstringExtraValue((*map[string]ExtraValue)(yyv19), d)
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} ***REMOVED*** {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x ExtraValue) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} ***REMOVED*** {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} ***REMOVED*** if z.HasExtensions() && z.EncExt(x) {
		} ***REMOVED*** {
			h.encExtraValue((ExtraValue)(x), e)
		}
	}
}

func (x *ExtraValue) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} ***REMOVED*** if z.HasExtensions() && z.DecExt(x) {
	} ***REMOVED*** {
		h.decExtraValue((*ExtraValue)(x), d)
	}
}

func (x codecSelfer1234) encMapstringExtraValue(v map[string]ExtraValue, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeMapStart(len(v))
	for yyk1, yyv1 := range v {
		z.EncSendContainerState(codecSelfer_containerMapKey1234)
		yym2 := z.EncBinary()
		_ = yym2
		if false {
		} ***REMOVED*** {
			r.EncodeString(codecSelferC_UTF81234, string(yyk1))
		}
		z.EncSendContainerState(codecSelfer_containerMapValue1234)
		if yyv1 == nil {
			r.EncodeNil()
		} ***REMOVED*** {
			yyv1.CodecEncodeSelf(e)
		}
	}
	z.EncSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x codecSelfer1234) decMapstringExtraValue(v *map[string]ExtraValue, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyl1 := r.ReadMapStart()
	yybh1 := z.DecBasicHandle()
	if yyv1 == nil {
		yyrl1, _ := z.DecInferLen(yyl1, yybh1.MaxInitLen, 40)
		yyv1 = make(map[string]ExtraValue, yyrl1)
		*v = yyv1
	}
	var yymk1 string
	var yymv1 ExtraValue
	var yymg1 bool
	if yybh1.MapValueReset {
		yymg1 = true
	}
	if yyl1 > 0 {
		for yyj1 := 0; yyj1 < yyl1; yyj1++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk1 = ""
			} ***REMOVED*** {
				yyv2 := &yymk1
				yym3 := z.DecBinary()
				_ = yym3
				if false {
				} ***REMOVED*** {
					*((*string)(yyv2)) = r.DecodeString()
				}
			}

			if yymg1 {
				yymv1 = yyv1[yymk1]
			} ***REMOVED*** {
				yymv1 = nil
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv1 = nil
			} ***REMOVED*** {
				yyv4 := &yymv1
				yyv4.CodecDecodeSelf(d)
			}

			if yyv1 != nil {
				yyv1[yymk1] = yymv1
			}
		}
	} ***REMOVED*** if yyl1 < 0 {
		for yyj1 := 0; !r.CheckBreak(); yyj1++ {
			z.DecSendContainerState(codecSelfer_containerMapKey1234)
			if r.TryDecodeAsNil() {
				yymk1 = ""
			} ***REMOVED*** {
				yyv5 := &yymk1
				yym6 := z.DecBinary()
				_ = yym6
				if false {
				} ***REMOVED*** {
					*((*string)(yyv5)) = r.DecodeString()
				}
			}

			if yymg1 {
				yymv1 = yyv1[yymk1]
			} ***REMOVED*** {
				yymv1 = nil
			}
			z.DecSendContainerState(codecSelfer_containerMapValue1234)
			if r.TryDecodeAsNil() {
				yymv1 = nil
			} ***REMOVED*** {
				yyv7 := &yymv1
				yyv7.CodecDecodeSelf(d)
			}

			if yyv1 != nil {
				yyv1[yymk1] = yymv1
			}
		}
	} // ***REMOVED*** len==0: TODO: Should we clear map entries?
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x codecSelfer1234) encExtraValue(v ExtraValue, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv1 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yym2 := z.EncBinary()
		_ = yym2
		if false {
		} ***REMOVED*** {
			r.EncodeString(codecSelferC_UTF81234, string(yyv1))
		}
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decExtraValue(v *ExtraValue, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []string{}
			yyc1 = true
		} ***REMOVED*** if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} ***REMOVED*** if yyl1 > 0 {
		var yyrr1, yyrl1 int
		var yyrt1 bool
		_, _ = yyrl1, yyrt1
		yyrr1 = yyl1 // len(yyv1)
		if yyl1 > cap(yyv1) {

			yyrl1, yyrt1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 16)
			if yyrt1 {
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} ***REMOVED*** {
					yyv1 = make([]string, yyrl1)
				}
			} ***REMOVED*** {
				yyv1 = make([]string, yyrl1)
			}
			yyc1 = true
			yyrr1 = len(yyv1)
		} ***REMOVED*** if yyl1 != len(yyv1) {
			yyv1 = yyv1[:yyl1]
			yyc1 = true
		}
		yyj1 := 0
		for ; yyj1 < yyrr1; yyj1++ {
			yyh1.ElemContainerState(yyj1)
			if r.TryDecodeAsNil() {
				yyv1[yyj1] = ""
			} ***REMOVED*** {
				yyv2 := &yyv1[yyj1]
				yym3 := z.DecBinary()
				_ = yym3
				if false {
				} ***REMOVED*** {
					*((*string)(yyv2)) = r.DecodeString()
				}
			}

		}
		if yyrt1 {
			for ; yyj1 < yyl1; yyj1++ {
				yyv1 = append(yyv1, "")
				yyh1.ElemContainerState(yyj1)
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = ""
				} ***REMOVED*** {
					yyv4 := &yyv1[yyj1]
					yym5 := z.DecBinary()
					_ = yym5
					if false {
					} ***REMOVED*** {
						*((*string)(yyv4)) = r.DecodeString()
					}
				}

			}
		}

	} ***REMOVED*** {
		yyj1 := 0
		for ; !r.CheckBreak(); yyj1++ {

			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, "") // var yyz1 string
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			if yyj1 < len(yyv1) {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = ""
				} ***REMOVED*** {
					yyv6 := &yyv1[yyj1]
					yym7 := z.DecBinary()
					_ = yym7
					if false {
					} ***REMOVED*** {
						*((*string)(yyv6)) = r.DecodeString()
					}
				}

			} ***REMOVED*** {
				z.DecSwallow()
			}

		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} ***REMOVED*** if yyj1 == 0 && yyv1 == nil {
			yyv1 = []string{}
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}
