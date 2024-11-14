package stack_test

import (
	"testing"

	stack "github.com/Binit-Dhakal/LeetCode-Go/Stack"
)

func TestMinStack(t *testing.T) {
	obj := stack.Constructor()
	obj.Push(10)
	obj.Push(-1)
	obj.Push(20)
	obj.Push(1)
	obj.Push(0)
	obj.Pop()

	if obj.Top() != 1 {
		t.Errorf("got: %v, want: %v", obj.Top(), 1)
	}

	if obj.GetMin() != -1 {
		t.Errorf("got: %v, want: %v", obj.Top(), -1)
	}
}
