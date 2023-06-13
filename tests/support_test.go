package tests

import "testing"

func TestGetTestDB(t *testing.T) {
	var count int

	db := GetTestDB()
	defer db.Close()

	db.Get(&count, "SELECT count(*) FROM sensors")
	if count != 3 {
		t.Fatal("sensors test table not correct")
	}

	db.Get(&count, "SELECT count(*) FROM sensors_data")
	if count != 5 {
		t.Fatal("sensors_data test table not correct")
	}
}

func TestGetEmptyTestDB(t *testing.T) {
	var count int

	db := GetTestDB() // get full db with rows
	db.Close()

	db = GetEmptyTestDB()
	defer db.Close()

	db.Get(&count, "SELECT count(*) FROM sensors")
	if count != 0 {
		t.Fatal("sensors test table not empty")
	}

}
