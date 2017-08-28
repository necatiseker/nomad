// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package agent

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	io "io"
	"reflect"
	"runtime"
)

const (
	// ----- content types ----
	codecSelferC_UTF85213 = 1
	codecSelferC_RAW5213  = 0
	// ----- value types used ----
	codecSelferValueTypeArray5213 = 10
	codecSelferValueTypeMap5213   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey5213    = 2
	codecSelfer_containerMapValue5213  = 3
	codecSelfer_containerMapEnd5213    = 4
	codecSelfer_containerArrayElem5213 = 6
	codecSelfer_containerArrayEnd5213  = 7
)

var (
	codecSelferBitsize5213                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr5213 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer5213 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 io.Reader
		_ = v0
	}
}

func (x *ReadCloserWrapper) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [2]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(2)
			} else {
				yynn2 = 2
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem5213)
				if x.Reader == nil {
					r.EncodeNil()
				} else {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Reader) {
					} else {
						z.EncFallback(x.Reader)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey5213)
				r.EncodeString(codecSelferC_UTF85213, string("Reader"))
				z.EncSendContainerState(codecSelfer_containerMapValue5213)
				if x.Reader == nil {
					r.EncodeNil()
				} else {
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Reader) {
					} else {
						z.EncFallback(x.Reader)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem5213)
				if x.Closer == nil {
					r.EncodeNil()
				} else {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Closer) {
					} else {
						z.EncFallback(x.Closer)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey5213)
				r.EncodeString(codecSelferC_UTF85213, string("Closer"))
				z.EncSendContainerState(codecSelfer_containerMapValue5213)
				if x.Closer == nil {
					r.EncodeNil()
				} else {
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else if z.HasExtensions() && z.EncExt(x.Closer) {
					} else {
						z.EncFallback(x.Closer)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd5213)
			}
		}
	}
}

func (x *ReadCloserWrapper) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap5213 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd5213)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray5213 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr5213)
		}
	}
}

func (x *ReadCloserWrapper) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
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
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey5213)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue5213)
		switch yys3 {
		case "Reader":
			if r.TryDecodeAsNil() {
				x.Reader = nil
			} else {
				yyv4 := &x.Reader
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv4) {
				} else {
					z.DecFallback(yyv4, true)
				}
			}
		case "Closer":
			if r.TryDecodeAsNil() {
				x.Closer = nil
			} else {
				yyv6 := &x.Closer
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv6) {
				} else {
					z.DecFallback(yyv6, true)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd5213)
}

func (x *ReadCloserWrapper) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj8 int
	var yyb8 bool
	var yyhl8 bool = l >= 0
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5213)
	if r.TryDecodeAsNil() {
		x.Reader = nil
	} else {
		yyv9 := &x.Reader
		yym10 := z.DecBinary()
		_ = yym10
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv9) {
		} else {
			z.DecFallback(yyv9, true)
		}
	}
	yyj8++
	if yyhl8 {
		yyb8 = yyj8 > l
	} else {
		yyb8 = r.CheckBreak()
	}
	if yyb8 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5213)
	if r.TryDecodeAsNil() {
		x.Closer = nil
	} else {
		yyv11 := &x.Closer
		yym12 := z.DecBinary()
		_ = yym12
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv11) {
		} else {
			z.DecFallback(yyv11, true)
		}
	}
	for {
		yyj8++
		if yyhl8 {
			yyb8 = yyj8 > l
		} else {
			yyb8 = r.CheckBreak()
		}
		if yyb8 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem5213)
		z.DecStructFieldNotFound(yyj8-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
}

