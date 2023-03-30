package piscine

func BasicJoin(elements []string) string {
	str := ""
	for i := 0; i < len(elements); i++ {
		str = Concat(str, elements[i])
	}
	return str
}
