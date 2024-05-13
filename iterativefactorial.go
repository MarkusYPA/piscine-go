package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 25 {
		return 0
	}

	fctrl := 1
	for i := 1; i <= nb; i++ {
		fctrl *= i
	}

	return fctrl
}
