package piscine

func ForEach(f func(int), a []int) {
	for _, numbers := range a {
		f(numbers)
	}
}
