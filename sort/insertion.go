package sort

func InsertionSortInt(data []int) {
	for i := 1; i < len(data); i++ {
		j := i - 1

		for i >= 0 && data[j] > data[j+1] {
			data[j], data[j+1] = data[j+1], data[j]
			j--
		}
	}
}
