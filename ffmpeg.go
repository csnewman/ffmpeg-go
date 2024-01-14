package ffmpeg

import (
	"fmt"
	"sync"
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

// AVError represents a non-positive return code from FFmpeg.
type AVError struct {
	Code int
}

func (e AVError) Error() string {
	buf := AllocCStr(uint(AVErrorMaxStringSize))
	defer buf.Free()

	AVStrerror(e.Code, buf, uint64(AVErrorMaxStringSize))

	return fmt.Sprintf("averror %v: %v", e.Code, buf.String())
}

// WrapErr returns a AVError if the code is less than zero, otherwise nil.
func WrapErr(code int) error {
	if code >= 0 {
		return nil
	}

	return AVError{Code: code}
}

// CStr is a string allocated in the C memory space. You may need to call Free to clean up the string depending on the
// owner and use-case.
type CStr struct {
	ptr      *C.char
	dontFree bool
}

// AllocCStr allocates an empty string with the given length. The buffer will be initialised to 0.
func AllocCStr(len uint) *CStr {
	ptr := (*C.char)(C.calloc(C.ulong(len), C.sizeof_char))

	return &CStr{
		ptr: ptr,
	}
}

// ToCStr allocates a new CStr with the given content. The CStr will not be automatically garbage collected.
func ToCStr(val string) *CStr {
	return &CStr{
		ptr: C.CString(val),
	}
}

var (
	strMap  = map[string]*CStr{}
	strLock = sync.RWMutex{}
)

// GlobalCStr resolves the given string to a CStr. Multiple calls with the same input string will return the same CStr.
// You should not attempt to free the CStr returned. When passing to FFmpeg, you may need to call Dup to create a copy
// if the FFmpeg code expects to take ownership and will likely free the string.
func GlobalCStr(val string) *CStr {
	var (
		ptr *CStr
		ok  bool
	)

	strLock.RLock()
	ptr, ok = strMap[val]
	strLock.RUnlock()

	if ok {
		return ptr
	}

	strLock.Lock()
	defer strLock.Unlock()

	ptr, ok = strMap[val]
	if ok {
		return ptr
	}

	ptr = ToCStr(val)
	ptr.dontFree = true
	strMap[val] = ptr

	return ptr
}

func wrapCStr(ptr *C.char) *CStr {
	if ptr == nil {
		return nil
	}

	return &CStr{
		ptr: ptr,
	}
}

// Dup is a wrapper for AVStrdup.
func (s *CStr) Dup() *CStr {
	return AVStrdup(s)
}

// String converts the CStr to a Go string.
func (s *CStr) String() string {
	return C.GoString(s.ptr)
}

// Free frees the backing memory for this string. You should only call this function if you are the owner of the memory.
func (s *CStr) Free() {
	if s.dontFree {
		return
	}

	C.free(unsafe.Pointer(s.ptr))
}

// RawPtr returns a raw reference to the underlying allocation.
func (s *CStr) RawPtr() unsafe.Pointer {
	return unsafe.Pointer(s.ptr)
}

// Array is a helper utility for accessing arrays of FFmpeg types. You can not directly allocate this type, and you must
// use one of the inbuilt constructors, such as AllocAVCodecIDArray.
//
// Arrays have no inbuilt length, matching the behaviour of C code. Getting or setting an out of bound index will lead
// to undefined behaviour.
type Array[T any] struct {
	ptr      unsafe.Pointer
	elemSize uintptr
	loadPtr  func(pointer unsafe.Pointer) T
	storePtr func(pointer unsafe.Pointer, value T)
}

// Get returns the element at the ith offset.
func (a *Array[T]) Get(i uintptr) T {
	ptr := unsafe.Add(a.ptr, i*a.elemSize)
	return a.loadPtr(ptr)
}

// Set sets the element at the ith offset.
func (a *Array[T]) Set(i uintptr, value T) {
	ptr := unsafe.Add(a.ptr, i*a.elemSize)
	a.storePtr(ptr, value)
}

// Free deallocates the underlying memory. You should only call this if you own the array.
func (a *Array[T]) Free() {
	AVFree(a.ptr)
}

// RawPtr returns a raw handle the underlying allocation.
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
