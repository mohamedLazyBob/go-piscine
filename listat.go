package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	index := 0
	for l != nil {
		if index == pos {
			return l
		}
		index++
		l = l.Next
	}
	return nil
}
