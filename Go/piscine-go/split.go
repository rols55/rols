package piscine

func Split(s, charset string) []string {
	ch := len(charset)
	sl := len(s)
	size := 0
	for ind := 0; ind <= sl-ch; ind++ {
		if string(s[ind:ind+ch]) == charset {
			size++
		}
	}
	resArr := make([]string, size+1)
	i := 0
	start := 0
	ind := 0
	for ; ind <= sl-ch; ind++ {
		if string(s[ind:ind+ch]) == charset {
			resArr[i] = string(s[start:ind])
			i++
			start = ind + ch
		}
		if ind == sl-ch {
			resArr[i] = string(s[start:])
		}
	}
	return resArr
}
