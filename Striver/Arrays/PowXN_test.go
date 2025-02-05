package arrays_test

import (
	"testing"

	arrays "github.com/Binit-Dhakal/LeetCode-Go/Striver/Arrays"
)

func TestPowXN(t *testing.T) {
	testCases := []struct {
		name   string
		x      float64
		n      int
		output float64
	}{
		{"Ex-1", 2.00000, 10, 1024.000000},
		{"Ex-2", 2.10000, 3, 9.26100},
		{"Ex-3", 2.00000, -2, 0.25000},
		{"Ex-4", 0.00000, 10, 0},
		{"Ex-5", 1.0000, 0, 1.0000},
	}

	t.Run("Brute Force", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := arrays.PowXN(test.x, test.n)
				pos := output - test.output
				if pos < 0 {
					pos = -pos
				}
				if pos > 1e-9 {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
