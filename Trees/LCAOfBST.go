package trees

func LCAOfBSTRecursion(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return LCAOfBSTRecursion(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return LCAOfBSTRecursion(root.Right, p, q)
	} else {
		return root
	}
}

func LCAOfBSTIteration(root, p, q *TreeNode) *TreeNode {
	cur := root

	for cur != nil {
		if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}
	return nil
}
