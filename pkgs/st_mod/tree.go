package st_mod

import (
	"fmt"
	"sync"
)

type Tree struct {
	root  *Node
	nodes []*Node
}

// NewTree creates a new Tree
func NewTree(nodes []*Node, rootID int) *Tree {
	root := nodes[rootID]
	return &Tree{
		root:  root,
		nodes: nodes,
	}
}

func (t *Tree) Print() {
	t.printTreeRec(t.root, "", true)
}

// PrintTree prints the structure of the tree starting from the node
func (t *Tree) printTreeRec(node *Node, prefix string, isTail bool) {
	// Determine the line connector based on whether the node is the last child
	connector := "└── "
	if !isTail {
		connector = "├── "
	}

	// Print the current node's ID
	fmt.Println(prefix + connector + fmt.Sprintf("Node ID: %d", node.ID))

	// Prepare the next level's prefix
	nextPrefix := prefix
	if isTail {
		nextPrefix += "    "
	} else {
		nextPrefix += "│   "
	}

	// Recursively print each child
	for i, child := range node.Children {
		t.printTreeRec(child, nextPrefix, i == len(node.Children)-1)
	}
}

// Run starts the goroutine for the tree
func (t *Tree) Run() {
	var wg sync.WaitGroup

	// Run nodes
	for _, node := range t.nodes {
		wg.Add(1)
		go node.Run(&wg)
	}

	wg.Wait()
}
