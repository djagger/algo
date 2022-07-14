package sort

func BubbleSortInt(data []int) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[j], data[i] = data[i], data[j]
			}
		}
	}
}
