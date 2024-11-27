package trees

func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	var findDiameter func(root *TreeNode) int
	findDiameter = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := findDiameter(root.Left)
		right := findDiameter(root.Right)

		maxDiameter = max(maxDiameter, left+right)

		return 1 + max(left, right)
	}
	findDiameter(root)
	return maxDiameter
}
