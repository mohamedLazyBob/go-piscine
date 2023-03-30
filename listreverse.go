package piscine

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	l.Tail = l.Head
	current := l.Head
	var prev *NodeL
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
	l.Tail.Next = nil
}
