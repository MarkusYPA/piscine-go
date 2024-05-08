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
	tempN := n
	bn := 48

	for tempN > 0 {
		runes = append(runes, rune((tempN%10)+bn))
		tempN = tempN / 10
	}

	for i := len(runes) - 1; i > -1; i-- {
		z01.PrintRune(runes[i])
	}
}
