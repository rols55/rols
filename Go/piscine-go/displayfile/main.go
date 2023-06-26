package main

import (
	"fmt"
	"os"
)

func main() {
	lenght := len([]string(os.Args))
	if lenght < 2 {
		fmt.Println("File name missing")
	} else if lenght == 2 {
		filename, _ := os.ReadFile(os.Args[1]) // For read access.
		fmt.Print(string(filename))
	} else if lenght >= 3 {
		fmt.Println("Too many arguments")
	}
}
