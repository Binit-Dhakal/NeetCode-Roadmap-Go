package linkedlist

func LLCycleIIHashMap(head *ListNode) *ListNode {
	existsMap := make(map[*ListNode]struct{})
	cur := head

	for cur != nil {
		if _, ok := existsMap[cur]; ok {
			return cur
		}
		existsMap[cur] = struct{}{}
		cur = cur.Next
	}
	return nil
}

func LLCycleFloyds(head *ListNode) *ListNode {
	fastPtr := head
	slowPtr := head

	for fastPtr != nil && fastPtr.Next != nil {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next.Next

		if fastPtr == slowPtr {
			break
		}
	}

	// no cycle
	if fastPtr != slowPtr {
		return nil
	}

	// detect entrance of cycle
	slowPtr = head
	for slowPtr != fastPtr {
		slowPtr = slowPtr.Next
		fastPtr = fastPtr.Next
	}
	return slowPtr
}
