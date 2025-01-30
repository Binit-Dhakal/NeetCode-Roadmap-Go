package dp1

func LongestIncreasingSubsequenceTopDown(nums []int) int {
	dp := make([]int, len(nums)+1)
	for i := range dp {
		dp[i] = -1
	}
	dp[len(nums)] = 0

	var dfs func(idx int) int

	dfs = func(idx int) int {
		if dp[idx] != -1 {
			return dp[idx]
		}

		cnt := 1
		for i := idx + 1; i < len(nums); i++ {
			if nums[idx] < nums[i] {
				cnt = max(cnt, 1+dfs(i))
			}
		}
		dp[idx] = cnt

		return cnt
	}

	maxCnt := 0
	for i := 0; i < len(nums); i++ {
		maxCnt = max(maxCnt, dfs(i))
	}
	return maxCnt
}

func LongestIncreasingSubsequenceBottomUp(nums []int) int {
	// O(n^2)
	n := len(nums) - 1
	dp := make([]int, len(nums))

	maxCnt := 0
	for i := n; i >= 0; i-- {
		cnt := 1
		for j := i + 1; j <= n; j++ {
			if nums[i] < nums[j] {
				cnt = max(1+dp[j], cnt)
			}
		}
		dp[i] = cnt
		maxCnt = max(maxCnt, cnt)
	}
	return maxCnt
}

func LongestIncreasingSubsequenceBinarySearch(nums []int) int {
	// O(nlogn)
	sequence := make([]int, 0)
	sequence = append(sequence, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > sequence[len(sequence)-1] {
			sequence = append(sequence, nums[i])
			continue
		}

		// binary search to fit nums[i] into sequence
		l, r := 0, len(sequence)-1

		for l <= r {
			if l == r {
				sequence[l] = nums[i]
				break
			}
			mid := (r + l) / 2
			if sequence[mid] == nums[i] {
				break
			} else if sequence[mid] < nums[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}
	return len(sequence)
}
