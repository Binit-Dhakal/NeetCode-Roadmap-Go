package linkedlist

func ReverseNodesInKGroupRecursion(head *ListNode, k int) *ListNode {
	var rec func(root *ListNode) *ListNode
	rec = func(root *ListNode) *ListNode {
		cur := root
		for range k {
			if cur == nil {
				return root
			}
			cur = cur.Next
		}

		// reversing from root to kth node
		var prev *ListNode
		cur = root
		for range k {
			temp := cur.Next
			cur.Next = prev
			prev = cur
			cur = temp
		}

		root.Next = rec(cur)

		return prev
	}

	return rec(head)
}

func ReverseNodesInKGroup(head *ListNode, k int) *ListNode {
	// returns reverse head of root after reversing
	reverse := func(root *ListNode) *ListNode {
		h := root
		var prev *ListNode

		for range k {
			temp := h.Next
			h.Next = prev
			prev = h
			h = temp
		}
		return prev
	}

	newHead := &ListNode{}
	curPtr := newHead

	cur := head
	for cur != nil {
		h := cur
		for range k {
			// we dont have enough remaining node
			if cur == nil {
				curPtr.Next = h
				return newHead.Next
			}
			cur = cur.Next
		}

		reverseHead := reverse(h)
		curPtr.Next = reverseHead
		curPtr = h

	}
	return newHead.Next
}
