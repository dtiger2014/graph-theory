package graph

import "fmt"

// GraphDFS 图 深度优先遍历
type GraphDFS struct {
	g       *Graph
	visited []bool
	pre     []int
	post    []int
}

// Init 初始化，图深度优先遍历（DFS）
func (gdfs *GraphDFS) Init(g *Graph) error {
	// 赋值 初始化
	gdfs.g = g
	gdfs.visited = make([]bool, g.V())
	gdfs.pre = make([]int, 0)
	gdfs.post = make([]int, 0)

	// 遍历所有顶点
	var err error
	for v := 0; v < g.V(); v++ {
		if gdfs.visited[v] == false {
			err = gdfs.dfs(v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (gdfs *GraphDFS) dfs(v int) error {
	gdfs.visited[v] = true
	gdfs.pre = append(gdfs.pre, v)
	vertexs, err := gdfs.g.Adj(v)
	if err != nil {
		return err
	}
	for vertex := range vertexs {
		if gdfs.visited[vertex] == false {
			gdfs.dfs(vertex)
		}
	}
	gdfs.post = append(gdfs.post, v)

	return nil
}

// Pre 深度优先遍历 前序
func (gdfs *GraphDFS) Pre() []int {
	return gdfs.pre
}

// Post 深度优先遍历 后序
func (gdfs *GraphDFS) Post() []int {
	return gdfs.post
}

// String
func (gdfs *GraphDFS) String() string {
	res := ""

	// 第一行
	res += fmt.Sprintf("GraphDFS: V = %d, E = %d\n", gdfs.g.V(), gdfs.g.E())

	// 前序
	res += fmt.Sprintf("PreOrder: %v \n", gdfs.Pre())
	// 后续
	res += fmt.Sprintf("PostOrder: %v \n", gdfs.Post())

	return res
}
