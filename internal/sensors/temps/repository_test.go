package temps

import (
	"os"
	"testing"
)

func TestRepository_Save(t *testing.T) {
	stat := Stat{
		Temp: 25.0,
		Key:  "test",
	}

	rep, err := NewRepository()
	defer os.Remove("db.sqlite3")

	if err != nil {
		t.Fatal(err)
	}

	if err := rep.Save(&stat); err != nil {
		t.Fatal(err)
	}

	if err := rep.DB.Close(); err != nil {
		t.Fatal(err)
	}

	if err := rep.Save(&stat); err == nil {
		t.Fatal("error stat has been saving :C")
	}
}

//func TestRepository_GetBySeconds(t *testing.T) {
//	rep, err := NewRepository()
//	defer os.Remove("db.sqlite3")
//
//	stats := []Stat{
//		{
//			Temp: 25,
//			Key:  "test",
//		},
//	}
//}
