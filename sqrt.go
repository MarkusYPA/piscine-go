package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 || nb == 1 {
		return nb
	}

	result := sqrtInt(nb)

	if result*result == nb {
		return result
	} else {
		return 0
	}
}

// Returns the equal or next integer of the square root of nb
func sqrtInt(nb int) int {
	powa := 1

	for i := 1; i < nb; i++ {
		powa = i * i

		if powa >= nb {
			return i
		}
	}

	return 0
}
