package arrays

func PowXN(x float64, n int) float64 {
	var res float64 = 1.0
	invFlag := false
	nn := int64(n)

	// here we can have buffer overflow
	if nn < 0 {
		invFlag = true
		nn = -nn
	}

	for nn > 0 {
		if nn&1 == 1 {
			res *= x
		}

		x *= x   // square x
		nn >>= 1 // divide exponent by 2
	}

	if invFlag {
		return 1 / res
	}
	return res
}
