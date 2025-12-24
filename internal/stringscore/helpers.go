package stringscore

import (
	"math"
)

func floatSlicesEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		e := math.Abs(a[i] - b[i])
		if e > 0.001 {
			return false
		}
	}
	return true
}
