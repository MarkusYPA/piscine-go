package piscine

func TrimAtoi(s string) int {
	numRunes := []rune{}

	for _, v := range s {
		if (v >= '0' && v <= '9') || v == '-' {
			numRunes = append(numRunes, v)
		}
	}

	return Atoi(string(numRunes))
}
