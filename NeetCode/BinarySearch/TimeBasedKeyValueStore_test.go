package binarysearch_test

import (
	"testing"

	binarysearch "github.com/Binit-Dhakal/LeetCode-Go/BinarySearch"
)

func TestTimeBasedDataStructure(t *testing.T) {
	t.Run("test multiple lower indexes", func(t *testing.T) {
		timeMap := binarysearch.Constructor()
		timeMap.Set("foo", "bar3", 3)
		timeMap.Set("foo", "bar5", 5)
		timeMap.Set("foo", "bar10", 10)
		timeMap.Set("foo", "bar20", 20)
		testCases := []struct {
			name           string
			inputTimeStamp int
			output         string
		}{
			{"for ans 10", 15, "bar10"},
			{"for ans 4", 4, "bar3"},
			{"for timestamp that exist", 10, "bar10"},
			{"for timestamp with no lower value", 1, ""},
		}

		for _, test := range testCases {
			got := timeMap.Get("foo", test.inputTimeStamp)
			if got != test.output {
				t.Errorf("got: %v, want: %v", got, test.output)
			}
		}
	})
}
