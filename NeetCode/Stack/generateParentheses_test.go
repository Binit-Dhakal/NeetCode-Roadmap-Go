package stack_test

import (
	"testing"

	stack "github.com/Binit-Dhakal/LeetCode-Go/Stack"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestGenerateParentheses(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output []string
	}{
		{"for input=1", 1, []string{"()"}},
		{"for input=3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	t.Run("Test Generate Parentheses", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.GenerateParenthesesBruteForce, test.input)
		}
	})

	t.Run("Test BackTracking", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.GenerateParenthesesBackTracking, test.input)
		}
	})
}
