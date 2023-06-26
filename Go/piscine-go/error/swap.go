package piscine

func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
} // a,b = b,a
