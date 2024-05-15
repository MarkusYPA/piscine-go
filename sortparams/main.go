package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := []string(os.Args[1:])

	for i := 0; i < len(params)-1; i++ {
		for j := i + 1; j < len(params); j++ {
			if params[i] > params[j] {
				params[i], params[j] = params[j], params[i]
			}
		}
	}

	for _, v := range params {
		for _, char := range v {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
