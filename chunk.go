package piscine

import "fmt"

func Chunk(slice []int, size int) {
	if len(slice) != 0 && size != 0 {
		chunk := [][]int(nil)
		for i := 0; ; i += size {
			if i+size >= len(slice) {
				chunk = append(chunk, slice[i:])
				break
			} else {
				chunk = append(chunk, slice[i:i+size])
			}
		}
		fmt.Println(chunk)
	} else if size == 0 {
		fmt.Println()
	} else {
		fmt.Println([]int{})
	}
}
