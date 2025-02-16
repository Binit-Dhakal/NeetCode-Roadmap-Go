package linkedlist

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	twoSumHash := func(i int) {
		visited := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {
			complement := -nums[i] - nums[j]
			if _, ok := visited[complement]; ok {
				result = append(result, []int{nums[i], complement, nums[j]})

				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
			}
			visited[nums[j]] = struct{}{}
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			twoSumHash(i)
		}
	}

	return result
}

func ThreeSumWithoutSort(nums []int) [][]int {
	dups := make(map[int]struct{})
	visited := make(map[int]int)
	result := make([][]int, 0)
	seenTriplet := make(map[[3]int]struct{})

	for i, num := range nums {
		if _, ok := dups[num]; ok {
			continue
		}
		dups[num] = struct{}{}

		for _, n := range nums[i+1:] {
			complement := -num - n
			if v, ok := visited[complement]; ok {
				if v == i {
					triplet := [3]int{num, complement, n}
					sort.Ints(triplet[:])

					if _, ok := seenTriplet[triplet]; !ok {
						seenTriplet[triplet] = struct{}{}
						result = append(result, []int{triplet[0], triplet[1], triplet[2]})
					}
				}
			}
			visited[n] = i
		}
	}
	return result
}
