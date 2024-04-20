package piscine

func NRune(a string, s int) rune {
	if s > 0 && s <= len(a) {
		r := []rune(a)
		return (r[s-1])
	} else {
		return 0
	}
}
