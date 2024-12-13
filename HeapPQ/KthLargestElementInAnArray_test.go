package heappq_test

import (
	"testing"

	heappq "github.com/Binit-Dhakal/LeetCode-Go/HeapPQ"
)

func TestKthLargestElementInAnArray(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		k      int
		output int
	}{
		{"Ex-1", []int{3, 2, 1, 5, 6, 4}, 2, 5},
		{"Ex-2", []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
		{"Ex-3", []int{2, 3, 1, 5, 4}, 2, 4},
		{"Ex-4", []int{2, 3, 1, 1, 5, 5, 4}, 3, 4},
	}

	t.Run("Heap approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := make([]int, len(test.nums))
				copy(input, test.nums)
				output := heappq.KthLargestElementInAnArray(input, test.k)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Quick select approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := make([]int, len(test.nums))
				copy(input, test.nums)
				output := heappq.KthLargestElementInAnArrayQuickSelect(input, test.k)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
	t.Run("Bucket Sort approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				input := make([]int, len(test.nums))
				copy(input, test.nums)
				output := heappq.KthLargestElementInAnArrayBucketSort(input, test.k)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
