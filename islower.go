package piscine

func IsLower(s string) bool {
	notLower := false
	for _, v := range s {
		if v < 'a' || v > 'z' {
			notLower = true
		}
	}
	return !notLower
}
