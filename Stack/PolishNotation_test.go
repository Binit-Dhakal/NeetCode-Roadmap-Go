package stack_test

import (
	"testing"

	stack "github.com/Binit-Dhakal/LeetCode-Go/Stack"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestPolishNotation(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		output int
	}{
		{"simple with add and mul", []string{"2", "1", "+", "3", "*"}, 9},
		{"with add and div", []string{"4", "13", "5", "/", "+"}, 6},
		{"long calc", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	t.Run("test polish notation", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.EvalRPNStack, test.input)
		}
	})
}
