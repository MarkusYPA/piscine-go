package piscine

func IntToString(i int) string {
	out := ""
	minus := ""
	if i < 0 {
		if i == -9223372036854775808 {
			return "-9223372036854775808"
		}
		i *= -1
		minus = "-"
	}
	if i == 0 {
		return "0"
	}
	for i > 0 {
		r := rune(i%10 + 48)
		out = string(r) + out
		i /= 10
	}
	return minus + out
}

func StringToInt(s string) int {
	num := 0
	minus := 1
	mult := 1
	valid := true
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			minus = -1
		}
		s = s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		if !(s[i] >= '0' && s[i] <= '9') {
			valid = false
			break
		} else {
			n := int(s[i] - 48) // Start from last digit
			num += n * mult
			mult *= 10
		}
	}
	if !valid {
		return -1
	} else {
		return num * minus
	}
}
