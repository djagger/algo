package chessbits

import "math/bits"

func knight(pos int) (int, uint64) {
	k := uint64(1) << pos

	kAndNoH := k & noH
	kAndNoGH := k & noGH
	kAndNoA := k & noA
	kAndNoAB := k & noAB

	var mask uint64
	mask = (kAndNoH >> 15) | (kAndNoA >> 17) |
		(kAndNoGH >> 6) | (kAndNoAB >> 10) |
		(kAndNoGH << 10) | (kAndNoAB << 6) |
		(kAndNoH << 17) | (kAndNoA << 15)

	moves := bits.OnesCount64(mask)
	return moves, mask
}
