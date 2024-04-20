package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []rune(os.Args[0])
	r := []rune(arg)
	for i := 2; i < len(arg); i++ {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}
