package piscine

func DescendAppendRange(max, min int) []int {
	descRange := []int{}
	for i := max; i > min; i-- {
		descRange = append(descRange, i)
	}
	return descRange
}
