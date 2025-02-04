package arrays

import (
	"slices"
)

func MergeIntervals(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] == b[0] {
			if a[1] < b[1] {
				return -1
			} else if a[1] == b[1] {
				return 0
			}
		}
		return 1
	})
	res := make([][]int, 1)
	res[0] = intervals[0]
	cur := 0

	for i := 1; i < len(intervals); i++ {
		if res[cur][1] >= intervals[i][0] {
			res[cur][1] = max(intervals[i][1], res[cur][1])
		} else {
			// no overlapping
			res = append(res, intervals[i])
			cur++
		}
	}
	return res
}
