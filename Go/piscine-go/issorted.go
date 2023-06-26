package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	sortUp := true
	sortDown := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			sortUp = false
		}
		if f(a[i], a[i+1]) < 0 {
			sortDown = false
		}
	}
	return sortUp || sortDown
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
