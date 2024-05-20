package main

import (
	"piscine"

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

	xString := piscine.NbrBase(points.x, "0123456789")
	yString := piscine.NbrBase(points.y, "0123456789")

	printStr("x = " + xString)
	printStr("y = " + yString)
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
