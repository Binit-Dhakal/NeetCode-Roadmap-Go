package arrays_test

import (
	"reflect"
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestMajorityElementII(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"Ex-1", []int{1, 2, 2, 3, 2}, []int{2}},
		{"Ex-2", []int{11, 33, 33, 11, 33, 11}, []int{11, 33}},
		{"Ex-3", []int{3, 2, 3}, []int{3}},
	}

	t.Run("Boyer-Moore Algorithm", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.MajorityElementIIVoting(test.nums)
				if !reflect.DeepEqual(output, test.output) {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
