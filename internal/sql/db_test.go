package sql

import "testing"

func TestCreateAndConnectDB(t *testing.T) {
	db, err := Connect("test")

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec("SELECT * FROM sensors")
	if err != nil {
		t.Fatal(err)
	}
}
