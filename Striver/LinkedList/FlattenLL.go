package linkedlist

func FlattenLL(root *Node) *Node {
	// slice to store n head
	heads := make([]*Node, 0)
	cur := root
	for cur != nil {
		heads = append(heads, cur)
		cur = cur.Next
	}

	dummyHead := &Node{}
	curHead := dummyHead

	for {
		smallest := &Node{Val: -1}
		smallIdx := -1

		for i := 0; i < len(heads); i++ {
			if heads[i] != nil {
				if heads[i].Val < smallest.Val {
					smallest = heads[i]
					smallIdx = i
				}
			}
		}

		// all values in list are nil
		if smallIdx == -1 {
			break
		}

		curHead.Next = heads[smallIdx]
		curHead = heads[smallIdx]
		heads[smallIdx] = heads[smallIdx].Child
	}

	return curHead
}
