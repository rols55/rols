package piscine

func ListPushFront(l *List, data interface{}) {
	x := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = x
	} else {
		x.Next = l.Head
		l.Head = x

	}
}
