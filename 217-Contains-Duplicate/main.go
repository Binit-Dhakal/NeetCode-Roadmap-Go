package main

import (
	"fmt"
	"sort"
)

func using_multiset(nums []int) bool {
	// Multiset is a generalization of set that allows multiple instance of same element.
	// Time: O(2n)=O(n) and Space: O(n)
	set := NewMultiSet()
	set.AddNums(nums)

	for _, num := range nums {
		if set.HasDuplicate(num) {
			return true
		}
	}
	return false
}

func using_hashmap(nums []int) bool {
	// Time: O(n) and Space: O(n)
	hashmap := make(map[int]bool)

	for _, num := range nums {
		if hashmap[num] {
			return true
		}
	}
	return false

}

func sorting(nums []int) bool {
	// We sort all the numbers and check if pairs are beside each other
	// Time - O(nlogn) and Space - O(logn)
	if len(nums) <= 1 {
		return false
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 1}
	res := sorting(nums)
	fmt.Println(res)

}
