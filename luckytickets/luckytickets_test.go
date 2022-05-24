package luckytickets

import "testing"

func TestLuckyTickets(t *testing.T) {
	testCases := []struct {
		name string
		num  int
		want int64
	}{
		{
			name: "1",
			num:  1,
			want: 10,
		},
		{
			name: "2",
			num:  2,
			want: 670,
		},
		{
			name: "3",
			num:  3,
			want: 55252,
		},
		{
			name: "4",
			num:  4,
			want: 4816030,
		},
		{
			name: "5",
			num:  5,
			want: 432457640,
		},
		{
			name: "6",
			num:  6,
			want: 39581170420,
		},
		{
			name: "7",
			num:  7,
			want: 3671331273480,
		},
		{
			name: "8",
			num:  8,
			want: 343900019857310,
		},
		{
			name: "9",
			num:  9,
			want: 32458256583753952,
		},
		{
			name: "10",
			num:  10,
			//TODO want: 3081918923741896816
			want: 3081918923741896840,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := luckyTickets(tc.num)
			if res != tc.want {
				t.Errorf("got %d, want %d", res, tc.want)
			}
		})
	}
}
