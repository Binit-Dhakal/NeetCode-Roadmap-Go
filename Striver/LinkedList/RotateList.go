package linkedlist

func RotateList(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// count length of nodes in head
	length := 1
	cur := head // will keep track of last element

	for cur.Next != nil {
		length++
		cur = cur.Next
	}

	// to account for if k is more than length
	k = length - k%length

	// now we apply the logic for removing nth node from last

	ptr := head
	for range k - 1 {
		ptr = ptr.Next
	}

	cur.Next = head
	head = ptr.Next
	ptr.Next = nil // last element

	return head
}
