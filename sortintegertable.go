package piscine

// Bubble Sort
func SortIntegerTable(table []int) {
	for i := range table {
		for j := range table {
			if table[i] < table[j] {
				Swap(&table[i], &table[j])
			}
		}
	}
}
