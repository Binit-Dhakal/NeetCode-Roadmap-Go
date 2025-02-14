package linkedlist

func FlattenAMultiLevelDLL(root *Node) *Node {
	if root == nil {
		return nil
	}

	dummyHead := &Node{}
	curHead := dummyHead

	var traverseChild func(parent *Node)
	traverseChild = func(parent *Node) {
		if parent == nil {
			return
		}

		var nextNode *Node

		parent.Prev = curHead
		curHead.Next = parent
		curHead = parent

		nextNode = parent.Next

		if parent.Child != nil {
			traverseChild(parent.Child)
			parent.Child = nil
		}
		traverseChild(nextNode)
	}

	traverseChild(root)
	dummyHead.Next.Prev = nil
	return dummyHead.Next
}

func FlattenAMultiLevelDLLIterative(root *Node) *Node {
	if root == nil {
		return nil
	}

	dummyHead := &Node{}
	curHead := dummyHead

	nextNodes := make([]*Node, 0)
	nextNodes = append(nextNodes, root)

	for len(nextNodes) != 0 {
		nextNode := nextNodes[len(nextNodes)-1]
		curNode := nextNode
		nextNodes = nextNodes[:len(nextNodes)-1]

		curHead.Next = curNode
		curNode.Prev = curHead
		curHead = curNode
		if curNode.Next != nil {
			nextNodes = append(nextNodes, curNode.Next)
		}

		if curNode.Child != nil {
			nextNodes = append(nextNodes, curNode.Child)
			curNode.Child = nil
		}

	}

	dummyHead.Next.Prev = nil
	return dummyHead.Next
}
