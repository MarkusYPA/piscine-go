package piscine

func IsUpper(s string) bool {
	notUpper := false
	for _, v := range s {
		if v < 'A' || v > 'Z' {
			notUpper = true
		}
	}
	return !notUpper
}
