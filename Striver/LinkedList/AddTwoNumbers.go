package linkedlist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	leftPtr := l1
	rightPtr := l2

	dummyNode := &ListNode{}
	cur := dummyNode
	carry := 0

	for leftPtr != nil || rightPtr != nil || carry != 0 {
		sum := carry
		if leftPtr != nil {
			sum += leftPtr.Val
			leftPtr = leftPtr.Next
		}

		if rightPtr != nil {
			sum += rightPtr.Val
			rightPtr = rightPtr.Next
		}

		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return dummyNode.Next
}
