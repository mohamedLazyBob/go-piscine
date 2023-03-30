package piscine

func Delete(integers []int, position int) []int {
	position--
	if position < 0 || position > len(integers) {
		return integers
	}
	return append(integers[:position], integers[position+1:]...)
}
