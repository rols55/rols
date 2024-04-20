package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, value := range tab {
		if f(value) {
			counter++
		}
	}
	return counter
}
