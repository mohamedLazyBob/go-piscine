package piscine

func CountIf(f func(string) bool, tab []string) int {
	var c int
	for _, v := range tab {
		if f(v) {
			c++
		}
	}
	return c
}
