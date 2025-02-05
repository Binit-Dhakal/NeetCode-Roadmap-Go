// https://www.naukri.com/code360/problems/missing-and-repeating-numbers_6828164
package arrays

func RepeatAndMissingNumber(nums []int) []int {
	// returns []int = {duplicateNum, missingNum}
	// X+Y = (AS2-S2)/(AS1-S1)
	// X-Y = (AS1-S1)

	n := len(nums)
	Sum1 := (n * (n + 1)) / 2
	Sum2 := (n * (n + 1) * (2*n + 1)) / 6

	var ASum1, ASum2 int

	for _, num := range nums {
		ASum1 += num
		ASum2 += num * num
	}

	sum1diff := Sum1 - ASum1
	sum2diff := Sum2 - ASum2
	res := make([]int, 2)
	res[0] = (sum2diff/sum1diff - sum1diff) / 2
	res[1] = (sum2diff/sum1diff + sum1diff) / 2
	return res
}

func RepeatAndMissingNumberXOR(nums []int) []int {
	xorAll, xorExpected := 0, 0
	for i, num := range nums {
		xorAll ^= num
		xorExpected ^= (i + 1)
	}

	// res will be P^Q
	res := xorAll ^ xorExpected

	// identify rightmost set bit
	rightMostBit := res & -res

	group1 := 0
	group2 := 0
	for i, num := range nums {
		if num&rightMostBit == 0 {
			group1 ^= num
		} else {
			group2 ^= num
		}

		if (i+1)&rightMostBit == 0 {
			group1 ^= (i + 1)
		} else {
			group2 ^= (i + 1)
		}
	}

	// we have P and Q in group1 and group2 not in order
	qExists := false
	for _, num := range nums {
		if num == group2 {
			qExists = true
			break
		}
	}

	if qExists {
		return []int{group2, group1}
	} else {
		return []int{group1, group2}
	}
}
