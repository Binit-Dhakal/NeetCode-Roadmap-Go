package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestCourseSchedule(t *testing.T) {
	testCases := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		output        bool
	}{
		{"Ex-1", 2, [][]int{{1, 0}}, true},
		{"Ex-2", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"Ex-3", 4, [][]int{{0, 2}, {0, 1}, {1, 2}, {2, 3}, {2, 0}}, false},
	}

	t.Run("DFS cycle detection", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.CourseSchedule(test.numCourses, test.prerequisites)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("Kahns Cycle detection", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.CourseScheduleKahnAlgo(test.numCourses, test.prerequisites)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
