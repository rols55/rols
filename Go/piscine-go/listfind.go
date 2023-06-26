package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for l.Head != nil {
		if comp(ref, l.Head.Data) {
			return &l.Head.Data
		}
		l.Head = l.Head.Next
	}
	return nil
}
