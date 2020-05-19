package main

import (
	"fmt"
	"graph-theory/graph"
)

func main() {
	adjMatrix := new(graph.AdjMatrix)

	err := adjMatrix.Init("files/g.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(adjMatrix)
}
