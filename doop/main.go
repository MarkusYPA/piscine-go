package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		os.Exit(1)
	}
	args := os.Args[1:]
	operator := args[1]
	operand1, err1 := stringToInt(args[0])
	operand2, err2 := stringToInt(args[2])

	if err1 != "" || err2 != "" {
		os.Exit(1)
	}

	checkOverflow(operand1, operand2)

	funcs := []func(int, int) int{plus, minus, divi, multi, modulo}

	switch operator {
	case "+":
		println(funcs[0](operand1, operand2))
	case "-":
		println(funcs[1](operand1, operand2))
	case "/":
		println(funcs[2](operand1, operand2))
	case "*":
		println(funcs[3](operand1, operand2))
	case "%":
		println(funcs[4](operand1, operand2))
	default:
		os.Exit(1)
	}
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func divi(a, b int) int {
	if b == 0 {
		println("No division by 0")
		os.Exit(1)
	}
	return a / b
}

func multi(a, b int) int {
	return a * b
}

func modulo(a, b int) int {
	if b == 0 {
		println("No modulo by 0")
		os.Exit(1)
	}
	return a % b
}

func stringToInt(s string) (int, string) {
	num := 0
	minus := 1
	mult := 1
	valid := true
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			minus = -1
		}
		s = s[1:]
	}
	for i := len(s) - 1; i >= 0; i-- {
		if !(s[i] >= '0' && s[i] <= '9') {
			valid = false
			break
		} else {
			n := int(s[i] - 48) // Start from last digit
			num += n * mult
			mult *= 10
		}
	}
	if !valid {
		return -1, "err"
	} else {
		return num * minus, ""
	}
}

func checkOverflow(a, b int) {
	if a < -922337203685477580 || a > 922337203685477580 { // minimum and maximun values without last digit
		os.Exit(1)
	}
	if b < -922337203685477580 || b > 922337203685477580 {
		os.Exit(1)
	}
}
