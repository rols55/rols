package piscine

func TrimAtoi(s string) int {
	r := []rune(s)
	answer := 0
	for i := 0; i < len(r); i++ {
		if r[i] >= '0' && r[i] <= '9' {
			answer *= 10
			answer += int(r[i]) - '0'
		}
	}

	if HasMinus(s) {
		answer *= -1
	}
	return answer
}

func HasMinus(is string) bool {
	has := false
	for _, ch := range is {
		if ch >= 48 && ch <= 57 { // 0-9
			has = true
		} else if ch == 45 { // -

			if has == false {
				return true
			}
		}
	}
	return false
}
