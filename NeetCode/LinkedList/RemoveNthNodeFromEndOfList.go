package linkedlist

func RemoveNthNodeFromEndBF(head *ListNode, n int) *ListNode {
	// count total number of item in list
	length := 0
	current := head
	for current != nil {
		length++
		current = current.Next
	}

	toRemoveNode := length - n
	if toRemoveNode == 0 {
		return head.Next
	}

	i := 0
	current = head
	for current != nil {
		if i == toRemoveNode-1 {
			current.Next = current.Next.Next
			break
		}
		i++
		current = current.Next
	}

	return head
}

func RemoveNthNodeFromEndOptimized(head *ListNode, n int) *ListNode {
	var isNthNodeFromLast func(*ListNode) (*ListNode, bool)
	isNthNodeFromLast = func(current *ListNode) (*ListNode, bool) {
		c := current
		for i := 0; i < n; i++ {
			c = c.Next
		}
		// we have to remove head of ll
		if c == nil {
			return nil, true
		}

		// checking if it is last element
		if c.Next == nil {
			return current, true
		}
		return nil, false
	}

	current := head
	if current.Next == nil {
		return current.Next
	}
	for current != nil {
		node, ok := isNthNodeFromLast(current)
		if node == nil && !ok {
			current = current.Next
			continue
		} else if node == nil && ok {
			head = head.Next
			break
		}
		node.Next = node.Next.Next
		break
	}

	return head
}

func RemoveNthNodeFromEndTwoPointers(head *ListNode, n int) *ListNode {
	rightNode := head
	for i := 0; i < n; i++ {
		rightNode = rightNode.Next
	}

	// then we have to remove first element
	if rightNode == nil {
		return head.Next
	}

	leftNode := head

	for rightNode.Next != nil {
		leftNode = leftNode.Next
		rightNode = rightNode.Next
	}

	leftNode.Next = leftNode.Next.Next
	return head
}

func RemoveNthNodeFromEndRecursion(head *ListNode, n int) *ListNode {
	var remove func(head *ListNode, n *int) *ListNode
	remove = func(head *ListNode, n *int) *ListNode {
		if head == nil {
			return nil
		}
		head.Next = remove(head.Next, n)
		(*n)--

		if *n == 0 {
			return head.Next
		}
		return head
	}

	return remove(head, &n)
}
