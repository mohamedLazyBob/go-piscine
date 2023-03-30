package piscine

func ListSize(l *List) int {
	count := 0
	if l.Head == nil {
		return 0
	}
	actual := l.Head
	count++
	for actual.Next != nil {
		count++
		actual = actual.Next
	}
	return count
}
