package linkedlist

func ReorderList(head *ListNode) {
	// O(n) - O(n)
	dlist := []*ListNode{}

	current := head
	for current != nil {
		dlist = append(dlist, current)
		current = current.Next
	}

	leftPtr, rightPtr := 1, len(dlist)-1
	leftTurn := false

	current = head
	for leftPtr <= rightPtr {
		if leftTurn {
			current.Next = dlist[leftPtr]
			leftPtr++
			leftTurn = false
		} else {
			current.Next = dlist[rightPtr]
			rightPtr--
			leftTurn = true
		}
		current = current.Next
	}
	current.Next = nil
}

func ReorderListWithReverse(head *ListNode) {
	dummyHead := &ListNode{Val: 0, Next: head}

	fast, slow := dummyHead, dummyHead
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow.Next
	slow.Next = nil

	// Reverse Linked List
	var prev *ListNode
	for fast != nil {
		tmp := fast.Next
		fast.Next = prev
		prev = fast
		fast = tmp
	}

	first := head
	second := prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}

func ReorderListWithRecursion(head *ListNode) {
	if head == nil {
		return
	}

	var reorderList func(root, current *ListNode) *ListNode

	reorderList = func(root, current *ListNode) *ListNode {
		if current == nil {
			return root
		}
		root = reorderList(root, current.Next)

		if root == nil {
			return nil
		}

		var temp *ListNode
		if root == current || root.Next == current {
			current.Next = nil
		} else {
			temp = root.Next
			root.Next = current
			current.Next = temp
		}

		return temp
	}

	head = reorderList(head, head.Next)
}
