package piscine

func StrLen(s string) int {
	count := 0
	pik := []rune(s)
	for i := range pik {
		count = i + 1
	}
	return count
}
