package math

func countPrimesSimple(n int) int64 {
	count := int64(0)

	for num := int64(0); num <= int64(n); num++ {
		if isPrimeSimple(num) {
			count++
		}
	}

	return count
}

func isPrimeSimple(num int64) bool {
	if num < 2 {
		return false
	}

	if num < 4 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := int64(3); i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func countPrimesSieveOfEratosthenes(n int) int64 {
	integers := make([]bool, n+1)
	for i := 2; i < n+1; i++ {
		integers[i] = true
	}

	for p := 2; p <= n; p++ {
		for i := p * p; i <= n; i += p {
			integers[i] = false
		}
	}

	count := 0
	for _, i := range integers {
		if i {
			count++
		}
	}

	return int64(count)
}
