package piscine

func RecursiveFactorial(nb int) int {
	if nb >= 0 && nb <= 20 { // check if nb is negative
		if nb == 0 {
			nb = nb + 1
		}
		if nb == 1 { // calls itself until nb is 1
			return nb
		} else {
			return nb * RecursiveFactorial(nb-1)
		}
	} else {
		return 0
	}
}
