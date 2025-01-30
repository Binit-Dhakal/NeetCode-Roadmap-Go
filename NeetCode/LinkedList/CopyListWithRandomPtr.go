package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyListWithRandomPtr(head *Node) *Node {
	hashNode := make(map[*Node]*Node)

	current := head

	for current != nil {
		hashNode[current] = &Node{Val: current.Val}

		current = current.Next
	}

	current = head

	for current != nil {
		hashNode[current].Next = hashNode[current.Next]
		hashNode[current].Random = hashNode[current.Random]

		current = current.Next
	}

	return hashNode[head]
}

func CopyListWithRandomPtrRecursion(head *Node) *Node {
	hashNode := make(map[*Node]*Node)

	var copyList func(*Node) *Node
	copyList = func(head *Node) *Node {
		if head == nil {
			return nil
		}

		if val, ok := hashNode[head]; ok {
			return val
		}

		hashNode[head] = &Node{Val: head.Val}
		hashNode[head].Next = copyList(head.Next)
		hashNode[head].Random = hashNode[head.Random]

		return hashNode[head]
	}

	return copyList(head)
}

func CopyListWithRandomPtrInterWeaving(head *Node) *Node {
	current := head
	for current != nil {
		current.Next = &Node{Val: current.Val, Next: current.Next}
		current = current.Next.Next
	}

	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}

		current = current.Next.Next
	}

	current = head
	var currentNew *Node
	for current != nil {
		currentNew = current.Next
		current.Next = currentNew.Next
		if currentNew.Next != nil {
			currentNew.Next = currentNew.Next.Next
		}
		current = current.Next
	}

	return head.Next
}
