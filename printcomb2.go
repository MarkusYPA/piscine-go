package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {

	bn := 48 // base number for printing digits

	for i := 0; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			z01.PrintRune(rune((i / 10) + bn))
			z01.PrintRune(rune((i % 10) + bn))
			z01.PrintRune(' ')

			z01.PrintRune(rune((j / 10) + bn))
			z01.PrintRune(rune((j % 10) + bn))

			if i < 98 || j < 98 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
