package piscine

func Map(f func(int) bool, a []int) []bool {
	var answer []bool
	for _, numbers := range a {
		answer = append(answer, f(numbers))
	}
	return answer
}
