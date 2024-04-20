package piscine

import "fmt"

func QuadC(x, y int) {
	for h := 1; h <= y; h++ {
		for w := 1; w <= x; w++ {
			if h == 1 && w == 1 || h == 1 && w == x {
				fmt.Print("A")
			} else if h == y && w == x || h == y && w == 1 {
				fmt.Print("C")
			} else if w == 1 || w == x || w > 1 && h == 1 || w > 1 && h == y {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
