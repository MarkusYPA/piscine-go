package piscine

func Any(f func(string) bool, a []string) bool {
	isAnyTrue := false
	for _, str := range a {
		if f(str) {
			isAnyTrue = true
		}
		if isAnyTrue {
			break
		}
	}
	return isAnyTrue
}
