package trees

import "reflect"

func SameTreeBFS(p *TreeNode, q *TreeNode) bool {
	var BFS func(*TreeNode) []*int
	BFS = func(root *TreeNode) []*int {
		stack := []*TreeNode{root}
		res := []*int{}

		for len(stack) > 0 {
			node := stack[0]

			stack = stack[1:]
			if node == nil {
				res = append(res, nil)
				continue
			}
			res = append(res, &node.Val)

			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}

		return res
	}

	res1 := BFS(p)
	res2 := BFS(q)

	if !reflect.DeepEqual(res1, res2) {
		return false
	}
	return true
}

func SameTreeDFS(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p != nil && q != nil && p.Val == q.Val {
		return SameTreeDFS(p.Left, q.Left) && SameTreeDFS(p.Right, q.Right)
	}
	return false
}
