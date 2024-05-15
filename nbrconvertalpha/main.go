package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	capitals := 0
	skip1 := 0

	if os.Args[1] == "--upper" {
		capitals = -32
		skip1 = 1
	}

	for i, v := range os.Args {
		if i > 0+skip1 {
			num := piscine.Atoi(v) - 1
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
