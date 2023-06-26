package piscine

func Capitalize(s string) string {
	r := []rune(s)

	for i, e := range r {
		if IsAlphaNumeric(e) {
			if i == 0 || IsAlphaNumeric(r[i-1]) == false {
				if r[i] >= 'a' && r[i] <= 'z' {
					r[i] = e - 32
				}
			} else {
				if r[i] >= 'A' && r[i] <= 'Z' {
					r[i] = e + 32
				}
			}
		}
	}
	return string(r)
}

func IsAlphaNumeric(r rune) bool {
	if (r < 'A' || r > 'Z') &&
		(r < 'a' || r > 'z') &&
		(r < '0' || r > '9') {
		return false
	}
	return true
}
