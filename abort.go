package piscine

func Abort(a, b, c, d, e int) int {
	ints := []int{a, b, c, d, e}
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			if ints[i] > ints[j] {
				ints[i], ints[j] = ints[j], ints[i]
			}
		}
	}
	return ints[2]
}
