package math

import (
	"fmt"
	"testing"
	"time"
)

func TestPrime_Simple(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int64
	}{
		{
			name: "0-10",
			num:  10,
			want: 4,
		},
		{
			name: "0-1",
			num:  1,
			want: 0,
		},
		{
			name: "0-2",
			num:  2,
			want: 1,
		},
		{
			name: "0-3",
			num:  3,
			want: 2,
		},
		{
			name: "0-4",
			num:  4,
			want: 2,
		},
		{
			name: "0-5",
			num:  5,
			want: 3,
		},
		{
			name: "0-100",
			num:  100,
			want: 25,
		},
		{
			name: "0-1000",
			num:  1000,
			want: 168,
		},
		{
			name: "0-10000",
			num:  10000,
			want: 1229,
		},
		{
			name: "0-100000",
			num:  100000,
			want: 9592,
		},
		{
			name: "0-1000000",
			num:  1000000,
			want: 78498,
		},
		{
			name: "0-10000000",
			num:  10000000,
			want: 664579,
		},
		// A lot of time.
		//{
		//	name: "0-100000000",
		//	num:  100000000,
		//	want: 5761455,
		//},
		//{
		//	name: "0-1000000000",
		//	num:  1000000000,
		//	want: 50847534,
		//},
		//{
		//	name: "0-123456789",
		//	num:  123456789,
		//	want: 7027260,
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			res := countPrimesSimple(tc.num)
			fmt.Println(time.Since(start))

			if res != tc.want {
				t.Errorf("got %d, want %d", res, tc.want)
			}
		})
	}
}

func TestPrime_SieveOfEratosthenes(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int64
	}{
		{
			name: "0-10",
			num:  10,
			want: 4,
		},
		{
			name: "0-1",
			num:  1,
			want: 0,
		},
		{
			name: "0-2",
			num:  2,
			want: 1,
		},
		{
			name: "0-3",
			num:  3,
			want: 2,
		},
		{
			name: "0-4",
			num:  4,
			want: 2,
		},
		{
			name: "0-5",
			num:  5,
			want: 3,
		},
		{
			name: "0-100",
			num:  100,
			want: 25,
		},
		{
			name: "0-1000",
			num:  1000,
			want: 168,
		},
		{
			name: "0-10000",
			num:  10000,
			want: 1229,
		},
		{
			name: "0-100000",
			num:  100000,
			want: 9592,
		},
		{
			name: "0-1000000",
			num:  1000000,
			want: 78498,
		},
		{
			name: "0-10000000",
			num:  10000000,
			want: 664579,
		},
		{
			name: "0-100000000",
			num:  100000000,
			want: 5761455,
		},
		// ~13 seconds.
		//{
		//	name: "0-1000000000",
		//	num:  1000000000,
		//	want: 50847534,
		//},
		{
			name: "0-123456789",
			num:  123456789,
			want: 7027260,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			res := countPrimesSieveOfEratosthenes(tc.num)
			fmt.Println(time.Since(start))

			if res != tc.want {
				t.Errorf("got %d, want %d", res, tc.want)
			}
		})
	}
}
