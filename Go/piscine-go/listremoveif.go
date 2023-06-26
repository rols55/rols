package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	n := l.Head
	previous := l.Head
	for n != nil && n.Data == data_ref {
		l.Head = n.Next
		n = l.Head
	}
	for n != nil {
		for n != nil && n.Data != data_ref {
			previous = n
			n = n.Next
		}
		if n == nil {
			return
		}
		previous.Next = n.Next
		n = previous.Next
	}
}
