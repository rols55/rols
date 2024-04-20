package piscine

import "fmt"

func QuadE(x, y int) {
	for h := 1; h <= y; h++ {
		for w := 1; w <= x; w++ {
			if h == 1 && w == 1 || h == y && w == x && y != 1 && x != 1 {
				fmt.Print("A")
			} else if h == 1 && w == x || h == y && w == 1 {
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
