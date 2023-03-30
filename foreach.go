package piscine

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
