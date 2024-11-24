package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{values[0], nil}

	node := head
	for i := 1; i < len(values); i++ {
		node.Next = &ListNode{Val: values[i], Next: nil}
		node = node.Next
	}
	return head
}

func TraverseLinkedList(head *ListNode) []int {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func ReverseLinkedList(head *ListNode) *ListNode {
	var prev, next *ListNode
	current := head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
