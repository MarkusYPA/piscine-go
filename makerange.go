package piscine

func MakeRange(min, max int) []int {
	length := max - min

	if length <= 0 {
		var empty []int
		return empty
	}

	ints := make([]int, length)

	for i := 0; i < length; i++ {
		ints[i] = i
	}

	return ints
}
