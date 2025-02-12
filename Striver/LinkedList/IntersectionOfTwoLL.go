package linkedlist

func IntersectionOfTwoLL(headA, headB *ListNode) *ListNode {
	commonMap := make(map[*ListNode]struct{})
	curA := headA
	curB := headB
	for curA != nil || curB != nil {
		if curA != nil {
			if _, ok := commonMap[curA]; ok {
				return curA
			}
			commonMap[curA] = struct{}{}

			curA = curA.Next
		}

		if curB != nil {
			if _, ok := commonMap[curB]; ok {
				return curB
			}
			commonMap[curB] = struct{}{}

			curB = curB.Next
		}

	}
	return nil
}

func IntersectionOfTwoLLOptimized(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	d1 := headA
	d2 := headB

	for d1 != d2 {
		if d1 != nil {
			d1 = d1.Next
		} else {
			d1 = headB
		}

		if d2 != nil {
			d2 = d2.Next
		} else {
			d2 = headA
		}
	}

	return d1
}
