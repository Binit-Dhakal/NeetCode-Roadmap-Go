package trees

func CountGoodNodesInBT(root *TreeNode) int {
	count := 1
	parentNodes := make(map[*TreeNode]*TreeNode) // node: valid parentNode

	var traverseNodes func(*TreeNode)
	traverseNodes = func(root *TreeNode) {
		if root == nil {
			return
		}

		nodes := make([]*TreeNode, 0)
		if root.Left != nil {
			nodes = append(nodes, root.Left)
		}
		if root.Right != nil {
			nodes = append(nodes, root.Right)
		}

		for i := 0; i < len(nodes); i++ {
			node := nodes[i]

			if node.Val >= root.Val {
				// this is valid node
				if parentNode, ok := parentNodes[root]; ok && parentNode.Val > node.Val {
					parentNodes[node] = parentNode
				} else {
					count++
				}
			} else {
				if parentNode, ok := parentNodes[root]; ok {
					parentNodes[node] = parentNode
				} else {
					parentNodes[node] = root
				}
			}
		}

		traverseNodes(root.Left)
		traverseNodes(root.Right)
	}

	traverseNodes(root)
	return count
}

func CountGoodNodesInBTOptimal(root *TreeNode) int {
	count := 0
	var check func(root *TreeNode, highest int)

	check = func(root *TreeNode, highest int) {
		if root == nil {
			return
		}

		if root.Val >= highest {
			count++
			highest = root.Val
		}

		check(root.Left, highest)
		check(root.Right, highest)
	}

	check(root, root.Val)
	return count
}
