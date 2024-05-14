package piscine

func IsAlpha(s string) bool {
	notAlphaNum := false
	for _, v := range s {
		if !((v >= '0' && v <= '9') || (v >= 'A' && v <= 'Z') || (v >= 'a' && v <= 'z')) {
			notAlphaNum = true
		}
	}
	return !notAlphaNum
}
