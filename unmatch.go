package piscine

func Unmatch(a []int) int {
	for i := range a {
		isFound := false
		for j := range a {
			if i != j && a[i] == a[j] {
				isFound = true
				a[i] = -1
				a[j] = -1
				break
			}
		}
		if !isFound {
			return a[i]
		}
	}
	return -1
}
