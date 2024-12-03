package trees

func BTMaxPathSum(root *TreeNode) int {
	var maximum int = root.Val

	var maxPath func(*TreeNode) int
	maxPath = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMax := maxPath(root.Left)
		rightMax := maxPath(root.Right)

		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)

		maximum = max(maximum, root.Val+rightMax+leftMax)

		return root.Val + max(leftMax, rightMax)
	}

	maxPath(root)

	return maximum
}
