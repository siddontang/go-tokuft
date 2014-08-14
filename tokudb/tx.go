package tokudb

/*
 #include "tokudb.h"
 #include "tokudb_ext.h"
*/
import "C"

import (
	"errors"
	"runtime"
)

var (
	ErrTxDone = errors.New("transaction has done")
)

type Tx struct {
	tx *C.DB_TXN

	e *Env
}

func (e *Env) BeginTx(parent *Tx, flags uint) (*Tx, error) {
	tx := new(Tx)

	tx.e = e

	runtime.LockOSThread()

	var ptx *C.DB_TXN = nil
	if parent != nil {
		ptx = parent.tx
	}

	r := C.env_txn_begin(e.env, ptx, &tx.tx, C.uint32_t(flags))
	if r != SUCCESS {
		runtime.UnlockOSThread()
		return nil, errno(r)
	}

	return tx, nil
}

func (tx *Tx) Commit() error {
	if tx.tx == nil {
		return ErrTxDone
	}

	r := C.txn_commit(tx.tx, 0)
	runtime.UnlockOSThread()
	tx.tx = nil
	return errno(r)
}

func (tx *Tx) Abort() error {
	if tx.tx == nil {
		return ErrTxDone
	}

	r := C.txn_abort(tx.tx)
	runtime.UnlockOSThread()
	tx.tx = nil
	return errno(r)
}

func (tx *Tx) Get(db *DB, key []byte) ([]byte, error) {
	ckey := Wrap(key)
	var cval Val
	ret := C.db_get(db.db, tx.tx, (*C.DBT)(&ckey), (*C.DBT)(&cval), 0)
	if Errno(ret) == DB_NOTFOUND {
		return nil, nil
	}
	return cval.Bytes(), errno(ret)
}

func (tx *Tx) Put(db *DB, key []byte, value []byte) error {
	ckey := Wrap(key)
	cval := Wrap(value)
	ret := C.db_put(db.db, tx.tx, (*C.DBT)(&ckey), (*C.DBT)(&cval), 0)
	return errno(ret)
}

func (tx *Tx) Delete(db *DB, key []byte) error {
	ckey := Wrap(key)
	ret := C.db_del(db.db, tx.tx, (*C.DBT)(&ckey), 0)
	if Errno(ret) == DB_NOTFOUND {
		return nil
	}

	return errno(ret)
}
