package piscine

// Convert string s in base b to int
func ToDec(s, b string) int {
	num := 0
	mult := 1

	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			mult = -1
		}
		s = s[1:]
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(b); j++ {
			if s[i] == b[j] {
				toAdd := pawa(len(b), len(s)-1-i) * j
				num += toAdd
			}
		}
	}

	return num * mult
}

// Convert number n to base b
func ToBase(n int, b string) string {
	result := ""
	mult := ""
	minInt := false

	if n == -9223372036854775808 {
		n = 9223372036854775807
		mult = "-"
		minInt = true
	}

	if n < 0 {
		mult = "-"
		n *= -1
	}

	if n == 0 {
		return string(b[0])
	}

	for n > 0 {
		index := n % len(b)
		if minInt {
			index++
			if len(b) == index {
				index = 0
			}
		}
		result = string(b[index]) + result
		n /= len(b)
	}

	return mult + result
}

// Returns number n to the power p
func pawa(n, p int) int {
	res := 1

	for i := 0; i < p; i++ {
		res = res * n
	}

	return res
}
