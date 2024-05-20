package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		os.Exit(0)
	}

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
	}

	arr := make([]byte, fileInfo.Size())
	file.Read(arr)
	fmt.Println(string(arr))
}
