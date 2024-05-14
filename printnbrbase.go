package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	bsLen := len(base)

	if bsLen < 2 || !allUnique(base) || containsPlusMinus(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		z01.PrintRune('0')
		return
	}

	if nbr == 9223372036854775807 {
		z01.PrintRune('7')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		z01.PrintRune('F')
		return
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr *= -1
	}

	exponent := HiPower(bsLen, nbr)
	fmt.Println(exponent)
	slotValue := Power(bsLen, exponent)
	multiplier := nbr / slotValue
	remains := nbr - slotValue*multiplier

	// Print how many times the slot value can be multiplied
	z01.PrintRune(rune(base[multiplier]))

	// Print zeros to next digit slots if necessary
	for i := 0; i < exponent; i++ {
		nextSlot := Power(bsLen, exponent-1-i)
		if remains/nextSlot < 1 {
			z01.PrintRune(rune(base[0]))
		}
	}

	if remains == 0 {
		return
	}

	PrintNbrBase(remains, base)
}

// Returns the highest power for n that is below or equal to target
func HiPower(n, target int) int {
	calc := 1
	pows := 0
	for i := 0; calc <= target; i++ {
		calc *= n
		pows = i
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

func containsPlusMinus(s string) bool {
	for _, v := range s {
		if v == 43 || v == 45 {
			return true
		}
	}
	return false
}
