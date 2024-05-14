package piscine

func Capitalize(s string) string {
	rns := []rune(s)
	isLetter := false
	isSmall := false
	isCapital := false
	prevIsLetter := false

	for i := 0; i < len(rns); i++ {
		if (rns[i] >= '0' && rns[i] <= '9') || (rns[i] >= 'A' && rns[i] <= 'Z') || (rns[i] >= 'a' && rns[i] <= 'z') {
			isLetter = true

			if rns[i] >= 'a' && rns[i] <= 'z' {
				isSmall = true
			}

			if rns[i] >= 'A' && rns[i] <= 'Z' {
				isCapital = true
			}
		}

		if !prevIsLetter && isSmall {
			rns[i] = rns[i] - 32
		}

		if prevIsLetter && isCapital {
			rns[i] = rns[i] + 32
		}

		prevIsLetter = isLetter
		isLetter = false
		isCapital = false
		isSmall = false
	}

	return string(rns)
}
