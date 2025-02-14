package linkedlist

func CopyListWithRandomPointer(head *Node) *Node {
	// O(n) space with O(2n) time complexity
	// Not acceptable in interview
	mapNodes := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		mapNodes[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	dummyHead := &Node{}
	curHead := dummyHead
	for cur != nil {
		temp := mapNodes[cur]
		temp.Next = mapNodes[cur.Next]
		temp.Random = mapNodes[cur.Random]
		cur = cur.Next
		curHead.Next = temp
		curHead = curHead.Next
	}

	return dummyHead.Next
}

func CopyListWithRandomPointerInterWeaving(head *Node) *Node {
	cur := head
	// interweaved two nodes
	for cur != nil {
		temp := cur.Next
		node := &Node{Val: cur.Val, Next: temp}
		cur.Next = node
		cur = temp
	}

	// taking random into consideration
	cur = head
	for cur != nil {
		random := cur.Random
		if random != nil {
			cur.Next.Random = random.Next
		} else {
			cur.Next.Random = nil
		}
	}

	// separate original and deepcopied linked list
	cur = head
	newHead := &Node{}
	newCur := newHead

	for cur != nil {
		newCur.Next = cur.Next
		newCur = newCur.Next

		cur.Next = newCur.Next
		cur = cur.Next
	}

	return newHead.Next
}

func CopyListWithRandomPointerIterative(head *Node) *Node {
	mapNodes := make(map[*Node]*Node) // oldNode:newNode
	mapNodes[nil] = nil

	cur := head
	newHead := &Node{}
	newCur := newHead

	for cur != nil {
		newNode, ok := mapNodes[cur]
		if !ok {
			newNode = &Node{Val: cur.Val}
			mapNodes[cur] = newNode
		}

		randomNode, ok := mapNodes[cur.Random]
		if !ok {
			randomNode = &Node{Val: cur.Random.Val}
			mapNodes[cur.Random] = randomNode
		}
		newNode.Random = randomNode

		nextNode, ok := mapNodes[cur.Next]
		if !ok {
			nextNode = &Node{Val: cur.Next.Val}
			mapNodes[cur.Next] = nextNode
		}
		newNode.Next = nextNode

		newCur.Next = newNode
		newCur = newCur.Next
		cur = cur.Next
	}

	return newHead.Next
}

func CopyListWithRandomPointerRecursive(head *Node) *Node {
	visited := make(map[*Node]*Node)
	visited[nil] = nil

	var rec func(node *Node) *Node
	rec = func(node *Node) *Node {
		if clonedNode, ok := visited[node]; ok {
			return clonedNode
		}

		newNode := &Node{Val: node.Val}
		visited[node] = newNode

		newNode.Random = rec(node.Random)
		newNode.Next = rec(node.Next)
		return newNode
	}

	return rec(head)
}
