package trees

func BTLevelOrderTraversalDFS(root *TreeNode) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}

	dfs(root, 0)
	return res
}

func BTLevelOrderTraversalBFS(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	levelTraverse := make([][]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		intStack := make([]int, 0)

		size := len(stack)
		for size > 0 {
			node := stack[0]
			stack = stack[1:]

			intStack = append(intStack, node.Val)

			if node.Left != nil {
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			size--
		}

		levelTraverse = append(levelTraverse, intStack)
	}
	return levelTraverse
}
