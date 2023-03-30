package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref, Next: nil}
	return SortListInsertNode(l, newNode)
}

func SortListInsertNode(l *NodeI, newNode *NodeI) *NodeI {
	newNode.Next = nil
	for current, prev := l, l; current != nil; prev, current = current, current.Next {
		if current.Data > newNode.Data {
			newNode.Next = current
			if prev == current {
				return newNode
			} else {
				prev.Next = newNode
				break
			}
		} else if current.Next == nil {
			current.Next = newNode
			break
		}
	}
	return l
}
