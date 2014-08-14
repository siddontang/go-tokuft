package tokudb

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
	DB_RUNRECOVERY                 = Errno(C.DB_RUNRECOVERY)
	DB_KEYEXIST                    = Errno(C.DB_KEYEXIST)
	DB_LOCK_DEADLOCK               = Errno(C.DB_LOCK_DEADLOCK)
	DB_LOCK_NOTGRANTED             = Errno(C.DB_LOCK_NOTGRANTED)
	DB_NOTFOUND                    = Errno(C.DB_NOTFOUND)
	DB_SECONDARY_BAD               = Errno(C.DB_SECONDARY_BAD)
	DB_DONOTINDEX                  = Errno(C.DB_DONOTINDEX)
	DB_BUFFER_SMALL                = Errno(C.DB_BUFFER_SMALL)
	DB_BADFORMAT                   = Errno(C.DB_BADFORMAT)
	TOKUDB_OUT_OF_LOCKS            = Errno(C.TOKUDB_OUT_OF_LOCKS)
	TOKUDB_SUCCEEDED_EARLY         = Errno(C.TOKUDB_SUCCEEDED_EARLY)
	TOKUDB_FOUND_BUT_REJECTED      = Errno(C.TOKUDB_FOUND_BUT_REJECTED)
	TOKUDB_USER_CALLBACK_ERROR     = Errno(C.TOKUDB_USER_CALLBACK_ERROR)
	TOKUDB_DICTIONARY_TOO_OLD      = Errno(C.TOKUDB_DICTIONARY_TOO_OLD)
	TOKUDB_DICTIONARY_TOO_NEW      = Errno(C.TOKUDB_DICTIONARY_TOO_NEW)
	TOKUDB_DICTIONARY_NO_HEADER    = Errno(C.TOKUDB_DICTIONARY_NO_HEADER)
	TOKUDB_CANCELED                = Errno(C.TOKUDB_CANCELED)
	TOKUDB_NO_DATA                 = Errno(C.TOKUDB_NO_DATA)
	TOKUDB_ACCEPT                  = Errno(C.TOKUDB_ACCEPT)
	TOKUDB_MVCC_DICTIONARY_TOO_NEW = Errno(C.TOKUDB_MVCC_DICTIONARY_TOO_NEW)
	TOKUDB_UPGRADE_FAILURE         = Errno(C.TOKUDB_UPGRADE_FAILURE)
	TOKUDB_TRY_AGAIN               = Errno(C.TOKUDB_TRY_AGAIN)
	TOKUDB_NEEDS_REPAIR            = Errno(C.TOKUDB_NEEDS_REPAIR)
	TOKUDB_CURSOR_CONTINUE         = Errno(C.TOKUDB_CURSOR_CONTINUE)
	TOKUDB_BAD_CHECKSUM            = Errno(C.TOKUDB_BAD_CHECKSUM)
	TOKUDB_HUGE_PAGES_ENABLED      = Errno(C.TOKUDB_HUGE_PAGES_ENABLED)
	TOKUDB_OUT_OF_RANGE            = Errno(C.TOKUDB_OUT_OF_RANGE)
	TOKUDB_INTERRUPTED             = Errno(C.TOKUDB_INTERRUPTED)
)
