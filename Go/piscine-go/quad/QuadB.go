package piscine

import "fmt"

func QuadB(x, y int) {
	for h := 1; h <= y; h++ {
		for w := 1; w <= x; w++ {
			if w == 1 && h == 1 || h != 1 && h == y && w == x && w != 1 {
				fmt.Print("/")
			} else if w == x && h == 1 || h == y && w == 1 {
				fmt.Print("\\")
			} else if h > 1 && w == 1 || h > 1 && w == x {
				fmt.Print("*")
			} else if h == 1 && w > 1 || h == y && w > 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
