package linkedlist

func ReverseNodesInKGroupRecursion(head *ListNode, k int) *ListNode {
	cur := head
	group := 0

	for cur != nil && group < k {
		cur = cur.Next
		group++
	}

	if group == k {
		cur := ReverseNodesInKGroupRecursion(cur, k)
		for group > 0 {
			tmp := head.Next
			head.Next = cur
			cur = head
			head = tmp
			group--
		}
		head = cur
	}
	return head
}
