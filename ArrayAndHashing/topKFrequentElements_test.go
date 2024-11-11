package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestTopKFrequentElement(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		k      int
		output []int
	}{
		{"Normal Case", []int{1, 2, 2, 3, 3, 3}, 2, []int{2, 3}},
		{"Same number case", []int{7, 7}, 1, []int{7}},
		{"Only one number case", []int{1}, 1, []int{1}},
	}

	t.Run("Test Normal sort", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.TopKSorting, test.input, test.k)
		}
	})

	t.Run("Test Bucket sort", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.TopKBucketSort, test.input, test.k)
		}
	})

	t.Run("Test Heap solution", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, arrayandhashing.TopKHeap, test.input, test.k)
		}
	})
}
