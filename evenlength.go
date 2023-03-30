package piscine

func EvenLength(strings []string) []string {
	_strings := []string(nil)
	for _, str := range strings {
		if len(str)%2 == 0 {
			_strings = append(_strings, str)
		}
	}
	return _strings
}
