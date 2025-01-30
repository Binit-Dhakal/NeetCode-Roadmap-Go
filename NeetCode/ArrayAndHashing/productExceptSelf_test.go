package arrayandhashing_test

import (
	"testing"

	arrayandhashing "github.com/Binit-Dhakal/LeetCode-Go/ArrayAndHashing"
	utils "github.com/Binit-Dhakal/LeetCode-Go/Utils"
)

func TestProductExceptSelf(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{"Normal Case", []int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{"Zero case", []int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, test := range testCases {
		utils.HelperTest(t, test.name, test.output, arrayandhashing.ProductExceptSelf, test.input)
	}
}
