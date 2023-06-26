package piscine

func BasicAtoi2(s string) int {
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
