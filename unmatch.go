package piscine

func Unmatch(a []int) int {
	for _, v1 := range a {
		mathces := 0
		for _, v2 := range a {
			if v1 == v2 {
				mathces++
			}
		}
		if mathces%2 == 1 {
			return v1
		}
	}
	return -1
}
