package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 || nb == 1 {
		return nb
	}

	result := powers(nb)

	if result*result == nb {
		return result
	} else {
		return 0
	}
}

func powers(nb int) int {
	powa := 1

	for i := 1; i < nb; i++ {
		powa = i * i

		if powa >= nb {
			return i
		}
	}

	return 0
}
