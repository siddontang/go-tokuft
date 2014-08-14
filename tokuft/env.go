package tokuft

/*
 #include "tokudb.h"
 #include "tokudb_ext.h"
 #include <stdlib.h>
*/
import "C"

import (
	"os"
	"unsafe"
)

type Env struct {
	env *C.DB_ENV
}

func NewEnv() (*Env, error) {
	e := new(Env)

	r := C.db_env_create(&e.env, 0)
	if r != SUCCESS {
		return nil, errno(r)
	}

	return e, nil
}

func (e *Env) Open(path string, flags uint, mode int) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}

	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	r := C.env_open(e.env, cpath, C.uint32_t(flags), C.int(mode))
	if r != SUCCESS {
		return errno(r)
	}

	return nil
}

func (e *Env) Close() error {
	r := C.env_close(e.env, 0)
	if r != SUCCESS {
		return errno(r)
	}

	return nil
}
