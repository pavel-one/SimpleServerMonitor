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
