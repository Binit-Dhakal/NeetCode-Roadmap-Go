package trees

func KthSmallestElementBSTDFS(root *TreeNode, k int) int {
	var res int

	var smallest func(root *TreeNode)

	smallest = func(root *TreeNode) {
		if root == nil {
			return
		}

		smallest(root.Left)

		k--
		if k == 0 {
			res = root.Val
			return
		}

		smallest(root.Right)
	}
	smallest(root)

	return res
}

func KthSmallestElementBSTIterative(root *TreeNode, k int) int {
	cur := root
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return cur.Val
		}

		cur = cur.Right
	}
	return 0
}

func KthSmallestElementBSTMorrisTraversal(root *TreeNode, k int) int {
	// Morris InOrder Traversal  Space- O(1) Time - O(n)
	cur := root
	for cur != nil {
		if cur.Left == nil {
			k--
			if k == 0 {
				return cur.Val
			}
			cur = cur.Right
		} else {
			// find the predecessor(right most node in left subtree)
			prev := cur.Left
			for prev.Right != nil && prev.Right != cur {
				prev = prev.Right
			}

			if prev.Right == nil {
				// we establish the temporary link
				prev.Right = cur
				cur = cur.Left
			} else {
				// if already linked, remove the link, move to right child
				prev.Right = nil
				k--
				if k == 0 {
					return cur.Val
				}
				cur = cur.Right
			}

		}
	}
	return 0
}
