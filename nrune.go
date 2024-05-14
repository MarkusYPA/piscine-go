package piscine

func NRune(s string, n int) rune {
	word := []rune(s)
	lng := StrLen(s)

	if n < 0 || n >= lng {
		return 0
	}

	return word[n]
}
