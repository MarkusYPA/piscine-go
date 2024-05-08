package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	}

	runes := []rune{}
	bn := 48

	maxfile := false

	if n == -9223372036854775808 {
		n = 922337203685477580
		maxfile = true
	}

	for n > 0 {
		runes = append(runes, rune((n%10)+bn))
		n = n / 10
	}

	for i := len(runes) - 1; i > -1; i-- {
		z01.PrintRune(runes[i])
	}

	if maxfile {
		z01.PrintRune('8')
	}
}
