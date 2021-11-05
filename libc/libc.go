package libc

import (
	"unsafe"
)

//go:linkname runtime_libcCall runtime.libcCall
func runtime_libcCall(fn, arg unsafe.Pointer) int32

//go:nosplit
func Malloc(size uintptr) uintptr {
	args := struct {
		size uintptr
		ret  uintptr
	}{
		size: size,
		ret:  0,
	}
	runtime_libcCall(unsafe.Pointer(libc_malloc), unsafe.Pointer(&args))
	return args.ret
}

//go:cgo_import_dynamic libc_malloc malloc "libc.dylib"

//go:linkname libc_malloc libc_malloc
var libc_malloc uintptr
