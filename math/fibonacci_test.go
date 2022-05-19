package math

import (
	"testing"
)

type fibonacciTestCase struct {
	name string
	num  int
	want int64
}

var fibonacciTestCases = []fibonacciTestCase{
	{
		name: "0",
		num:  0,
		want: 0,
	},
	{
		name: "1",
		num:  1,
		want: 1,
	},
	{
		name: "2",
		num:  2,
		want: 1,
	},
	{
		name: "3",
		num:  3,
		want: 2,
	},
	{
		name: "4",
		num:  4,
		want: 3,
	},
	{
		name: "5",
		num:  5,
		want: 5,
	},
	{
		name: "10",
		num:  10,
		want: 55,
	},

	// TODO big.Int
	//{
	//	name: "100",
	//	num:  100,
	//	want: 354224848179261915075,
	//},
	//{
	//	name: "1000",
	//	num:  1000,
	//	want: 43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875,
	//},
}

func TestFibonacci_GoldenRatio(t *testing.T) {
	for _, tc := range fibonacciTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res := fibonacciGoldenRatio(tc.num)
			if res != tc.want {
				t.Errorf("got %d, want %d", res, tc.want)
			}
		})
	}
}

func TestFibonacci_Matrix(t *testing.T) {
	for _, tc := range fibonacciTestCases {
		t.Run(tc.name, func(t *testing.T) {
			res := fibonacciMatrix(tc.num)
			if res != tc.want {
				t.Errorf("got %d, want %d", res, tc.want)
			}
		})
	}
}
