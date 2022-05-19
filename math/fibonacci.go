package math

import (
	"math"
)

func fibonacciGoldenRatio(n int) int64 {
	sqrt5 := math.Sqrt(5)
	phi := (sqrt5 + 1) / 2

	powPhi := math.Pow(phi, float64(n))
	return int64(powPhi/sqrt5 + 0.5)
}

// O(LogN)
func fibonacciMatrix(n int) int64 {
	a, b := 1, 1
	c, rc, tc := 1, 0, 0
	d, rd := 0, 1

	for n != 0 {
		if n&1 == 1 {
			tc = rc
			rc = rc*a + rd*c
			rd = tc*b + rd*d
		}

		ta := a
		tb := b
		tc = c
		a = a*a + b*c
		b = ta*b + b*d
		c = c*ta + d*c
		d = tc*tb + d*d

		n >>= 1
	}

	return int64(rc)
}
