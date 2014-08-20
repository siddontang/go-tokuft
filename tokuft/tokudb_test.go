package tokuft

import (
	"fmt"
	"os"
	"testing"
)

func TestBase(t *testing.T) {
	env, err := NewEnv()
	if err != nil {
		t.Fatal(err)
	}

	dir := "./var"
	os.RemoveAll(dir)
	if err = env.Open(dir, CREATE|PRIVATE|INIT_MPOOL|INIT_LOCK|INIT_TXN, 0644); err != nil {
		t.Fatal(err)
	}

	var tx *Tx
	if tx, err = env.BeginTx(nil, 0); err != nil {
		t.Fatal(err)
	}
	defer tx.Abort()

	var db *DB
	if db, err = tx.OpenDB("test.db", CREATE, 0644); err != nil {
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

	for i := 0; i < 10; i++ {
		if err := tx.Put(db, []byte(fmt.Sprintf("%d", i)), nil); err != nil {
			t.Fatal(err)
		}
	}

	var c *Cursor
	if c, err = tx.Cursor(db); err != nil {
		t.Fatal(err)
	}

	if k, _, err := c.Get(nil, nil, FIRST); err != nil {
		t.Fatal(err)
	} else if string(k) != "0" {
		t.Fatal(string(k))
	}

	for i := 1; i < 10; i++ {
		if k, _, err := c.Get(nil, nil, NEXT); err != nil {
			t.Fatal(err)
		} else if string(k) != fmt.Sprintf("%d", i) {
			t.Fatal(string(k))
		}

	}

	if k, _, err := c.Get(nil, nil, LAST); err != nil {
		t.Fatal(err)
	} else if string(k) != "9" {
		t.Fatal(string(k))
	}

	if k, _, err := c.Get(nil, nil, NEXT); err != nil {
		t.Fatal(err)
	} else if k != nil {
		t.Fatal("must nil")
	}

	if k, _, err := c.Get(nil, nil, LAST); err != nil {
		t.Fatal(err)
	} else if string(k) != "9" {
		t.Fatal(string(k))
	}

	for i := 8; i >= 0; i-- {
		if k, _, err := c.Get(nil, nil, PREV); err != nil {
			t.Fatal(err)
		} else if string(k) != fmt.Sprintf("%d", i) {
			t.Fatal(string(k))
		}
	}

	if k, _, err := c.Get(nil, nil, PREV); err != nil {
		t.Fatal(err)
	} else if k != nil {
		t.Fatal("must nil")
	}

	if k, _, err := c.Get([]byte("5"), nil, SET); err != nil {
		t.Fatal(err)
	} else if string(k) != "5" {
		t.Fatal(string(k))
	}

	if _, err := tx.Get(db, []byte("1")); err != nil {
		t.Fatal(err)
	}

	if err := tx.Delete(db, []byte("5")); err != nil {
		t.Fatal(err)
	}

	if k, _, err := c.Get([]byte("5"), nil, SET_RANGE); err != nil {
		t.Fatal(err)
	} else if string(k) != "6" {
		t.Fatal(string(k))
	}

	if err := c.Close(); err != nil {
		t.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		t.Fatal(err)
	}

	tx, _ = env.BeginTx(nil, 0)

	if err = tx.Put(db, []byte("invalid_key"), []byte("invalid_value")); err != nil {
		t.Fatal(err)
	}

	tx.Abort()

	tx, _ = env.BeginTx(nil, TXN_SNAPSHOT|TXN_READ_ONLY)

	if v, err := tx.Get(db, []byte("invalid_key")); err != nil {
		t.Fatal(err)
	} else if v != nil {
		t.Fatal("must nil")
	}

	tx.Abort()

	tx, _ = env.BeginTx(nil, TXN_SNAPSHOT|TXN_READ_ONLY)

	c, err = tx.Cursor(db)
	if err != nil {
		t.Fatal(err)
	}

	if tx1, err := env.BeginTx(nil, 0); err != nil {
		t.Fatal(err)
	} else {
		if err = tx1.Delete(db, []byte("key")); err != nil {
			t.Fatal(err)
		}

		if err = tx1.Commit(); err != nil {
			t.Fatal(err)
		}
	}

	tx.Abort()

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}

	if err := env.Close(); err != nil {
		t.Fatal(err)
	}
}
