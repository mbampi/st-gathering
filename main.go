package main

import (
	"fmt"
	"os"
	"stgathering/pkgs/creator"
	"stgathering/pkgs/st"
	"stgathering/pkgs/stats"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <st|st_mod> <numNodes> ")
		os.Exit(1)
	}

	treeType := os.Args[1]

	numNodes, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Usage: go run main.go <st|st_mod>  <numNodes>")
		os.Exit(1)
	}

	var tree st.ITree
	if treeType == "st_mod" {
		tree = creator.CreateSTModTree(numNodes)
	} else {
		tree = creator.CreateSTTree(numNodes)
	}

	startTime := time.Now()
	tree.Run()
	duration := time.Since(startTime)

	fmt.Printf("Total messages sent: %d\n", stats.GetNumMessagesSent())
	fmt.Printf("Total time: %s\n", duration)
}
