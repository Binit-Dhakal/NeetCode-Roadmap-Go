package linkedlist

func MergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var current, head, node *ListNode
	current = &ListNode{}
	head = current

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			node = list1
			list1 = list1.Next
		} else {
			node = list2
			list2 = list2.Next
		}

		current.Next = node
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return head.Next
}

func MergeTwoSortedListsRecursion(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = MergeTwoSortedListsRecursion(list1.Next, list2)
		return list1
	}

	list2.Next = MergeTwoSortedListsRecursion(list1, list2.Next)
	return list2
}
