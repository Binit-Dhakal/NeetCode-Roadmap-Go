package graph_test

import (
	"reflect"
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestCourseScheduleII(t *testing.T) {
	t.Run("Using reverse", func(t *testing.T) {
		output := graph.CourseScheduleIIReverseNodes(2, [][]int{{1, 0}})
		if !reflect.DeepEqual(output, []int{0, 1}) {
			t.Errorf("Want: %v, got: %v", []int{0, 1}, output)
		}

		output = graph.CourseScheduleIIReverseNodes(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
		if !(reflect.DeepEqual(output, []int{0, 2, 1, 3}) || reflect.DeepEqual(output, []int{0, 1, 2, 3})) {
			t.Errorf("Got: %v", output)
		}

		output = graph.CourseScheduleIIReverseNodes(2, [][]int{{0, 1}, {1, 0}})
		if len(output) != 0 {
			t.Errorf("Got: %v, want: []int", output)
		}
	})

	t.Run("Using Kahn Topological ordering", func(t *testing.T) {
		output := graph.CourseScheduleIIKahnAlgo(2, [][]int{{1, 0}})
		if !reflect.DeepEqual(output, []int{0, 1}) {
			t.Errorf("Want: %v, got: %v", []int{0, 1}, output)
		}

		output = graph.CourseScheduleIIKahnAlgo(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
		if !(reflect.DeepEqual(output, []int{0, 2, 1, 3}) || reflect.DeepEqual(output, []int{0, 1, 2, 3})) {
			t.Errorf("Got: %v", output)
		}

		output = graph.CourseScheduleIIKahnAlgo(2, [][]int{{0, 1}, {1, 0}})
		if len(output) != 0 {
			t.Errorf("Got: %v, want: []int", output)
		}

		output = graph.CourseScheduleIIKahnAlgo(4, [][]int{{1, 0}, {2, 1}, {3, 2}, {2, 0}, {2, 3}})
		if len(output) != 0 {
			t.Errorf("Got: %v, want: []int", output)
		}
	})
}
