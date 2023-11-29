package creator

import "stgathering/pkgs/st_mod"

func CreateSTModTree(numNodes int) *st_mod.Tree {
	nodes := make([]*st_mod.Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = st_mod.NewNode(i)
	}

	assignSTModChildren(nodes, numNodes)

	updateTotalChildren(nodes[0])

	tree := st_mod.NewTree(nodes, 0)
	return tree
}

func assignSTModChildren(nodes []*st_mod.Node, numNodes int) {
	for i := 0; i < numNodes; i++ {
		leftChildIndex := 2*i + 1
		rightChildIndex := 2*i + 2

		if leftChildIndex < numNodes {
			nodes[i].AddChild(nodes[leftChildIndex])
		}
		if rightChildIndex < numNodes {
			nodes[i].AddChild(nodes[rightChildIndex])
		}
	}
}

func updateTotalChildren(node *st_mod.Node) int {
	if node == nil {
		return 0
	}

	totalChildren := 0

	// Calculate total children for each child node
	for _, child := range node.Children {
		totalChildren += updateTotalChildren(child) + 1
	}

	// Update the total children count for the current node
	node.TotalChildren = totalChildren

	return totalChildren
}
