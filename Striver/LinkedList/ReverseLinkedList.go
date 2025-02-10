package linkedlist

func ReverseList(head *ListNode) *ListNode {
	next := head // just so it wont be null (loop condn)
	cur := head
	var prev *ListNode

	for next != nil {
		cur = next
		next = cur.Next
		cur.Next = prev
		prev = cur
	}

	return cur
}
