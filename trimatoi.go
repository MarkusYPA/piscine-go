package piscine

func TrimAtoi(s string) int {
	numRunes := []rune{}
	numsFound := false

	for _, v := range s {
		if v >= '0' && v <= '9' {
			numRunes = append(numRunes, v)
			numsFound = true
		}

		if v == '-' && !numsFound {
			numRunes = append(numRunes, v)
		}
	}

	return Atoi(string(numRunes))
}
