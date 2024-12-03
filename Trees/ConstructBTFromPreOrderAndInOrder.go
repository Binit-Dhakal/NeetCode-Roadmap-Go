package trees

func ConstructTreeFromPreOrderAndInOrder(preorder []int, inorder []int) *TreeNode {
	preIdx := 0
	inOrderHash := make(map[int]int) // value: index
	for i, val := range inorder {
		inOrderHash[val] = i
	}

	var constructTree func(left, right int) *TreeNode
	constructTree = func(left, right int) *TreeNode {
		if right < left {
			return nil
		}

		inIdx := inOrderHash[preorder[preIdx]]

		rootNode := &TreeNode{Val: preorder[preIdx]}
		preIdx++
		rootNode.Left = constructTree(left, inIdx-1)
		rootNode.Right = constructTree(inIdx+1, right)
		return rootNode
	}

	return constructTree(0, len(inorder)-1)
}
