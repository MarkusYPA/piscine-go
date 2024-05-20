package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for _, arg := range args {

		/* 		if i != 0 {
			printStr("")
		} */

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
		arr = append(arr, byte('\n'))
		file.Read(arr)
		os.Stdout.Write(arr)
		// printStr("")
		// os.Stdout.Write([]byte("\n"))
		// os.Stdout.Sync()
	}
	// printStr("")
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
