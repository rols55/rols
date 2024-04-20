package piscine

func IsPrintable(s string) bool {
	r := []rune(s)
	for i := 0; i <= len(r)-1; i++ {
		if r[i] < 32 {
			return false
		}
	}
	return true
}
