package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	var items string
	for _, e := range str {
		if e == 32 {
			list[items] += 1
			items = ""
		} else if e != 32 {
			items += string(byte(e))
		}
	}
	list[items] += 1

	return list
}
