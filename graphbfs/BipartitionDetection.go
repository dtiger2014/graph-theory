package graphbfs

import "graph-theory/graph"

// BipartitionDetection 二分图检测
type BipartitionDetection struct {
	g           *graph.Graph
	visited     []bool
	colors      []int
	isBipartite bool
}

// Init 初始化
func (bd *BipartitionDetection) Init(g *graph.Graph) error {
	bd.g = g
	bd.visited = make([]bool, bd.g.V())
	bd.colors = make([]int, bd.g.V())
	for i := 0; i < bd.g.V(); i++ {
		bd.colors[i] = -1
	}
	bd.isBipartite = true

	for v := 0; v < bd.g.V(); v++ {
		if bd.visited[v] == false {
			if bd.bfs(v) == false {
				bd.isBipartite = false
				break
			}
		}
	}
	return nil
}

func (bd *BipartitionDetection) bfs(s int) bool {
	queue := []int{s}
	bd.visited[s] = true
	bd.colors[s] = 0

	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		vertexes, err := bd.g.Adj(v)
		for err != nil {
			return false
		}

		for w := range vertexes {
			if bd.visited[w] == false {
				queue = append(queue, w)
				bd.visited[w] = true
				bd.colors[w] = 1 - bd.colors[v]
			} else if bd.colors[v] == bd.colors[w] {
				return false
			}
		}
	}
	return true
}

// IsBipartite 判断是否为二分图
func (bd *BipartitionDetection) IsBipartite() bool {
	return bd.isBipartite
}
