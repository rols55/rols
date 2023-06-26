package piscine

func ListReverse(l *List) {
	newList := &List{}
	for l.Head != nil {
		ListPushFront(newList, l.Head.Data)
		l.Head = l.Head.Next
	}
	*l = *newList
}
