package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	nums := []int{}

	if n == 0 {
		z01.PrintRune('0')
	}

	for n != 0 {
		nums = append(nums, n%10)
		n /= 10
	}

	SortIntegerTable(nums)

	for _, v := range nums {
		z01.PrintRune(rune(v + 48))
	}
}
