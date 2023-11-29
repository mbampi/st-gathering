package st

import (
	"fmt"
	"stgathering/pkgs/stats"
	"sync"
)

// Node represents a node in the spanning tree
type Node struct {
	ID            int
	Children      []*Node
	Parent        chan<- int
	Data          int
	ReceivingChan chan int
}

// NewNode creates a new Node
func NewNode(id int) *Node {
	return &Node{
		ID:            id,
		Children:      make([]*Node, 0),
		Parent:        nil,
		Data:          0,
		ReceivingChan: make(chan int, 200),
	}
}

// AddChild adds a child to the node
func (n *Node) AddChild(child *Node) {
	child.Parent = n.ReceivingChan
	n.Children = append(n.Children, child)
}

// Run starts the goroutine for the node
func (n *Node) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate data gathering
	n.Data = 1

	// Collect children data
	fmt.Printf("Node %d collecting data from its children (%d)\n", n.ID, len(n.Children))
	childData := 0
	for range n.Children {
		childData += <-n.ReceivingChan
	}
	fmt.Printf("Node %d collected %d from its children\n", n.ID, childData)

	// Aggregate child data
	n.Data += childData

	// Send data to parent
	if n.Parent != nil {
		fmt.Printf("Node %d sent %d to its parent\n", n.ID, n.Data)
		n.Parent <- n.Data
		stats.IncrementMessagesSent()
	}

	// If root, print the result
	if n.Parent == nil {
		fmt.Printf("Total data gathered by the root (Node %d): %d\n", n.ID, n.Data)
	}
}
