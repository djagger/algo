package math

import "testing"

func TestPower0(t *testing.T) {
	testCases := []struct {
		name  string
		num   float64
		power int
		want  float64
	}{
		{
			name: "2^10",
			num:  2, power: 10,
			want: 1024,
		},
		{
			name: "zero power",
			num:  123456789, power: 0,
			want: 1,
		},
		{
			name: "1.001^1000",
			num:  1.001, power: 1000,
			want: 2.716923932235562,
		},
		{
			name: "1.0001^10000",
			num:  1.0001, power: 10000,
			want: 2.7181459268244557,
		},
		{
			name: "1.00001^100000",
			num:  1.00001, power: 100000,
			want: 2.718268237195739,
		},
		{
			name: "1.000001^1000000",
			num:  1.000001, power: 1000000,
			want: 2.7182804691276115,
		},
		{
			name: "1.0000001^10000000",
			num:  1.0000001, power: 10000000,
			want: 2.7182816940044554,
		},
		{
			name: "1.00000001^100000000",
			num:  1.00000001, power: 100000000,
			want: 2.718281790326462,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := power0(tc.num, tc.power)
			if res != tc.want {
				t.Errorf("got %f, want %f", res, tc.want)
			}
		})
	}
}

func TestPower1(t *testing.T) {
	testCases := []struct {
		name  string
		num   float64
		power int
		want  float64
	}{
		{
			name: "2^10",
			num:  2, power: 10,
			want: 1024,
		},
		{
			name: "zero power",
			num:  123456789, power: 0,
			want: 1,
		},
		{
			name: "1.001^1000",
			num:  1.001, power: 1000,
			want: 2.7169239322355203,
		},
		{
			name: "1.0001^10000",
			num:  1.0001, power: 10000,
			want: 2.718145926824356,
		},
		{
			name: "1.00001^100000",
			num:  1.00001, power: 100000,
			want: 2.7182682371975284,
		},
		{
			name: "1.000001^1000000",
			num:  1.000001, power: 1000000,
			want: 2.7182804691564275,
		},
		{
			name: "1.0000001^10000000",
			num:  1.0000001, power: 10000000,
			want: 2.7182816939803724,
		},
		{
			name: "1.00000001^100000000",
			num:  1.00000001, power: 100000000,
			want: 2.7182817863957975,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := power1(tc.num, tc.power)
			if res != tc.want {
				t.Errorf("got %f, want %f", res, tc.want)
			}
		})
	}
}
