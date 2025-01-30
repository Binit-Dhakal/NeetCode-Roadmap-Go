package trees

func InvertBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	InvertBinaryTree(root.Left)
	InvertBinaryTree(root.Right)

	return root
}

func InvertBinaryTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) != 0 {
		node := queue[0]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return root
}
