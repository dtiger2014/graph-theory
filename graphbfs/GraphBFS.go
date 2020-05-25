package graphbfs

import (
	"graph-theory/graph"
)

// GraphBFS 图 广度优先遍历
type GraphBFS struct {
	g       *graph.Graph
	visited []bool
	order   []int
}

// Init 初始化
func (gbfs *GraphBFS) Init(g *graph.Graph) error {
	gbfs.g = g
	gbfs.visited = make([]bool, gbfs.g.V())
	gbfs.order = make([]int, 0)

	for v := 0; v < gbfs.g.V(); v++ {
		if gbfs.visited[v] == false {
			err := gbfs.bfs(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (gbfs *GraphBFS) bfs(s int) error {
	queue := []int{s}

	gbfs.visited[s] = true

	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]
		gbfs.order = append(gbfs.order, v)

		vertexs, err := gbfs.g.Adj(v)
		if err != nil {
			return err
		}

		for vertex := range vertexs {
			if gbfs.visited[vertex] == false {
				err := gbfs.bfs(vertex)
				if err != nil {
					return err
				}
				gbfs.visited[vertex] = true
			}
		}
	}
	return nil
}

// Order 输出 顶点
func (gbfs *GraphBFS) Order() []int {
	return gbfs.order
}
