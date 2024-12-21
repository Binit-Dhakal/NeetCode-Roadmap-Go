package graph_test

import (
	"testing"

	graph "github.com/Binit-Dhakal/LeetCode-Go/Graph"
)

func TestWordLadder(t *testing.T) {
	testCases := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		output    int
	}{
		{"Ex-1", "hit", "cog", []string{"hot", "dog", "lot", "dot", "log", "cog"}, 5},
		{"Ex-2", "hit", "cog", []string{"hot", "dog", "dot", "lot", "log"}, 0},
	}

	t.Run("BFS", func(t *testing.T) {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {
				output := graph.WordLadder(test.beginWord, test.endWord, test.wordList)
				if output != test.output {
					t.Errorf("Got: %v, want: %v", output, test.output)
				}
			})
		}
	})
}
