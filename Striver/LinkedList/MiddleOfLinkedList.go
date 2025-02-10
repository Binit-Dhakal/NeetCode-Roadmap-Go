package linkedlist

func MiddleOfLinkedListFastSlowPtr(head *ListNode) *ListNode {
	slowPtr := head
	fastPtr := head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next
	}

	return slowPtr
}
