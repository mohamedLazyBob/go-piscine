package piscine

func ListForEachIf(l *List, f func(*NodeL), condition func(*NodeL) bool) {
	node := l.Head
	for node != nil {
		if condition(node) {
			f(node)
		}
		node = node.Next
	}
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}
