package chessbits

import "testing"

func TestChess_King(t *testing.T) {
	testCases := []struct {
		pos       int
		wantMoves int
		wantMask  uint64
	}{
		{
			pos:       0,
			wantMoves: 3,
			wantMask:  770,
		},
		{
			pos:       1,
			wantMoves: 5,
			wantMask:  1797,
		},
		{
			pos:       7,
			wantMoves: 3,
			wantMask:  49216,
		},
		{
			pos:       8,
			wantMoves: 5,
			wantMask:  197123,
		},
		{
			pos:       10,
			wantMoves: 8,
			wantMask:  920078,
		},
		{
			pos:       15,
			wantMoves: 5,
			wantMask:  12599488,
		},
		{
			pos:       54,
			wantMoves: 8,
			wantMask:  16186183351374184448,
		},
		{
			pos:       55,
			wantMoves: 5,
			wantMask:  13853283560024178688,
		},
		{
			pos:       56,
			wantMoves: 3,
			wantMask:  144959613005987840,
		},
		{
			pos:       63,
			wantMoves: 3,
			wantMask:  4665729213955833856,
		},
	}

	for _, tc := range testCases {
		t.Run("chess bits - king", func(t *testing.T) {
			moves, mask := king(tc.pos)

			if moves != tc.wantMoves {
				t.Errorf("got %d, wantMoves %d", moves, tc.wantMoves)
			}

			if mask != tc.wantMask {
				t.Errorf("got %d, wantMask %d", mask, tc.wantMask)
			}
		})
	}
}
