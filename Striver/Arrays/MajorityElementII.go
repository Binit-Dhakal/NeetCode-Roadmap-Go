package arrays

func MajorityElementIIVoting(nums []int) []int {
	// since for n/3 majority there can only be two element
	a, b := 0, 0
	cnt1, cnt2 := 0, 0

	for _, num := range nums {
		if cnt1 == 0 && num != b {
			cnt1++
			a = num
		} else if cnt2 == 0 && num != a {
			cnt2++
			b = num
		} else if num == a {
			cnt1++
		} else if num == b {
			cnt2++
		} else {
			cnt1--
			cnt2--
		}
	}

	cnt1, cnt2 = 0, 0
	for _, num := range nums {
		if num == a {
			cnt1++
		} else if num == b {
			cnt2++
		}
	}

	threshold := len(nums) / 3
	res := make([]int, 0)
	if cnt1 > threshold {
		res = append(res, a)
	}
	if cnt2 > threshold {
		res = append(res, b)
	}

	return res
}
