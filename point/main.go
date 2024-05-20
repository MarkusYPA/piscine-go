package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	xString := intToString(points.x)
	yString := intToString(points.y)

	printStr("x = " + xString + ", " + "y = " + yString)
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func intToString(i int) string {
	out := ""
	for i > 0 {
		r := rune(i%10 + 48)
		out = string(r) + out
		i /= 10
	}
	return out
}
