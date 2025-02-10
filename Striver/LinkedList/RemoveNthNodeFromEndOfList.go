package linkedlist

func RemoveNthNodeFromEndOfList(head *ListNode, n int) *ListNode {
	length := 0
	fastPtr := head
	slowPtr := head

	count := 0
	for fastPtr != nil && fastPtr.Next != nil {
		count++
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	if fastPtr == nil {
		length = count * 2
	} else if fastPtr.Next == nil {
		length = count*2 + 1
	}

	prevIdx := length - n - 1

	if prevIdx >= count {
		// we can use slowPtr
		for count != prevIdx {
			slowPtr = slowPtr.Next
			count++
		}

		// it is not last element we are deleting
		slowPtr.Next = slowPtr.Next.Next
	} else if prevIdx < 0 {
		// if prevIdx is -ve, then we have to delete first element
		head = head.Next
	} else {
		// prevIdx >= 0 && prevIdx < count
		slowPtr = head
		count = 0
		for count != prevIdx {
			slowPtr = slowPtr.Next
			count++
		}

		slowPtr.Next = slowPtr.Next.Next
	}

	return head
}

func RemoveNthNodeFromEndOfListOptimized(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	fastPtr := dummyNode
	slowPtr := dummyNode

	for range n {
		// as number of nodes is less than n
		fastPtr = fastPtr.Next
	}

	for fastPtr.Next != nil { // until we are end of list
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}

	slowPtr.Next = slowPtr.Next.Next

	return dummyNode.Next
}
