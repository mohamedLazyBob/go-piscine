package piscine

func Slice(a []string, nbr ...int) []string {
	lenTab := len(a)
	start := 0
	end := lenTab
	if len(nbr) >= 1 {
		start = nbr[0]
		if start < 0 {
			start = (start + lenTab) % lenTab
		}
		if len(nbr) == 2 {
			end = nbr[1]
			if end < 0 {
				end = (end + lenTab) % lenTab
			}
		}
		if start > end {
			return nil
		}
	}
	return a[start:end]
}
