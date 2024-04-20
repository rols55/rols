package piscine

func Any(f func(string) bool, a []string) bool {
	for _, numbers := range a {
		if f(numbers) {
			return true
		}
	}
	return false
}
