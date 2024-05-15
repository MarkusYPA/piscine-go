package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := []string(os.Args[1:])

	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	allArguments := joinThese(args, " ")
	vowelsInArgs := []rune{}

	for _, v := range allArguments {
		if isVowel(v) {
			vowelsInArgs = append(vowelsInArgs, v)
		}
	}

	if len(vowelsInArgs) == 0 {
		printString(allArguments)
		z01.PrintRune('\n')
		return
	}

	reverseRunes(vowelsInArgs)
	vowelIndex := 0

	output := make([]rune, len(allArguments))
	for i, v := range allArguments {
		if isVowel(v) {
			output[i] = vowelsInArgs[vowelIndex]
			vowelIndex++
		} else {
			output[i] = v
		}
	}

	printString(string(output))
	z01.PrintRune('\n')
}

// Reverse the order of a slice of runes
func reverseRunes(rns []rune) {
	for i := 0; i < (len(rns) / 2); i++ {
		rns[i], rns[len(rns)-1-i] = rns[len(rns)-1-i], rns[i]
	}
}

// Join a slice of strings, separate by sep
func joinThese(strs []string, sep string) string {
	result := ""
	for i, v := range strs {
		result += v
		if i < len(strs)-1 {
			result += sep
		}
	}
	return result
}

// See if rune is a vowel
func isVowel(r rune) bool {
	allVowels := []rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
	for _, v := range allVowels {
		if r == v {
			return true
		}
	}
	return false
}

// Prints a string
func printString(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
