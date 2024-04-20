package piscine

func MakeRange(min, max int) []int {
	var s []int
	if min < max {
		s = make([]int, max-min)
		for i := range s { // dafuq is this
			s[i] = i + min
		}
	}
	return s
}
