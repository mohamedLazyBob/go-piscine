package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i := 0; i < len(strs); i++ {
		str = Concat(str, strs[i])
		if i != len(strs)-1 {
			str = Concat(str, sep)
		}
	}
	return str
}
