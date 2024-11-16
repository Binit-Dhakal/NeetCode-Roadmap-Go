package stack_test

import (
	"testing"

	stack "github.com/Binit-Dhakal/LeetCode-Go/Stack"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestCarFleet(t *testing.T) {
	testCases := []struct {
		name     string
		target   int
		position []int
		speed    []int
		output   int
	}{
		{"different car with different speed", 12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},
		{"only one car", 10, []int{3}, []int{3}, 1},
		{"all car as one fleet", 100, []int{0, 2, 4}, []int{4, 2, 1}, 1},
	}

	t.Run("Car fleet using stack", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.CarFleet, test.target, test.position, test.speed)
		}
	})
}
