package temps

import "testing"

func TestStrExtractTemps(t *testing.T) {
	str := `+45.0°C (high = +80.0°C, crit = +100.0°C)`

	out := StrExtractTemps(str)

	if out[0] != 45.0 {
		t.Fatalf("0 index not correct: %f", out[0])
	}

	if out[1] != 80.0 {
		t.Fatalf("1 index not correct: %f", out[1])
	}

	if out[2] != 100 {
		t.Fatalf("2 index not correct: %f", out[2])
	}
}

func TestStrHasTemp(t *testing.T) {
	str := `+45.0°C (high = +80.0°C, crit = +100.0°C)`

	if !StrHasTemp(str) {
		t.Fatalf("StrHasTemp not correct: %s", str)
	}

	str = "ACPI interface"
	if StrHasTemp(str) {
		t.Fatalf("StrHasTemp not correct: %s", str)
	}

	str = `hi, its string`

	if StrHasTemp(str) {
		t.Fatalf("StrHasTemp not correct: %s", str)
	}
}
