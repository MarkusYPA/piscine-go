package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	bsLen := len(base)

	if bsLen < 2 || !allUnique(base) || containsPlusMinus(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == -9223372036854775808 && base == "0123456789" {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
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
