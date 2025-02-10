package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{}
	cur := head

	for i, num := range nums {
		cur.Val = num
		if i == len(nums)-1 {
			cur.Next = nil
			break
		}
		cur.Next = &ListNode{}
		cur = cur.Next
	}
	return head
}

func TraverseLinkedList(head *ListNode) []int {
	cur := head
	output := make([]int, 0)
	for cur != nil {
		output = append(output, cur.Val)
		cur = cur.Next
	}
	return output
}
