package linkedlist

func LLCycle(head *ListNode) bool {
	slowNode := head
	fastNode := head

	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next

		if slowNode == fastNode {
			return true
		}
	}

	return false
}
