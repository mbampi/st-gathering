package main

import (
	"fmt"
	"os"
	"stgathering/pkgs/st_mod"
	"stgathering/pkgs/stats"
	"time"
)

func main() {
	// Get number of nodes from command line arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <numNodes>")
		os.Exit(1)
	}

	numNodes := 0
	_, err := fmt.Sscanf(os.Args[1], "%d", &numNodes)
	if err != nil {
		fmt.Println("Usage: go run main.go <numNodes>")
		os.Exit(1)
	}

	nodes := make([]*st_mod.Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = st_mod.NewNode(i)
	}

	// Create tree

	// Assign children to nodes
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

	fmt.Println(updateTotalChildren(nodes[0]))

	tree := st_mod.NewTree(nodes, 0)

	startTime := time.Now()
	tree.Run()
	duration := time.Since(startTime)

	fmt.Printf("Total messages sent: %d\n", stats.GetNumMessagesSent())
	fmt.Printf("Total time: %s\n", duration)
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
