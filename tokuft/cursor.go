package tokuft

/*
#include "tokudb.h"
#include "tokudb_ext.h"
*/
import "C"

import ()

type Cursor struct {
	cursor *C.DBC
}

func (cursor *Cursor) Close() error {
	if cursor.cursor == nil {
		return nil
	}
	r := C.cursor_close(cursor.cursor)
	cursor.cursor = nil
	return errno(r)
}

func (cursor *Cursor) Get(key, val []byte, op uint) ([]byte, []byte, error) {
	ckey := Wrap(key)
	cval := Wrap(val)
	ret := C.cursor_get(cursor.cursor, (*C.DBT)(&ckey), (*C.DBT)(&cval), C.uint32_t(op))
	if ret != SUCCESS {
		if Errno(ret) == NOTFOUND {
			return nil, nil, nil
		}
		return nil, nil, errno(ret)
	}
	return ckey.Bytes(), cval.Bytes(), nil
}
