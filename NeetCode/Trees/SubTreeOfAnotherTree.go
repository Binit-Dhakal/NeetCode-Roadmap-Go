package trees

import "fmt"

func sameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return sameTree(a.Left, b.Left) && sameTree(a.Right, b.Right)
}

func SubTreeOfAnotherTreeDFS(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	isEqual := false
	if root.Val == subRoot.Val {
		// then we check other nodes
		isEqual = sameTree(root, subRoot)
	}

	return isEqual || SubTreeOfAnotherTreeDFS(root.Left, subRoot) || SubTreeOfAnotherTreeDFS(root.Right, subRoot)
}

func serialize(root *TreeNode) string {
	if root == nil {
		return "$#"
	}

	return "$" + fmt.Sprintf("%d", root.Val) + serialize(root.Left) + serialize(root.Right)
}

func zAlgorithm(input string) []int {
	size := len(input)
	zArr := make([]int, size)

	l, r := 0, 0
	for k := 1; k < size; k++ {
		if k > r {
			// outside z box
			l, r = k, k
			for r < size && input[r-k] == input[r] {
				r++
			}

			zArr[k] = r - l
			r--
		} else {
			if zArr[k-l] < r-k+1 {
				// inside the z-box
				zArr[k] = zArr[k-l]
			} else {
				l = k
				for r+1 < size && input[r+1] == input[r-l+1] {
					r++
				}

				zArr[k] = r - l + 1
				r--
			}
		}
	}
	return zArr
}

func SubTreeOfAnotherTreeSerialize(root *TreeNode, subRoot *TreeNode) bool {
	s1 := serialize(root)
	s2 := serialize(subRoot)

	combinedStr := s2 + "|" + s1

	zArr := zAlgorithm(combinedStr)

	subTreeLen := len(s2)

	for _, z := range zArr {
		if z == subTreeLen {
			return true
		}
	}

	return false
}
