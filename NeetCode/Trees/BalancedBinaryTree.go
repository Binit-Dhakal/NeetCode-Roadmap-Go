package trees

// import "container/list"

func BalancedBinaryTreeDFS(root *TreeNode) bool {
	if root == nil {
		return true
	}

	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}
	var findDepth func(root *TreeNode) int
	findDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := findDepth(root.Left)
		right := findDepth(root.Right)

		if left == -1 || right == -1 || abs(left-right) > 1 {
			return -1
		}

		return 1 + max(left, right)
	}

	s := findDepth(root)
	if s == -1 {
		return false
	}
	return true
}
