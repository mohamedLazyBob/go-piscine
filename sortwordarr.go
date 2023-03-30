package piscine

func SortWordArr(a []string) {
	for i := range a {
		for j := range a {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
