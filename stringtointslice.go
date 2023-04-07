package piscine

func StringToIntSlice(s string) []int {
	var sslice []int
	for i, _ := range s {
		sslice = append(sslice, int(s[i]))
	}
	return sslice
}
