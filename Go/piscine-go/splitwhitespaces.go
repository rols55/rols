package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	word := ""
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if len(word) > 0 {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	result = append(result, word)
	return result
}
