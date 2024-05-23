package main

import "os"

func main() {
	args := os.Args[1:]

	for _, a := range args {
		if a == "01" || a == "galaxy" || a == "galaxy 01" {
			os.Stdout.WriteString("Alert!!!\n")
			break
		}
	}
}
