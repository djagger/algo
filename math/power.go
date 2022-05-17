package math

// O(N/2+LogN) => O(N)
func power0(num float64, power int) float64 {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	if power == 0 {
		return 1
	}

	num0 := num

	i := 1

	// power of num: num^2, num^4 num^8, num^16, num^32...
	for ; i <= power/2; i *= 2 {
		num *= num
	}

	for ; i < power; i++ {
		num *= num0
	}

	return num
}

// O(2LogN) => O(LogN)
func power1(num float64, power int) float64 {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	if power == 0 {
		return 1
	}

	res := float64(1)
	for ; power > 1; power /= 2 {
		if power%2 == 1 {
			res *= num
		}

		num *= num
	}

	res *= num

	return res
}
