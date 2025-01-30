package backtracking_test

import (
	"testing"

	backtracking "github.com/Binit-Dhakal/LeetCode-Go/Backtracking"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestPalindromePartioning(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output [][]string
	}{
		// {"Ex-1", "aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		// {"Ex-2", "a", [][]string{{"a"}}},
		// {"Ex-3", "aabb", [][]string{{"a", "a", "b", "b"}, {"aa", "b", "b"}, {"aa", "bb"}, {"a", "a", "bb"}}},
		{"Ex-4", "civicabbac", [][]string{
			{"c", "i", "v", "i", "c", "a", "b", "b", "a", "c"},
			{"c", "i", "v", "i", "c", "a", "bb", "a", "c"},
			{"c", "i", "v", "i", "c", "abba", "c"},
			{"c", "i", "v", "i", "cabbac"},
			{"c", "ivi", "c", "a", "b", "b", "a", "c"},
			{"c", "ivi", "c", "a", "bb", "a", "c"},
			{"c", "ivi", "c", "abba", "c"},
			{"c", "ivi", "cabbac"},
			{"civic", "a", "b", "b", "a", "c"},
			{"civic", "a", "bb", "a", "c"},
			{"civic", "abba", "c"},
		}},
	}

	t.Run("backtracking", func(t *testing.T) {
		for _, test := range testCases {
			output := backtracking.PalindromePartitioning(test.input)
			if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
				t.Errorf("Got: %v, Want: %v", output, test.output)
			}
		}
	})

	t.Run("backtracking dp", func(t *testing.T) {
		for _, test := range testCases {
			output := backtracking.PalindromePartitioningDP(test.input)
			if ok, _ := utils.CheckUnorderedSliceEquality(output, test.output); !ok {
				t.Errorf("Got: %v, Want: %v", output, test.output)
			}
		}
	})
}

func BenchmarkPalindromePartioningDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := "civicabbacabcdef"
		b.StartTimer()
		backtracking.PalindromePartitioningDP(input)
	}
}

func BenchmarkPalindromePartioningBackTracking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := "civicabbacabcdef"
		b.StartTimer()
		backtracking.PalindromePartitioning(input)
	}
}
