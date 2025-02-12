package linkedlist

func DeleteNodeInLL(node *ListNode) {
	cur := node
	var prevNode *ListNode

	for cur.Next != nil {
		prevNode = cur
		cur.Val = cur.Next.Val
		cur = cur.Next
	}

	prevNode.Next = nil
}

func DeleteNodeInLLOptimal(node *ListNode) {
	// we can pretend that next element is the one
	// we have to delete and thus we will have "node" as
	// previous element
	cur := node
	nextNode := node.Next
	cur.Val = nextNode.Val

	cur.Next = nextNode.Next
}
