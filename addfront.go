package piscine

func AddFront(s string, slice []string) []string {
	_slice := []string{s}
	_slice = append(_slice, slice...)
	return _slice
}
