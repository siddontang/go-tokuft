package tokuft

// #include "tokudb.h"
import "C"

import (
	"unsafe"
)

type Val C.DBT

// Create a Val that points to p's data. the Val's data must not be freed
// manually and C references must not survive the garbage collection of p (and
// the returned Val).
func Wrap(p []byte) Val {
	if len(p) == 0 {
		return Val(C.DBT{})
	}

	return Val(C.DBT{
		data:  unsafe.Pointer(&p[0]),
		size:  C.uint32_t(len(p)),
		ulen:  0,
		flags: 0,
	})
}

// If val is nil, a empty slice is retured.
func (val Val) Bytes() []byte {
	return C.GoBytes(val.data, C.int(val.size))
}

// If val is nil, an empty string is returned.
func (val Val) String() string {
	return C.GoStringN((*C.char)(val.data), C.int(val.size))
}
