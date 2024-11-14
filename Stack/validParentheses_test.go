package stack_test

import (
	"testing"

	stack "github.com/Binit-Dhakal/LeetCode-Go/Stack"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestValidParentheses(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output bool
	}{
		{"only one bracket", ")", false},
		{"one valid brackets", "()", true},
		{"all valid brackets in correct order", "()[]{}", true},
		{"one invalid bracket", "(]", false},
		{"nested brackets", "([])", true},
	}

	t.Run("Solution using stack", func(t *testing.T) {
		for _, test := range testCases {
			utils.HelperTest(t, test.name, test.output, stack.ValidParenthesesStack, test.input)
		}
	})
}
