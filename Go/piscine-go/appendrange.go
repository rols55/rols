package piscine

func AppendRange(min, max int) []int {
	if min > max {
		var s []int
		return s
	} else {
		var s []int
		for i := min; i < max; i++ {
			s = append(s, i)
		}
		return s
	}
}
