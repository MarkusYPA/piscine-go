package piscine

func BasicAtoi(s string) int {
	lng := len(s)
	ints := make([]int, lng)
	result := 0
	multiplier := 1

	for i, v := range s {
		ints[i] = int(v - 48)
	}

	for i := lng - 1; i > -1; i-- {
		result += ints[i] * multiplier
		multiplier *= 10
	}

	return result
}

func stringLength(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}
