package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	// strl := []rune(s) //not needed
	for _, sec := range s {
		z01.PrintRune(sec)
	}
	z01.PrintRune('\n')
}

/*func main(){
	PrintStr(Hello!)
}*/
