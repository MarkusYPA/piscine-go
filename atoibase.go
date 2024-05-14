package piscine

func AtoiBase(s string, base string) int {
	bsLen := len(base)
	mult := 1

	// If s starts with minus
	if s[0] == 45 {
		mult = -1
		s = s[1:]
	}

	if bsLen < 2 || !allUniqueB(base) || containsPlusMinusB(base) || !stringIsValid(s, base) {
		return 0
	}

	if s[0] == base[0] && len(s) == 1 {
		return 0
	}

	result := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				expt := len(s) - 1 - i
				result += PowerB(len(base), expt) * j
			}
		}
	}

	return result * mult
}

func stringIsValid(s, b string) bool {
	validChars := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(b); j++ {
			if s[i] == b[j] {
				validChars++
			}
		}
	}
	return validChars == len(s)
}

// Returns n to the power p
func PowerB(n, p int) int {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return n
	}
	ns := n
	for i := 1; i < p; i++ {
		n *= ns
	}
	return n
}

func allUniqueB(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func containsPlusMinusB(s string) bool {
	for _, v := range s {
		if v == 43 || v == 45 {
			return true
		}
	}
	return false
}
