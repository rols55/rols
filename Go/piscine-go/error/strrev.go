package piscine

func StrRev(s string) string {
	kak := []rune(s)
	pik := len(s)
	for i, j := 0, pik-1; i < j; i, j = i+1, j-1 {
		kak[i], kak[j] = kak[j], kak[i]
	}
	s = string(kak)
	return s
}
