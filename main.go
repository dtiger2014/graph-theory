package main

import (
	"fmt"
	"graph-theory/graph"
)

func main() {
	var err error
	
	// graph
	g := new(graph.Graph)
	err = g.Init("files/g.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(g)

	// graphDFS
	gdfs := new(graph.GraphDFS)
	err = gdfs.Init(g)
	if err != nil {
		panic(err)
	}
	fmt.Println(gdfs)
}
