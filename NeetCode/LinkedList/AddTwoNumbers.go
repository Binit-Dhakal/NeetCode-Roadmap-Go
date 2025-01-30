package linkedlist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummyNode := &ListNode{Val: 0}

	current := dummyNode
	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry != 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummyNode.Next
}

func AddTwoNumbersRecursion(l1 *ListNode, l2 *ListNode) *ListNode {
	var add func(*ListNode, *ListNode, int) *ListNode
	add = func(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil && carry == 0 {
			return nil
		}

		s := 0
		if l1 != nil {
			s += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			s += l2.Val
			l2 = l2.Next
		}
		s += carry

		carry = s / 10

		nextNode := add(l1, l2, carry)
		return &ListNode{Val: s % 10, Next: nextNode}
	}

	return add(l1, l2, 0)
}