func (x *StreamFrame) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [4]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Offset != 0
			yyq2[1] = len(x.Data) != 0
			yyq2[2] = x.File != ""
			yyq2[3] = x.FileEvent != ""
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(4)
			} else {
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
				z.EncSendContainerState(codecSelfer_containerArrayElem5213)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeInt(int64(x.Offset))
					}
				} else {
					r.EncodeInt(0)
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey5213)
					r.EncodeString(codecSelferC_UTF85213, string("Offset"))
					z.EncSendContainerState(codecSelfer_containerMapValue5213)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeInt(int64(x.Offset))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem5213)
				if yyq2[1] {
					if x.Data == nil {
						r.EncodeNil()
					} else {
						yym7 := z.EncBinary()
						_ = yym7
						if false {
						} else {
							r.EncodeStringBytes(codecSelferC_RAW5213, []byte(x.Data))
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey5213)
					r.EncodeString(codecSelferC_UTF85213, string("Data"))
					z.EncSendContainerState(codecSelfer_containerMapValue5213)
					if x.Data == nil {
						r.EncodeNil()
					} else {
						yym8 := z.EncBinary()
						_ = yym8
						if false {
						} else {
							r.EncodeStringBytes(codecSelferC_RAW5213, []byte(x.Data))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem5213)
				if yyq2[2] {
					yym10 := z.EncBinary()
					_ = yym10
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF85213, string(x.File))
					}
				} else {
					r.EncodeString(codecSelferC_UTF85213, "")
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey5213)
					r.EncodeString(codecSelferC_UTF85213, string("File"))
					z.EncSendContainerState(codecSelfer_containerMapValue5213)
					yym11 := z.EncBinary()
					_ = yym11
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF85213, string(x.File))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem5213)
				if yyq2[3] {
					yym13 := z.EncBinary()
					_ = yym13
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF85213, string(x.FileEvent))
					}
				} else {
					r.EncodeString(codecSelferC_UTF85213, "")
				}
			} else {
				if yyq2[3] {
					z.EncSendContainerState(codecSelfer_containerMapKey5213)
					r.EncodeString(codecSelferC_UTF85213, string("FileEvent"))
					z.EncSendContainerState(codecSelfer_containerMapValue5213)
					yym14 := z.EncBinary()
					_ = yym14
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF85213, string(x.FileEvent))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd5213)
			}
		}
	}
}

func (x *StreamFrame) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap5213 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd5213)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray5213 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr5213)
		}
	}
}

func (x *StreamFrame) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
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
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey5213)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue5213)
		switch yys3 {
		case "Offset":
			if r.TryDecodeAsNil() {
				x.Offset = 0
			} else {
				yyv4 := &x.Offset
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*int64)(yyv4)) = int64(r.DecodeInt(64))
				}
			}
		case "Data":
			if r.TryDecodeAsNil() {
				x.Data = nil
			} else {
				yyv6 := &x.Data
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*yyv6 = r.DecodeBytes(*(*[]byte)(yyv6), false, false)
				}
			}
		case "File":
			if r.TryDecodeAsNil() {
				x.File = ""
			} else {
				yyv8 := &x.File
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else {
					*((*string)(yyv8)) = r.DecodeString()
				}
			}
		case "FileEvent":
			if r.TryDecodeAsNil() {
				x.FileEvent = ""
			} else {
				yyv10 := &x.FileEvent
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					*((*string)(yyv10)) = r.DecodeString()
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd5213)
}

func (x *StreamFrame) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5213)
	if r.TryDecodeAsNil() {
		x.Offset = 0
	} else {
		yyv13 := &x.Offset
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*int64)(yyv13)) = int64(r.DecodeInt(64))
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5213)
	if r.TryDecodeAsNil() {
		x.Data = nil
	} else {
		yyv15 := &x.Data
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*yyv15 = r.DecodeBytes(*(*[]byte)(yyv15), false, false)
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5213)
	if r.TryDecodeAsNil() {
		x.File = ""
	} else {
		yyv17 := &x.File
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else {
			*((*string)(yyv17)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5213)
	if r.TryDecodeAsNil() {
		x.FileEvent = ""
	} else {
		yyv19 := &x.FileEvent
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			*((*string)(yyv19)) = r.DecodeString()
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem5213)
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
}

func (x *StreamFramer) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [1]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(1)
			} else {
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
				z.EncSendContainerState(codecSelfer_containerArrayElem5213)
				if x.Err == nil {
					r.EncodeNil()
				} else {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						z.EncFallback(x.Err)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey5213)
				r.EncodeString(codecSelferC_UTF85213, string("Err"))
				z.EncSendContainerState(codecSelfer_containerMapValue5213)
				if x.Err == nil {
					r.EncodeNil()
				} else {
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						z.EncFallback(x.Err)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd5213)
			}
		}
	}
}

