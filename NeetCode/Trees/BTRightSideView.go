package trees

func BTRightSideViewDFS(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(root *TreeNode, depth int)

	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}

		if depth == len(res) {
			res = append(res, root.Val)
		}

		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}

	dfs(root, 0)

	return res
}

func BTRightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		lastNodeVal := 0
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			lastNodeVal = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
		res = append(res, lastNodeVal)
	}

	return res
}
