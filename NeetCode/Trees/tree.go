package trees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NilNode = int(^uint(0) >> 1)

func preOrderTraversal(nums []*int, parentIndex int) *TreeNode {
	if parentIndex >= len(nums) {
		return nil
	}

	if nums[parentIndex] == nil {
		return nil
	}

	root := &TreeNode{
		Val: *nums[parentIndex],
	}

	root.Left = preOrderTraversal(nums, 2*parentIndex+1)
	root.Right = preOrderTraversal(nums, 2*parentIndex+2)

	return root
}

func IntPtr(a int) *int {
	return &a
}

func BuildBinaryTree(nums []int) *TreeNode {
	numsPtr := make([]*int, len(nums))

	for idx, num := range nums {
		if num == NilNode {
			numsPtr[idx] = nil
		} else {
			numsPtr[idx] = IntPtr(num)
		}
	}
	root := preOrderTraversal(numsPtr, 0)
	return root
}

func isAllNil(s []*TreeNode) bool {
	for _, val := range s {
		if val != nil {
			return false
		}
	}
	return true
}

func TraverseBinaryTree(node *TreeNode) []int {
	res := make([]int, 0)
	queue := make([]*TreeNode, 0)

	queue = append(queue, node)

	for len(queue) > 0 {
		node := queue[0]
		if isAllNil(queue) {
			break
		}

		queue = queue[1:]
		if node == nil {
			res = append(res, NilNode)
			continue
		}
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		} else {
			queue = append(queue, nil)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		} else {
			queue = append(queue, nil)
		}
	}

	return res
}
