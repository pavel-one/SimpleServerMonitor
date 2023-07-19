package db

import (
	"os"
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect("test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("create table foo (id integer not null primary key, name text); delete from foo;")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Remove("./test.sqlite3"); err != nil {
		return
	}
}
