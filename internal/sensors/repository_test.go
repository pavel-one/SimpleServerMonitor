package sensors

import (
	"github.com/pavel-one/sensors/tests"
	"testing"
)

func TestSensorRepository_Find(t *testing.T) {
	db := tests.GetTestDB()
	if db == nil {
		t.Fatal("not init database")
	}

	rep := NewSensorRepository(db)

	model, err := rep.Find(1)
	if err != nil {
		t.Fatal(err)
	}

	if len(model.Data) != 3 {
		t.Fatal("model data not set: ", model)
	}

	model, err = rep.Find(3)
	if err != nil {
		t.Fatal(err)
	}

	if model.Data != nil {
		t.Fatal("model data is set: ", model)
	}
}

func TestSensorRepository_FindWithColumn(t *testing.T) {
	db := tests.GetTestDB()
	if db == nil {
		t.Fatal("not init database")
	}

	rep := NewSensorRepository(db)

	m, err := rep.FindWithColumn("name", "test")
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "test" {
		t.Fatalf("model not test: %s", m.Name)
	}

	m, err = rep.FindWithColumn("name", "test1")
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "test1" {
		t.Fatalf("model not test1: %s", m.Name)
	}
}

func TestSensorRepository_AddTemp(t *testing.T) {
	db := tests.GetEmptyTestDB()
	if db == nil {
		t.Fatal("not init database")
	}

	rep := NewSensorRepository(db)

	err := rep.AddTemp(&Sensor{
		Name:     "test123",
		Temp:     25.5,
		HighTemp: 50,
		CritTemp: 110,
	})
	if err != nil {
		t.Fatal(err)
	}

	m, err := rep.FindWithColumn("name", "test123")
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "test123" {
		t.Fatalf("model not test: %s", m.Name)
	}

	// add from exists record
	err = rep.AddTemp(&Sensor{
		Name:     "test123",
		Temp:     26,
		HighTemp: 50,
		CritTemp: 110,
	})
	if err != nil {
		t.Fatal(err)
	}

	m, err = rep.FindWithColumn("name", "test123")
	if err != nil {
		t.Fatal(err)
	}

	if m.Name != "test123" {
		t.Fatalf("model not test 2: %s", m.Name)
	}

	if len(m.Data) != 2 {
		t.Fatal("count data records not correct")
	}
}
