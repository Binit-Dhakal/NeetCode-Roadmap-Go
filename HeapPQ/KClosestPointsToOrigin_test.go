package heappq_test

import (
	"testing"

	heappq "github.com/Binit-Dhakal/LeetCode-Go/HeapPQ"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestKClosestPointsToOrigin(t *testing.T) {
	testCases := []struct {
		name   string
		points [][]int
		k      int
		output [][]int
	}{
		{"Ex-1", [][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{"Ex-2", [][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
		{"Ex-3", [][]int{{0, 2}, {2, 2}}, 1, [][]int{{0, 2}}},
		{"Ex-4", [][]int{{0, 2}, {2, 0}, {2, 2}}, 2, [][]int{{0, 2}, {2, 0}}},
	}

	t.Run("Using heap", func(t *testing.T) {
		for _, test := range testCases {
			output := heappq.KClosestPointsToOrigin(test.points, test.k)
			if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
				t.Errorf("Got: %v, want: %v", output, test.output)
			}
		}
	})
	t.Run("Using quick select", func(t *testing.T) {
		for _, test := range testCases {
			output := heappq.KClosestPointsToOriginQuickSelect(test.points, test.k)
			if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
				t.Errorf("Got: %v, want: %v", output, test.output)
			}
		}
	})
}
