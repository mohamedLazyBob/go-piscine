package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := range a {
		for j := range a {
			if f(string(a[i]), string(a[j])) == -1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
