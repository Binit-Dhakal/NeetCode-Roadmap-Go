package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(nums []int, parentIndex int) *TreeNode {
	if parentIndex >= len(nums) {
		return nil
	}

	root := &TreeNode{
		Val: nums[parentIndex],
	}

	root.Left = preOrderTraversal(nums, 2*parentIndex+1)
	root.Right = preOrderTraversal(nums, 2*parentIndex+2)

	return root
}

func BuildBinaryTree(nums []int) *TreeNode {
	return preOrderTraversal(nums, 0)
}

func bfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	bfs(node.Left, res)
	bfs(node.Right, res)
}

func TraverseBinaryTree(node *TreeNode) []int {
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, node)

	for len(queue) > 0 {
		node := queue[0]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}

	return res
}
