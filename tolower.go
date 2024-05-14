package piscine

func ToLower(s string) string {
	runes := []rune(s)

	for i := 0; i < StrLen(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			runes[i] = rune(s[i]) + 32
		}
	}

	return string(runes)
}
