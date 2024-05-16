package piscine

func NbrBase(nbr int, base string) string {
	bsLen := len(base)
	toReturn := ""
	minInt := false

	if bsLen < 2 || !allUnique(base) || containsPlusMinus(base) {
		return "NV"
	}

	if nbr == 0 {
		return "0"
	}

	if nbr < 0 {
		if nbr == -9223372036854775808 {
			nbr = 9223372036854775807
			minInt = true
		} else {
			nbr *= -1
		}
		toReturn += "-"
	}

	exponent := 0
	if !minInt {
		exponent = HiPower(bsLen, nbr) // Highest exponent to base length not over nbr
	} else {
		exponent = hiPowerMin(bsLen)
	}
	remains := nbr

	for i := exponent; i >= 0; i-- {
		slotValue := Power(bsLen, i)
		multiplier := remains / slotValue
		remains = remains - slotValue*multiplier

		if minInt {
			multiplier++
			if multiplier > bsLen-1 {
				multiplier = 0
			}
		}

		// How many times the slot value can be multiplied
		toReturn += string(base[multiplier])
	}
	return toReturn
}

// Returns the highest power for n that is below or equal to target
func HiPower(base, target int) int {
	calc := 1
	prevCalc := 0
	pows := 0
	for i := 0; calc <= target; i++ {
		calc *= base
		pows = i
		if prevCalc > calc {
			return pows
		}
		prevCalc = calc
	}
	return pows
}

// Returns n to the power p
func Power(n, p int) int {
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

// Are all characters in string unique
func allUnique(s string) bool {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

// Does string contain + or -
func containsPlusMinus(s string) bool {
	for _, v := range s {
		if v == 43 || v == 45 {
			return true
		}
	}
	return false
}

// Returns the highest power for base (length of base) that is below or equal to the minimun int64 value
func hiPowerMin(base int) int {
	pows := 0

	switch base {
	case 2:
		pows = 63
	case 3:
		pows = 39
	case 4:
		pows = 31
	case 5:
		pows = 27
	case 6:
		pows = 24
	case 7:
		pows = 22
	case 8:
		pows = 21
	case 9:
		pows = 19
	case 10:
		pows = 18
	case 11:
		pows = 18
	case 12:
		pows = 17
	case 13:
		pows = 17
	case 14, 15:
		pows = 16
	case 16, 17, 18:
		pows = 15
	case 19, 20, 21, 22:
		pows = 14
	case 23, 24, 25, 26, 27, 28:
		pows = 13
	case 29, 30, 31, 32, 33, 34, 35, 36, 37, 38:
		pows = 12
	default:
		pows = 11
	}

	return pows
}
