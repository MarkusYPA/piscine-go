package piscine

func IsNumeric(s string) bool {
	notNum := false
	for _, v := range s {
		if v < '0' || v > '9' {
			notNum = true
		}
	}
	return !notNum
}
