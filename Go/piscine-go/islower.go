package piscine

func IsLower(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}
