package trees

import (
	"container/list"
)

func MaxDepthOfBinaryTreeRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depthLeft, depthRight := 1, 1

	depthLeft += MaxDepthOfBinaryTreeRecursion(root.Left)
	depthRight += MaxDepthOfBinaryTreeRecursion(root.Right)

	return max(depthLeft, depthRight)
}

func MaxDepthOfBinaryTreeBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	maxDepth := 0

	queue = append(queue, root)
	for len(queue) > 0 {
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		maxDepth++
	}
	return maxDepth
}

func MaxDepthOfBinaryTreeStack(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := list.New()
	stack.PushBack([]interface{}{root, 1})
	res := 0
	for stack.Len() != 0 {
		ele := stack.Back()
		stack.Remove(ele)

		pair := ele.Value.([]interface{})
		node := pair[0].(*TreeNode)
		depth := pair[1].(int)

		if node != nil {
			res = max(depth, res)
			stack.PushBack([]interface{}{node.Left, depth + 1})
			stack.PushBack([]interface{}{node.Right, depth + 1})
		}
	}
	return res
}
