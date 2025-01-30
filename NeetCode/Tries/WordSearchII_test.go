package tries_test

import (
	"testing"

	tries "github.com/Binit-Dhakal/LeetCode-Go/Tries"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestWordSearchII(t *testing.T) {
	testCases := []struct {
		name   string
		board  [][]byte
		words  []string
		output []string
	}{
		{"Ex-1 NC", [][]byte{{'a', 'b', 'c', 'd'}, {'s', 'a', 'a', 't'}, {'a', 'c', 'k', 'e'}, {'a', 'c', 'd', 'n'}}, []string{"bat", "cat", "back", "backend", "stack"}, []string{"cat", "back", "backend"}},
		{"Ex-2", [][]byte{{'x', 'o'}, {'x', 'o'}}, []string{"xoxo"}, []string{}},
		{"Ex-3", [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}}, []string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{"Ex-4", [][]byte{{'a', 'a'}}, []string{"aaa"}, []string{}},
		{"Ex-5", [][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}, []string{"gfedcbaaa"}, []string{"gfedcbaaa"}},
		{"Ex-6", [][]byte{{'a', 'b', 'c'}, {'b', 'a', 'd'}, {'e', 'c', 'a'}}, []string{"abc"}, []string{"abc"}},
	}

	t.Run("Backtracking", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := tries.WordSearchIIBackTracking(test.board, test.words)
				ok, _ := utils.CheckUnorderedSliceEquality(output, test.output)
				if !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})

	t.Run("With Tries - Array", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := tries.WordSearchIITrie(test.board, test.words)
				ok, _ := utils.CheckUnorderedSliceEquality(output, test.output)
				if !ok {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
