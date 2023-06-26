package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for one := '0'; one <= '9'; one++ {
		for two := '0'; two <= '9'; two++ {
			four := two + 1
			for three := one; three <= '9'; three++ {
				for ; four <= '9'; four++ {
					z01.PrintRune(one)
					z01.PrintRune(two)
					z01.PrintRune(' ')
					z01.PrintRune(three)
					z01.PrintRune(four)
					if one < '9' || two < '8' || three < '9' || four < '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				four = '0'
			}

		}
	}
	z01.PrintRune('\n')
}
