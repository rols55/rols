package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		var r []int
		for k := 0; 0 != n; k++ {
			r = append(r, n%10)
			n = n / 10
		}
		len := len(r)
		count := 0
		for j := 0; count < len-1; j++ {
			count = 0
			for i := 0; i < len-1; i++ {
				if r[i] > r[i+1] {
					r[i], r[i+1] = r[i+1], r[i]
				} else {
					count++
				}
			}

		}

		for o := 0; o <= len-1; o++ {
			z01.PrintRune(rune(r[o] + 48)) // ascii number starts from 48
		}
	}
}
