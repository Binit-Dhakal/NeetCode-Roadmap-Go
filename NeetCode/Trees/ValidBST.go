package trees

import (
	"math"
)

func ValidateBST(root *TreeNode) bool {
	var isValid func(*TreeNode, [2]int) bool

	isValid = func(root *TreeNode, numRange [2]int) bool {
		if root == nil {
			return true
		}

		if root.Val <= numRange[0] || root.Val >= numRange[1] {
			return false
		}

		return isValid(root.Left, [2]int{numRange[0], root.Val}) && isValid(root.Right, [2]int{root.Val, numRange[1]})
	}

	return isValid(root, [2]int{math.MinInt, math.MaxInt})
}
