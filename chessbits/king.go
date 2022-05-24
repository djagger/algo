package chessbits

import "math/bits"

func king(pos int) (int, uint64) {
	k := uint64(1) << pos

	kAndNoA := k & noA
	kAndNoH := k & noH

	var mask uint64
	mask = (kAndNoA << 7) | (k << 8) | (kAndNoH << 9) |
		(kAndNoA >> 1) | (kAndNoH << 1) |
		(kAndNoA >> 9) | (k >> 8) | (kAndNoH >> 7)

	moves := bits.OnesCount64(mask)
	return moves, mask
}
