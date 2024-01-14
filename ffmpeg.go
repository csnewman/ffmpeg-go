package ffmpeg

import (
	"fmt"
	"unsafe"
)

/*
#cgo pkg-config: libavcodec libavfilter libavformat libavutil

#include <errno.h>
#include <stdlib.h>

#include <libavutil/avutil.h>
*/
import "C"

const (
	ptrSize   = uintptr(C.sizeof_size_t)
	intSize   = uintptr(C.sizeof_int)
	int8Size  = uintptr(C.sizeof_int8_t)
	int64Size = uintptr(C.sizeof_int64_t)
)

var AVTimeBaseQ = &AVRational{value: C.AV_TIME_BASE_Q}

var (
	EAgain     = AVError{Code: -C.EAGAIN}
	AVErrorEOF = AVError{Code: AVErrorEofConst}
)

type AVError struct {
	Code int
}

func (e AVError) Error() string {
	buf := AllocCStr(uint(AVErrorMaxStringSize))
	defer buf.Free()

	AVStrerror(e.Code, buf, uint64(AVErrorMaxStringSize))

	return fmt.Sprintf("averror %v: %v", e.Code, buf.String())
}

func WrapErr(code int) error {
	if code >= 0 {
		return nil
	}

	return AVError{Code: code}
}

type CStr struct {
	ptr *C.char
}

func AllocCStr(len uint) *CStr {
	ptr := (*C.char)(C.calloc(C.ulong(len), C.sizeof_char))

	return &CStr{
		ptr: ptr,
	}
}

func ToCStr(val string) *CStr {
	return &CStr{
		ptr: C.CString(val),
	}
}

func wrapCStr(ptr *C.char) *CStr {
	if ptr == nil {
		return nil
	}

	return &CStr{
		ptr: ptr,
	}
}

func (s *CStr) Dup() *CStr {
	return AVStrdup(s)
}

func (s *CStr) String() string {
	return C.GoString(s.ptr)
}

func (s *CStr) Free() {
	C.free(unsafe.Pointer(s.ptr))
}

func (s *CStr) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

type Array[T any] struct {
	ptr      unsafe.Pointer
	elemSize uintptr
	loadPtr  func(pointer unsafe.Pointer) T
	storePtr func(pointer unsafe.Pointer, value T)
}

func (a *Array[T]) Get(i uintptr) T {
	ptr := unsafe.Add(a.ptr, i*a.elemSize)
	return a.loadPtr(ptr)
}

func (a *Array[T]) Set(i uintptr, value T) {
	ptr := unsafe.Add(a.ptr, i*a.elemSize)
	a.storePtr(ptr, value)
}

func (a *Array[T]) Free() {
	AVFree(a.ptr)
}

func (a *Array[T]) RawPtr() unsafe.Pointer {
	return a.ptr
}

func ToIntArray(ptr unsafe.Pointer) *Array[int] {
	if ptr == nil {
		return nil
	}

	return &Array[int]{
		elemSize: intSize,
		loadPtr: func(pointer unsafe.Pointer) int {
			ptr := (*C.int)(pointer)
			return int(*ptr)
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value int) {
			ptr := (*C.int)(pointer)
			*ptr = C.int(value)
		},
	}
}

func ToUint8Array(ptr unsafe.Pointer) *Array[uint8] {
	if ptr == nil {
		return nil
	}

	return &Array[uint8]{
		elemSize: int8Size,
		loadPtr: func(pointer unsafe.Pointer) uint8 {
			ptr := (*C.uint8_t)(pointer)
			return uint8(*ptr)
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value uint8) {
			ptr := (*C.uint8_t)(pointer)
			*ptr = C.uint8_t(value)
		},
	}
}

func ToInt64Array(ptr unsafe.Pointer) *Array[int64] {
	if ptr == nil {
		return nil
	}

	return &Array[int64]{
		elemSize: int64Size,
		loadPtr: func(pointer unsafe.Pointer) int64 {
			ptr := (*C.int64_t)(pointer)
			return int64(*ptr)
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value int64) {
			ptr := (*C.int64_t)(pointer)
			*ptr = C.int64_t(value)
		},
	}
}

func ToUint64Array(ptr unsafe.Pointer) *Array[uint64] {
	if ptr == nil {
		return nil
	}

	return &Array[uint64]{
		elemSize: int64Size,
		loadPtr: func(pointer unsafe.Pointer) uint64 {
			ptr := (*C.uint64_t)(pointer)
			return uint64(*ptr)
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value uint64) {
			ptr := (*C.uint64_t)(pointer)
			*ptr = C.uint64_t(value)
		},
	}
}

func ToUint8PtrArray(ptr unsafe.Pointer) *Array[unsafe.Pointer] {
	if ptr == nil {
		return nil
	}

	return &Array[unsafe.Pointer]{
		elemSize: ptrSize,
		loadPtr: func(pointer unsafe.Pointer) unsafe.Pointer {
			ptr := (*unsafe.Pointer)(pointer)
			return *ptr
		},
		ptr: ptr,
		storePtr: func(pointer unsafe.Pointer, value unsafe.Pointer) {
			ptr := (*unsafe.Pointer)(pointer)
			*ptr = value
		},
	}
}
