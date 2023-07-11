package temps

import (
	"testing"
)

func TestGetSensors(t *testing.T) {
	chips, err := GetChips()
	if err != nil {
		t.Fatal(err)
	}

	if chips == nil {
		t.Fatal("chips is nil")
	}

	if len(chips) == 0 {
		t.Fatal("not getting chips")
	}
}
