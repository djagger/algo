package luckytickets

import "math"

func luckyTickets(num int) int64 {
	result := int64(0)

	sl := make([]int64, 0)
	for i := 0; i < 10; i++ {
		sl = append(sl, 1)
	}

	for i := 0; i < num-1; i++ {
		sl = nextSlice(sl)
	}

	for _, v := range sl {
		powRes := math.Pow(float64(v), 2)
		result += int64(powRes)
	}

	return result
}

func nextSlice(prevSlice []int64) []int64 {
	sl := make([]int64, 0)

	newLen := len(prevSlice) + 9
	for i := 0; i < newLen; i++ {
		q := int64(0)

		for j := 0; j < 10; j++ {
			idx := i - j

			if idx < 0 {
				continue
			}

			if len(prevSlice) <= idx {
				continue
			}

			q += prevSlice[idx]
		}

		sl = append(sl, q)
	}

	return sl
}
