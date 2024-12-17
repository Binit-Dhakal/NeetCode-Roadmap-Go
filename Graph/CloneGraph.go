package graph

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	nodeMap := make(map[*Node]*Node, 0) // oldNode: newNode

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		if _, ok := nodeMap[node]; ok {
			return
		}

		newNode := &Node{Val: node.Val}
		nodeMap[node] = newNode

		neighbors := node.Neighbors
		for _, neighbor := range neighbors {
			dfs(neighbor)

			newNode.Neighbors = append(newNode.Neighbors, nodeMap[neighbor])
		}
	}

	dfs(node)

	return nodeMap[node]
}
