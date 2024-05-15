package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	capitals := 0
	skip1 := 0

	if len(os.Args) < 2 {
		return
	}

	if os.Args[1] == "--upper" {
		capitals = -32
		skip1 = 1
	}

	for i, v := range os.Args {
		if i > 0+skip1 {
			num := atoinNow(v) - 1
			myRune := 'a' + rune(num)
			if myRune < 'a' || myRune > 'z' {
				z01.PrintRune(' ')
			} else {
				runeToPrint := myRune + rune(capitals)
				z01.PrintRune(runeToPrint)
			}
		}
	}
	z01.PrintRune('\n')
}

func atoinNow(s string) int {
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
