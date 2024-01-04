package ffmpeg

import (
	"fmt"
	"unsafe"
)

/*
#cgo pkg-config: libavcodec libavfilter libavformat libavutil

#include <stdlib.h>
*/
import "C"

type AVError struct {
	Code int
}

func (e *AVError) Error() string {
	return fmt.Sprintf("AVError %v", e.Code)
}

func WrapError(code int) error {
	if code >= 0 {
		return nil
	}

	return &AVError{Code: code}
}

func arrayGet[T any](array *T, i uintptr) T {
	var inner T
	ptrPtr := (*T)(unsafe.Add(unsafe.Pointer(array), i*unsafe.Sizeof(inner)))
	return *ptrPtr
}

type CStr struct {
	ptr *C.char
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
