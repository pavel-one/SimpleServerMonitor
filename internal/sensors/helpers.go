package sensors

import (
	"regexp"
	"strconv"
	"strings"
)

func StrHasTemp(str string) bool {
	if !strings.Contains(str, "°C") {
		return false
	}

	return true
}

func StrExtractTemps(str string) []float32 {
	re := regexp.MustCompile(`\+([0-9.]+)°C`)
	matches := re.FindAllStringSubmatch(str, -1)
	out := make([]float32, len(matches))

	for i, v := range matches {
		vv, err := strconv.ParseFloat(v[1], 10)
		if err != nil {
			continue
		}

		out[i] = float32(vv)
	}

	return out
}
