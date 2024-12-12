package heappq_test

import (
	"testing"

	heappq "github.com/Binit-Dhakal/LeetCode-Go/HeapPQ"
)

func TestLastStoneWeight(t *testing.T) {
	testCases := []struct {
		name   string
		stones []int
		output int
	}{
		{"Ex-1", []int{2, 7, 4, 1, 8, 1}, 1},
		{"Ex-2", []int{1}, 1},
		{"Ex-3", []int{2, 3, 6, 2, 4}, 1},
		{"Ex-4", []int{1, 3}, 2},
	}

	t.Run("Heap Approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				temp := make([]int, len(test.stones))
				copy(temp, test.stones)
				output := heappq.LastStoneWeight(temp)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Bucket approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				temp := make([]int, len(test.stones))
				copy(temp, test.stones)
				output := heappq.LastStoneWeightBucketSort(temp)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
