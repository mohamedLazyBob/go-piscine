package piscine

func LoafOfBread(str string) string {
	if len(str) == 0 {
		return "\n"
	} else if len(str) < 5 {
		return "Invalid Output\n"
	} else {
		skip := ""
		i := 0
		for _, e := range str {
			if i%6 != 5 && e == ' ' {
				continue
			}
			if i%6 == 5 {
				skip += " "
			} else {
				skip += string(e)
			}
			i++
		}
		// remove trailing spaces
		for i := len(skip) - 1; i >= 0; i-- {
			if skip[i] != ' ' {
				skip = skip[:i+1]
				break
			}
		}
		return skip + "\n"
	}
}
