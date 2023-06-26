package piscine

func UltimateDivMod(a *int, b *int) {
	var c int = *a / *b
	var d int = *a % *b
	*a = c
	*b = d
}
