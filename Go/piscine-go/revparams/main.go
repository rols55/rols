package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []string(os.Args)
	for s := len(arg) - 1; s > 0; s-- {
		inner := []rune(arg[s])
		for i := 0; i < len(inner); i++ {
			z01.PrintRune(inner[i])
		}
		z01.PrintRune('\n')
	}
}
