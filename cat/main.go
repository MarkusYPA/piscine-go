package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {

		file, err := os.Open(arg)
		if err != nil {
			printStr("ERROR: " + err.Error())
			os.Exit(0)
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			printStr("ERROR: " + err.Error())
			os.Exit(0)
		}

		arr := make([]byte, fileInfo.Size())
		// arr = append(arr, byte('\n'))
		file.Read(arr)
		os.Stdout.Write(arr)
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
