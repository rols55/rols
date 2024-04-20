package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() { // makes array for args
	argsSorted := []string(os.Args[1:])      // appends string to array
	for i := 0; i < len(argsSorted)-1; i++ { // sorts array
		for j := 0; j < len(argsSorted)-1-i; j++ {
			a := []rune(argsSorted[j])
			b := []rune(argsSorted[j+1])
			if a[0] > b[0] {
				argsSorted[j], argsSorted[j+1] = argsSorted[j+1], argsSorted[j]
			}
		}
	}
	for i := range argsSorted { // converts to runes and prints
		kak := []rune(argsSorted[i])
		for s := range kak {
			z01.PrintRune(rune(kak[s]))
		}
		z01.PrintRune('\n')
	}
}
