package arrays_test

import (
	"fmt"
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestMergeSortedArray(t *testing.T) {
	testCases := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		output []int
	}{
		{"Ex-1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"Ex-2", []int{1}, 1, []int{}, 0, []int{1}},
		{"Ex-3", []int{0}, 0, []int{1}, 1, []int{1}},
		{"Ex-4", []int{10, 13, 14, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 10, 13, 14}},
	}

	t.Run("Binary Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				nums1 := make([]int, len(test.nums1))
				copy(nums1, test.nums1)
				arrays.MergeSortedArray(nums1, test.m, test.nums2, test.n)
				if !reflect.DeepEqual(nums1, test.output) {
					t.Errorf("Got: %v, want: %v", nums1, test.output)
				}
			})
		}
	})

	t.Run("Optimal Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				nums1 := make([]int, len(test.nums1))
				copy(nums1, test.nums1)
				fmt.Println(nums1)
				arrays.MergeSortedArrayOptimal(nums1, test.m, test.nums2, test.n)
				fmt.Println(nums1)
				if !reflect.DeepEqual(nums1, test.output) {
					t.Errorf("Got: %v, want: %v", nums1, test.output)
				}
			})
		}
	})
}
