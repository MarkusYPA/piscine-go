package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {

	for i := 0; i < 26; i++ {
		fmt.Printf("%c", 97+i)
	}
	z01.PrintRune('\n')
}
