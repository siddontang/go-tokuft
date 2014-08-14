package tokuft

/*
 #include "tokudb.h"
 #include "tokudb_ext.h"
 #include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

type DB struct {
	db *C.DB
}

func (tx *Tx) OpenDB(dbName string, flags uint, mode int) (*DB, error) {
	db := new(DB)

	r := C.db_create(&db.db, tx.e.env, 0)
	if r != SUCCESS {
		return nil, errno(r)
	}

	cname := C.CString(dbName)
	defer C.free(unsafe.Pointer(cname))

	if r = C.db_open(db.db, tx.tx, cname, C.uint32_t(flags), C.int(mode)); r != SUCCESS {
		return nil, errno(r)
	}
	return db, nil
}

func (db *DB) Close() error {
	r := C.db_close(db.db, 0)

	db.db = nil
	return errno(r)
}
