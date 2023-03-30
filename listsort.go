package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	for current := l; current != nil; current = current.Next {
		for index := current.Next; index != nil; index = index.Next {
			if current.Data > index.Data {
				current.Data, index.Data = index.Data, current.Data
			}
		}
	}
	return l
}
