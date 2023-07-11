package temps

import (
	"regexp"
	"strconv"
	"strings"
)

// StrHasTemp check valid sensor string
func StrHasTemp(str string) bool {
	return strings.Contains(str, "°C")
}

// StrExtractTemps extract temp slice from sensor text
func StrExtractTemps(str string) []float32 {
	re := regexp.MustCompile(`\+([0-9.]+)°C`)
	matches := re.FindAllStringSubmatch(str, -1)
	out := make([]float32, len(matches))

	for i, v := range matches {
		vv, err := strconv.ParseFloat(v[1], 32)
		if err != nil {
			continue
		}

		out[i] = float32(vv)
	}

	return out
}
