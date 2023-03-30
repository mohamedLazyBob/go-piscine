package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, compare func(a, b interface{}) bool) *interface{} {
	node := l.Head
	for node != nil {
		if compare(ref, node.Data) {
			return &node.Data
		}
		node = node.Next
	}
	return nil
}
