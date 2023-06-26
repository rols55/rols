package piscine

func FirstRune(s string) rune {
	r := []rune(s)
	for _, c := range r {
		return c
	}
	return 0
}
