package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	n1 = ListSort(n1)
	head := n1
	for current := n2; current != nil; {
		next := current.Next
		head = SortListInsertNode(head, current)
		current = next
	}
	return head
}
