package piscine

func Atoi(s string) int {

	if len(s) == 0 {
		return 0
	}

	factor := 1

	if s[0] == '+' || s[0] == '-' {

		if s[0] == '-' {
			factor = -1
		}

		s = s[1:]
	}

	ints := make([]int, len(s))

	for i, v := range s {
		if v < 48 || v > 57 {
			return 0
		}
		ints[i] = int(v - 48)
	}

	result := 0
	mult := 1

	for i := len(s) - 1; i > -1; i-- {
		result += ints[i] * mult
		mult *= 10
	}

	return result * factor

}
