package piscine

func LastRune(s string) rune {
	word := []rune(s)
	lng := stringsLen(s)
	return word[lng-1]
}

func stringsLen(s string) int {
	length := 0
	for range s {
		length++
	}
	return length
}
