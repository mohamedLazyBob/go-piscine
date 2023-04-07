package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
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
		return skip + "\n"
	}
}
