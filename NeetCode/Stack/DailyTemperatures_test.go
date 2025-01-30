package stack_test

import (
	"testing"

	stack "github.com/Binit-Dhakal/LeetCode-Go/Stack"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestDailyTemperature(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{"with lots of temperatures", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"with increasing tempt", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"with decreasing temp", []int{30, 20, 10}, []int{0, 0, 0}},
	}

	t.Run("Using stack", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.DailyTemperatures, test.input)
		}
	})
}
