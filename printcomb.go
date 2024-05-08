package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	bn := 48 // base number for printing digits

	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				z01.PrintRune(rune(i + bn))
				z01.PrintRune(rune(j + bn))
				z01.PrintRune(rune(k + bn))

				if i < 7 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}

	z01.PrintRune('\n')
}
