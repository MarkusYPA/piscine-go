package piscine

func NRune(s string, n int) rune {
	word := []rune(s)
	lng := StrLen(s)

	if n < 1 || n > lng {
		return 0
	}

	return word[n-1]
}
