package piscine

func Compact(ptr *[]string) int {
	newPtr := []string(nil)
	for _, str := range *ptr {
		if str != "" {
			newPtr = append(newPtr, str)
		}
	}
	*ptr = newPtr
	return len(newPtr)
}
