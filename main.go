package main

import (
	"fmt"
	"graph-theory/graph"
	"graph-theory/graphbfs"
)

func main() {
	test_bfs()
}

func test_bfs() {
	var err error

	// graph bfs
	g := new(graph.Graph)
	err = g.Init("files/g.txt")
	checkErr(err)
	fmt.Println(g)

	gbfs := new(graphbfs.GraphBFS)
	err = gbfs.Init(g)
	checkErr(err)
	fmt.Println(gbfs.Order())

	// Single Source Path
	fmt.Println("\nSingle Source Path ===")
	ss := new(graphbfs.SingleSourcePath)
	err = ss.Init(g, 0)
	checkErr(err)
	fmt.Println(ss.IsConnectedTo(6))
	fmt.Println(ss.Path(6))

	// Path
	fmt.Println("\nPath ====")
	path := new(graphbfs.Path)
	err = path.Init(g, 0, 5)
	checkErr(err)
	fmt.Println(path.IsConnected(), path.Path())

	// AllPairsPath
	fmt.Println("\nAll Pairs Path ====")
	ap := new(graphbfs.AllPairsPath)
	err = ap.Init(g)
	checkErr(err)
	fmt.Println(ap.IsConntectedTo(0, 4), ap.Path(0, 4))
	fmt.Println(ap.IsConntectedTo(3, 6), ap.Path(3, 6))
	fmt.Println(ap.IsConntectedTo(2, 5), ap.Path(2, 5))

	// CC
	fmt.Println("\nCC ====")
	cc := new(graphbfs.CC)
	err = cc.Init(g)
	checkErr(err)
	fmt.Println(cc.IsConnected(2, 5), cc.IsConnected(0, 6))
	fmt.Println(cc.Count())
	fmt.Println(cc.Components())

	// CycleDetection
	fmt.Println("\nCycle Detection ====")
	g2 := new(graph.Graph)
	err = g2.Init("files/g2.txt")
	checkErr(err)
	fmt.Println(g2)

	cd := new(graphbfs.CycleDetection)
	err = cd.Init(g2)
	checkErr(err)
	fmt.Println(cd.HasCycle())

	// graph 3
	g3 := new(graph.Graph)
	err = g3.Init("files/g3.txt")
	checkErr(err)
	fmt.Println(g3)

	// BipartitionDetection
	fmt.Println("\nBipartition Detection ====")
	bd := new(graphbfs.BipartitionDetection)
	err = bd.Init(g3)
	checkErr(err)
	fmt.Println(bd.IsBipartite())
}

func test_dfs() {
	var err error

	// graph
	g := new(graph.Graph)
	err = g.Init("files/g.txt")
	checkErr(err)
	fmt.Println(g)

	// graphDFS
	gdfs := new(graph.GraphDFS)
	err = gdfs.Init(g)
	checkErr(err)
	fmt.Println(gdfs)

	// CC 联通变量
	cc := new(graph.CC)
	err = cc.Init(g)
	checkErr(err)
	fmt.Println(cc)

	// SS 单源路径
	ss := new(graph.SingleSourcePath)
	err = ss.Init(g, 0)
	checkErr(err)
	fmt.Println(ss)
	fmt.Println(ss.Path(6))
	fmt.Println(ss.Path(5))

	// AllPairsPath
	ap := new(graph.AllPairsPath)
	err = ap.Init(g)
	checkErr(err)
	fmt.Println(ap.IsConnectedTo(3, 4))
	fmt.Println(ap.IsConnectedTo(2, 5))
	fmt.Println(ap.Path(3, 6))
	fmt.Println(ap.Path(2, 5))

	// Path
	p := new(graph.Path)
	err = p.Init(g, 3, 6)
	checkErr(err)
	fmt.Println(p.IsConnected())
	fmt.Println(p.Path())

	// CycleDetection
	// graph
	g2 := new(graph.Graph)
	err = g2.Init("files/g2.txt")
	checkErr(err)
	fmt.Println(g2)

	cd := new(graph.CycleDetection)
	err = cd.Init(g2)
	checkErr(err)
	fmt.Println(cd.HasCycle())

	// graph 3
	g3 := new(graph.Graph)
	err = g3.Init("files/g3.txt")
	checkErr(err)
	fmt.Println(g3)

	// BipartitionDetection
	bd := new(graph.BipartitionDetection)
	err = bd.Init(g3)
	checkErr(err)
	fmt.Println(bd.IsBipartite())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
	return
}
