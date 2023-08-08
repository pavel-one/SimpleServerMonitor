package memory

import "testing"

func TestGetStats(t *testing.T) {
	stats, err := GetStats()
	if err != nil {
		t.Fatal(err)
	}

	if stats.Memory == nil {
		t.Fatal("memory not load data")
	}

	if stats.Memory.Total == 0 {
		t.Fatal("total memory not loaded")
	}
}
