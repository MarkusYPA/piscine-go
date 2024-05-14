package piscine

func ToUpper(s string) string {
	runes := []rune(s)

	for i := 0; i < StrLen(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			runes[i] = rune(s[i]) - 32
		}
	}

	return string(runes)
}
