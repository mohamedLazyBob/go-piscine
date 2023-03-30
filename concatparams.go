package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, param := range args {
		str = Concat(str, param)
		if i != len(args)-1 {
			str = Concat(str, "\n")
		}
	}
	return str
}