func (x *StreamFramer) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap5213 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd5213)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray5213 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr5213)
		}
	}
}

func (x *StreamFramer) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
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
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey5213)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue5213)
		switch yys3 {
		case "Err":
			if r.TryDecodeAsNil() {
				x.Err = nil
			} else {
				yyv4 := &x.Err
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					z.DecFallback(yyv4, true)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd5213)
}

func (x *StreamFramer) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj6 int
	var yyb6 bool
	var yyhl6 bool = l >= 0
	yyj6++
	if yyhl6 {
		yyb6 = yyj6 > l
	} else {
		yyb6 = r.CheckBreak()
	}
	if yyb6 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem5213)
	if r.TryDecodeAsNil() {
		x.Err = nil
	} else {
		yyv7 := &x.Err
		yym8 := z.DecBinary()
		_ = yym8
		if false {
		} else {
			z.DecFallback(yyv7, true)
		}
	}
	for {
		yyj6++
		if yyhl6 {
			yyb6 = yyj6 > l
		} else {
			yyb6 = r.CheckBreak()
		}
		if yyb6 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem5213)
		z.DecStructFieldNotFound(yyj6-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
}

func (x *indexTuple) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [0]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(0)
			} else {
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
				z.EncSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd5213)
			}
		}
	}
}

func (x *indexTuple) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap5213 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd5213)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray5213 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr5213)
		}
	}
}

func (x *indexTuple) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
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
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey5213)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue5213)
		switch yys3 {
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd5213)
}

func (x *indexTuple) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj4 int
	var yyb4 bool
	var yyhl4 bool = l >= 0
	for {
		yyj4++
		if yyhl4 {
			yyb4 = yyj4 > l
		} else {
			yyb4 = r.CheckBreak()
		}
		if yyb4 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem5213)
		z.DecStructFieldNotFound(yyj4-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd5213)
}

func (x indexTupleArray) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			h.encindexTupleArray((indexTupleArray)(x), e)
		}
	}
}

func (x *indexTupleArray) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		h.decindexTupleArray((*indexTupleArray)(x), d)
	}
}

func (x codecSelfer5213) encindexTupleArray(v indexTupleArray, e *codec1978.Encoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv1 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem5213)
		yy2 := &yyv1
		yy2.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd5213)
}

func (x codecSelfer5213) decindexTupleArray(v *indexTupleArray, d *codec1978.Decoder) {
	var h codecSelfer5213
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []indexTuple{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else if yyl1 > 0 {
		var yyrr1, yyrl1 int
		var yyrt1 bool
		_, _ = yyrl1, yyrt1
		yyrr1 = yyl1 // len(yyv1)
		if yyl1 > cap(yyv1) {

			yyrg1 := len(yyv1) > 0
			yyv21 := yyv1
			yyrl1, yyrt1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 16)
			if yyrt1 {
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]indexTuple, yyrl1)
				}
			} else {
				yyv1 = make([]indexTuple, yyrl1)
			}
			yyc1 = true
			yyrr1 = len(yyv1)
			if yyrg1 {
				copy(yyv1, yyv21)
			}
		} else if yyl1 != len(yyv1) {
			yyv1 = yyv1[:yyl1]
			yyc1 = true
		}
		yyj1 := 0
		for ; yyj1 < yyrr1; yyj1++ {
			yyh1.ElemContainerState(yyj1)
			if r.TryDecodeAsNil() {
				yyv1[yyj1] = indexTuple{}
			} else {
				yyv2 := &yyv1[yyj1]
				yyv2.CodecDecodeSelf(d)
			}

		}
		if yyrt1 {
			for ; yyj1 < yyl1; yyj1++ {
				yyv1 = append(yyv1, indexTuple{})
				yyh1.ElemContainerState(yyj1)
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = indexTuple{}
				} else {
					yyv3 := &yyv1[yyj1]
					yyv3.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj1 := 0
		for ; !r.CheckBreak(); yyj1++ {

			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, indexTuple{}) // var yyz1 indexTuple
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			if yyj1 < len(yyv1) {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = indexTuple{}
				} else {
					yyv4 := &yyv1[yyj1]
					yyv4.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = []indexTuple{}
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}
