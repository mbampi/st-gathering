package main

import (
	"fmt"
	"os"
	"stgathering/pkgs/st_mod"
	"stgathering/pkgs/stats"
	"strconv"
	"time"
)

func main() {
	numNodes := getNumNodesFromArgs()

	nodes := make([]*st_mod.Node, numNodes)
	for i := 0; i < numNodes; i++ {
		nodes[i] = st_mod.NewNode(i)
	}

	assignChildren(nodes, numNodes)

	updateTotalChildren(nodes[0])

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

func assignChildren(nodes []*st_mod.Node, numNodes int) {
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

func getNumNodesFromArgs() int {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <numNodes>")
		os.Exit(1)
	}

	numNodes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Usage: go run main.go <numNodes>")
		os.Exit(1)
	}

	return numNodes
}
