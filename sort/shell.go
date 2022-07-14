package sort

func ShellSortInt(data []int) {
	for gap := len(data) / 2; gap > 0; gap /= 2 {
		for i := 0; i+gap < len(data); i++ {
			j := i + gap
			tmp := data[j]

			for j-gap >= 0 && data[j-gap] > tmp {
				data[j], data[j-gap] = data[j-gap], data[j]
				j -= gap
			}

			data[j] = tmp
		}
	}
}
