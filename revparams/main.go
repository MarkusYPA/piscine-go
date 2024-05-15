package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := []string(os.Args[1:])

	for i := len(params); i > 0; i-- {
		for _, argChar := range params[i-1] {
			z01.PrintRune(argChar)
		}
		z01.PrintRune('\n')
	}
}
