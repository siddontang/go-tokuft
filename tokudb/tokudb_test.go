package tokudb

import (
	//"os"
	"testing"
)

func TestBase(t *testing.T) {
	env, err := NewEnv()
	if err != nil {
		t.Fatal(err)
	}

	dir := "./var"
	//os.RemoveAll(dir)
	if err = env.Open(dir, DB_CREATE|DB_PRIVATE|DB_INIT_MPOOL|DB_INIT_LOCK|DB_INIT_TXN, 0644); err != nil {
		t.Fatal(err)
	}

	var tx *Tx
	if tx, err = env.BeginTx(nil, 0); err != nil {
		t.Fatal(err)
	}
	defer tx.Abort()

	var db *DB
	if db, err = tx.OpenDB("test.db", DB_CREATE, 0644); err != nil {
		t.Fatal(err)
	}

	if err := tx.Put(db, []byte("key"), []byte("value")); err != nil {
		t.Fatal(err)
	}

	if v, err := tx.Get(db, []byte("key")); err != nil {
		t.Fatal(err)
	} else if string(v) != "value" {
		t.Fatal(string(v))
	}

	if err := tx.Delete(db, []byte("key")); err != nil {
		t.Fatal(err)
	}

	if v, err := tx.Get(db, []byte("key")); err != nil {
		t.Fatal(err)
	} else if v != nil {
		t.Fatal("must nil")
	}

	if err := tx.Delete(db, []byte("key")); err != nil {
		t.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatal(err)
	}

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	if err := env.Close(); err != nil {
		t.Fatal(err)
	}
}
