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

func arrayGet[T any](array *T, i uintptr) T {
	var inner T
	ptrPtr := (*T)(unsafe.Add(unsafe.Pointer(array), i*unsafe.Sizeof(inner)))
	return *ptrPtr
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
