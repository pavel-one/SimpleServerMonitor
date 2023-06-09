package tests

import "testing"

func TestGetTestDB(t *testing.T) {
	var count int

	db := GetTestDB()

	db.Get(&count, "SELECT count(*) FROM sensors")
	if count != 3 {
		t.Fatal("sensors test table not correct")
	}

	db.Get(&count, "SELECT count(*) FROM sensors_data")
	if count != 5 {
		t.Fatal("sensors_data test table not correct")
	}
}
