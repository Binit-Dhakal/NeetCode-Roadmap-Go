package main

import "fmt"

func usingBruteForce(nums []int, target int) []int {
	// Time-O(n^2) and Space- O(1)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func usingHashmap(nums []int, target int) []int {
	// Time- O(n) and Space- O(n)

	hashmap := make(map[int]int)

	var y int
	for idx, num := range nums {
		y = target - num

		idy, ok := hashmap[num]
		if ok {
			return []int{idy, idx}
		} else {
			hashmap[y] = idx
		}
	}
	return nil

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(usingBruteForce(nums, target))
	fmt.Println(usingHashmap(nums, target))

}
