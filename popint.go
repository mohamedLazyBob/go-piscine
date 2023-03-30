package piscine

func PopInt(integers []int) []int {
	return integers[:len(integers)-1]
}
