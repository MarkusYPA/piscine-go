package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}

	if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	}

	return 1 // Return 1 if nb is 0 or 1
}
