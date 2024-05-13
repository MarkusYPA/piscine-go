package piscine

import "math"

func Sqrt(nb int) int {
	result := math.Sqrt(float64(nb))

	if result-float64(int(result)) != 0 {
		return 0
	} else {
		return int(result)
	}
}
