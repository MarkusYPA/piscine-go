package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := []string(os.Args[1:])
	for _, v := range params {
		for _, argChar := range v {
			z01.PrintRune(argChar)
		}
		z01.PrintRune('\n')
	}
}
