package heappq_test

import (
	"testing"

	heappq "github.com/Binit-Dhakal/LeetCode-Go/HeapPQ"
)

func TestTaskScheduler(t *testing.T) {
	testCases := []struct {
		name   string
		tasks  []byte
		n      int
		output int
	}{
		{"Ex-1", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{"Ex-2", []byte{'A', 'C', 'A', 'B', 'D', 'B'}, 1, 6},
		{"Ex-3", []byte{'X', 'X', 'Y', 'Y'}, 2, 5},
		{"Ex-4", []byte{'B', 'C', 'D', 'A', 'A', 'A', 'A', 'G'}, 1, 8},
		{"Ex-5", []byte{'A', 'A', 'A', 'B', 'B', 'B', 'C', 'C', 'C', 'D', 'D', 'E'}, 2, 12},
		{"Ex-6", []byte{'A', 'A', 'A', 'B', 'B', 'B'}, 3, 10},
	}

	t.Run("Brute force approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := heappq.TaskScheduler(test.tasks, test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Heap approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := heappq.TaskSchedulerHeap(test.tasks, test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Math approach", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := heappq.TaskSchedulerMath(test.tasks, test.n)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
