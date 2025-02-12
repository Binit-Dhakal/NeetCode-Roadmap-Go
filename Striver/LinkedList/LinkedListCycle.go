package linkedlist

func LLCycle(head *ListNode) bool {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	slowPtr := dummyNode
	fastPtr := dummyNode

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
		if fastPtr == slowPtr {
			return true
		}
	}

	return false
}
