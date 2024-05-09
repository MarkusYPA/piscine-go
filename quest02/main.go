package main

import (
	"fmt"

	"piscine/quest03/piscine"
	//"github.com/01-edu/z01"
)

func main() {
	/*
		piscine.IsNegative(-5)
		piscine.IsNegative(0)
		piscine.IsNegative(5)
	*/

	// piscine.PrintComb()
	// piscine.PrintComb2()

	// piscine.PrintNbr(-123)
	// piscine.PrintNbr(123456)
	// piscine.PrintNbr(0)
	// piscine.PrintNbr(123)
	// piscine.PrintNbr(-9223372036854775807)
	// piscine.PrintNbr(-9223372036854775808)

	piscine.PrintCombN(1)
	piscine.PrintCombN(3)
	piscine.PrintCombN(9)

	// test()
}

func test() {
	// bn := 48 // base number for printing digits

	for i := 0; i < 100; i++ {
		// z01.PrintRune(rune(i + bn))
		fmt.Print(i / 10)
		fmt.Print(i % 10)
		fmt.Println()
	}
}
