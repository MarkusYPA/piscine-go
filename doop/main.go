package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		os.Exit(0)
	}
	args := os.Args[1:]
	operator := args[1]
	operand1, err1 := stringToInt(args[0])
	operand2, err2 := stringToInt(args[2])

	if err1 != "" || err2 != "" {
		return
		// os.Exit(0)
	}

	checkOverflow(operand1, operand2)

	funcs := []func(int, int) int{plus, minus, divi, multi, modulo}

	switch operator {
	case "+":
		printToStdout(funcs[0](operand1, operand2))
	case "-":
		printToStdout(funcs[1](operand1, operand2))
	case "/":
		if operand2 == 0 {
			os.Stdout.WriteString("No division by 0") // Without line change
			return
		}
		printToStdout(funcs[2](operand1, operand2))
	case "*":
		printToStdout(funcs[3](operand1, operand2))
	case "%":
		if operand2 == 0 {
			os.Stdout.WriteString("No modulo by 0") // Without line change
			return
		}
		printToStdout(funcs[4](operand1, operand2))
	default:
		return
		// os.Exit(0)
	}
}

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func divi(a, b int) int {
	return a / b
}

func multi(a, b int) int {
	return a * b
}

func modulo(a, b int) int {
	return a % b
}

func printToStdout(i int) {
	str := intToString(i)
	os.Stdout.WriteString(str)
	os.Stdout.WriteString("\n")
}

func intToString(i int) string {
	out := ""
	minus := ""
	if i < 0 {
		if i == -9223372036854775808 {
			return "-9223372036854775808"
		}
		i *= -1
		minus = "-"
	}
	if i == 0 {
		return "0"
	}
	for i > 0 {
		r := rune(i%10 + 48)
		out = string(r) + out
		i /= 10
	}
	return minus + out
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
		os.Exit(0)
	}
	if b < -922337203685477580 || b > 922337203685477580 {
		os.Exit(0)
	}
}
