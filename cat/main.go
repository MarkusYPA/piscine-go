package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		// If no command-line argument is provided, read from stdin
		input := make([]byte, 1024)
		for {
			numberOfBytesRead, err := os.Stdin.Read(input) // writing Stdin to 'input'
			if err != nil {
				os.Exit(0) // Error: EOF > exit, In second round when no bytes have been read?
			}

			toPrint := string(input[:numberOfBytesRead])
			printStr(toPrint)

			input = make([]byte, 1024) // reset slice
		}
	} else {
		for _, arg := range args {

			file, err := os.Open(arg)
			if err != nil {
				printStr("ERROR: " + err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}
			defer file.Close()

			fileInfo, err := file.Stat()
			if err != nil {
				printStr("ERROR: " + err.Error())
				z01.PrintRune('\n')
				os.Exit(1)
			}

			arr := make([]byte, fileInfo.Size())
			file.Read(arr)
			os.Stdout.Write(arr)
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
