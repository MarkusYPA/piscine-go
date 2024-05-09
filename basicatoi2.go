package piscine

func BasicAtoi2(s string) int {
	lng := len(s)
	ints := make([]int, lng)
	result := 0
	multiplier := 1

	for i, v := range s {

		if v < 48 || v > 57 {
			return 0
		}

		ints[i] = int(v - 48)
	}

	for i := lng - 1; i > -1; i-- {
		result += ints[i] * multiplier
		multiplier *= 10
	}

	return result
}
