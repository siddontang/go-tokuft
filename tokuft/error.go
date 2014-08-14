package tokuft

// #include "tokudb.h"
import "C"

import (
	"fmt"
	"syscall"
)

const SUCCESS C.int = C.int(0)

type Errno C.int

func (e Errno) Error() string {
	s := C.GoString(C.db_strerror(C.int(e)))
	if s == "" {
		return fmt.Sprint("tokudb errno %d:%s", int(e), syscall.Errno(C.int(e)))
	} else {
		return fmt.Sprintf("tokudb error %d:%s", int(e), s)
	}
}

// for tests that can't import C
func _errno(ret int) error {
	return errno(C.int(ret))
}

func errno(ret C.int) error {
	if ret == SUCCESS {
		return nil
	}

	return Errno(ret)
}

const (
	RUNRECOVERY     = Errno(C.DB_RUNRECOVERY)
	KEYEXIST        = Errno(C.DB_KEYEXIST)
	LOCK_DEADLOCK   = Errno(C.DB_LOCK_DEADLOCK)
	LOCK_NOTGRANTED = Errno(C.DB_LOCK_NOTGRANTED)
	NOTFOUND        = Errno(C.DB_NOTFOUND)
	SECONDARY_BAD   = Errno(C.DB_SECONDARY_BAD)
	DONOTINDEX      = Errno(C.DB_DONOTINDEX)
	BUFFER_SMALL    = Errno(C.DB_BUFFER_SMALL)
	BADFORMAT       = Errno(C.DB_BADFORMAT)
)
