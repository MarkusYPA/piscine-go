package piscine

func Rot14(s string) string {
	newString := make([]rune, len(s))

	for i, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			newString[i] = r + 14
		} else {
			newString[i] = r
		}

		if (r >= 'A' && r <= 'Z') && newString[i] > 'Z' {
			newString[i] -= 26
		}

		if (r >= 'a' && r <= 'z') && newString[i] > 'z' {
			newString[i] -= 26
		}
	}

	return string(newString)
}
