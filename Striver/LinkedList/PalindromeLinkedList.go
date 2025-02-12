package linkedlist

func PalindromeLinkedList(head *ListNode) bool {
	reverse := func(node *ListNode) *ListNode {
		var prev *ListNode

		// reverse linked list from mid to last
		for node != nil {
			temp := node.Next
			node.Next = prev
			prev = node
			node = temp
		}
		return prev
	}

	dummyNode := &ListNode{}
	dummyNode.Next = head

	fastPtr, slowPtr := dummyNode, dummyNode
	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	midPtr := slowPtr

	reversed_node := reverse(slowPtr.Next)
	midPtr.Next = reversed_node

	firstHalfPtr := head
	midPtr = midPtr.Next
	flag := true
	for midPtr != nil {
		if firstHalfPtr.Val != midPtr.Val {
			flag = false
			break
		}
		firstHalfPtr = firstHalfPtr.Next
		midPtr = midPtr.Next
	}

	// change the reversed node to original node
	prev := reverse(reversed_node)
	slowPtr.Next = prev

	return flag
}

func PalindromeLinkedListRecursion(head *ListNode) bool {
	frontPointer := head
	var checkRecursively func(curNode *ListNode) bool
	checkRecursively = func(curNode *ListNode) bool {
		if curNode != nil {
			if !checkRecursively(curNode.Next) {
				return false
			}

			if curNode.Val != frontPointer.Val {
				return false
			}

			frontPointer = frontPointer.Next

		}
		return true
	}

	return checkRecursively(head)
}
