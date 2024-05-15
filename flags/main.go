package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []string(os.Args[1:])

	if len(os.Args) < 2 || help(args) {
		printHelp()
		return
	}

	lenA := len(args)
	doInsert, indexI, lenI := insert(args)
	doOrder := order(args)

	if doOrder && doInsert {
		toInsert := args[indexI][lenI:]
		output := args[lenA-1] + toInsert
		output = sortString(output)
		printString(output)
		z01.PrintRune('\n')
		return
	}

	if doOrder {
		output := sortString(args[lenA-1])
		printString(output)
		z01.PrintRune('\n')
		return
	}

	if doInsert {
		args[indexI] = args[indexI][lenI:]
		output := args[lenA-1] + args[indexI]
		printString(output)
		z01.PrintRune('\n')
		return
	}

	printString(os.Args[1])
	z01.PrintRune('\n')
}

func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}

// Sorts strings by ascii, inside and out
func sortString(str string) string {
	runes := []rune(str)
	for i := 0; i < len(runes)-1; i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

// Returns if insert flag was found, index of flag and flag length
func insert(strs []string) (bool, int, int) {
	for i, str := range strs {
		comp := "--insert="
		if len(str) >= len(comp) {
			if str[:len(comp)] == comp {
				return true, i, 9
			}
		}
	}
	for i, str := range strs {
		comp := "-i="
		if len(str) >= len(comp) {
			if str[:len(comp)] == comp {
				return true, i, 3
			}
		}
	}
	return false, 0, 0
}

// Returns if order flag was found, index of flag
func order(strs []string) bool {
	for _, str := range strs {
		comp := "--order"
		if len(str) >= len(comp) {
			if str[:len(comp)] == comp {
				return true
			}
		}
	}
	for _, str := range strs {
		comp := "-o"
		if len(str) >= len(comp) {
			if str[:len(comp)] == comp {
				return true
			}
		}
	}
	return false
}

// Return if help flag was found
func help(strs []string) bool {
	for _, str := range strs {
		if str == "--help" || str == "-h" {
			return true
		}
	}
	return false
}

func printHelp() {
	printString("--insert")
	z01.PrintRune('\n')
	printString("  -i")
	z01.PrintRune('\n')
	z01.PrintRune(9)
	printString(" This flag inserts the string into the string passed as argument.")
	z01.PrintRune('\n')
	printString("--order")
	z01.PrintRune('\n')
	printString("  -o")
	z01.PrintRune('\n')
	z01.PrintRune(9)
	printString(" This flag will behave like a boolean, if it is called it will order the argument.")
	z01.PrintRune('\n')
}
