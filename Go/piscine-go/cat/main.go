package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	files := os.Args[1:]
	for _, v := range files {
		file, err := os.ReadFile(v)
		printer(file)
		if err != nil {
			z01.PrintRune('E')
			z01.PrintRune('R')
			z01.PrintRune('R')
			z01.PrintRune('O')
			z01.PrintRune('R')
			z01.PrintRune(':')
			z01.PrintRune(' ')
			for _, v := range err.Error() {
				z01.PrintRune(rune(v))
			}
			z01.PrintRune(10)
			os.Exit(1)
		}
	}
	if len(os.Args) == 1 {
		reader, _ := ioutil.ReadAll(os.Stdin)
		for _, v := range reader {
			z01.PrintRune(rune(v))
		}
	}
}

func printer(contents []byte) {
	for _, v := range contents {
		z01.PrintRune(rune(v))
	}
}
