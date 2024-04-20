package piscine

// import "fmt"

func BasicAtoi(s string) int {
	pik := []rune(s)
	pikkus := len(pik)
	answer := 0
	for i := 0; i < pikkus; i++ {
		if pik[i] < '0' || pik[i] > '9' {
			return 0
		} else {
			answer *= 10
			answer += int(pik[i]) - '0'
		}
	}
	return answer
}

/*func main() {
	fmt.Println(BasicAtoi2("12345000"))
	fmt.Println(BasicAtoi2("76570000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}
*/
