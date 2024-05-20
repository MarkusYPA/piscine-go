package main

import (
	"fmt"
	"os"
)

func main() {
	exitNro := 0
	args := os.Args[1:]
	cut := 0

	if len(args) > 1 {
		cut = stringToInt(args[1])
	}

	if len(args) > 2 && args[0] == "-c" {
		for i, filename := range args[2:] {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Println(err.Error())
				exitNro = 1
			} else {
				if i != 0 {
					fmt.Println()
				}
				if len(args[2:]) > 1 {
					fmt.Println("==>", filename, "<==")
				}
				fileInfo, _ := file.Stat()
				arr := make([]byte, fileInfo.Size())
				file.Read(arr)
				if cut > len(arr) {
					cut = len(arr)
				}
				fmt.Print(string(arr[len(arr)-cut:]))
			}
			defer file.Close()
		}
	} else {
		exitNro = 1
	}
	os.Exit(exitNro)
}

func stringToInt(s string) int {
	num := 0
	goOn := true
	revInd := 0
	for i := len(s) - 1; i >= 0; i-- {
		if !(s[i] >= '0' && s[i] <= '9') {
			goOn = false
			break
		} else {
			n := int(s[i] - 48) // Start from last digit
			num += n * power(10, revInd)
		}
		revInd++
	}
	if !goOn {
		return -1
	} else {
		return num
	}
}

// Returns n to the power p
func power(n, p int) int {
	if p == 0 {
		return 1
	}
	if p == 1 {
		return n
	}
	ns := n
	for i := 1; i < p; i++ {
		n *= ns
	}
	return n
}
