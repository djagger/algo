package chessbits

import "testing"

func TestChess_Knight(t *testing.T) {
	testCases := []struct {
		pos       int
		wantMoves int
		wantMask  uint64
	}{
		{
			pos:       0,
			wantMoves: 2,
			wantMask:  132096,
		},
		{
			pos:       1,
			wantMoves: 3,
			wantMask:  329728,
		},
		{
			pos:       2,
			wantMoves: 4,
			wantMask:  659712,
		},
		{
			pos:       36,
			wantMoves: 8,
			wantMask:  11333767002587136,
		},
		{
			pos:       47,
			wantMoves: 4,
			wantMask:  4620693356194824192,
		},
		{
			pos:       48,
			wantMoves: 3,
			wantMask:  288234782788157440,
		},
		{
			pos:       54,
			wantMoves: 4,
			wantMask:  1152939783987658752,
		},
		{
			pos:       55,
			wantMoves: 3,
			wantMask:  2305878468463689728,
		},
		{
			pos:       56,
			wantMoves: 2,
			wantMask:  1128098930098176,
		},
		{
			pos:       63,
			wantMoves: 2,
			wantMask:  9077567998918656,
		},
	}

	for _, tc := range testCases {
		t.Run("chess bits - knight", func(t *testing.T) {
			moves, mask := knight(tc.pos)

			if moves != tc.wantMoves {
				t.Errorf("got %d, wantMoves %d", moves, tc.wantMoves)
			}

			if mask != tc.wantMask {
				t.Errorf("got %d, wantMask %d", mask, tc.wantMask)
			}
		})
	}
}
