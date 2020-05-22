package graph

// BipartitionDetection 二分图检测
type BipartitionDetection struct {
	g           *Graph
	visited     []bool
	colors      []int
	isBipartite bool
}

func (bd *BipartitionDetection) Init(g *Graph) error {
	bd.g = g
	bd.visited = make([]bool, bd.g.V())
	bd.colors = make([]int, bd.g.V())
	for k := range bd.colors {
		bd.colors[k] = -1
	}
	bd.isBipartite = true

	for v := 0; v < bd.g.V(); v++ {
		if bd.visited[v] == false {
			if bd.dfs(v, 0) {
				bd.isBipartite = false
				break
			}
		}
	}
	return nil
}

func (bd *BipartitionDetection) dfs (v, color int) bool {
	bd.
}