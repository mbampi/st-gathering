package creator

import (
	"stgathering/pkgs/st"
)

func CreateSTTree(numNodes int) *st.Tree {
	nodes := make([]*st.Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = st.NewNode(i)
	}

	assignSTChildren(nodes, numNodes)

	tree := st.NewTree(nodes, 0)
	return tree
}

func assignSTChildren(nodes []*st.Node, numNodes int) {
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
