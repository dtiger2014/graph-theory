package graphbfs

import "graph-theory/graph"

// USSSPath Unweighted Single Source Shortest Path 无权图单原 最短路径
type USSSPath struct {
	g       *graph.Graph
	s       int
	visited []bool
	pre     []int
	dis     []int
}

// Init 初始化
func (up *USSSPath) Init(g *graph.Graph, s int) error {
	up.g = g
	up.s = s
	up.visited = make([]bool, up.g.V())
	up.pre = make([]int, up.g.V())
	up.dis = make([]int, up.g.V())
	for i := 0; i < up.g.V(); i++ {
		up.pre[i] = -1
		up.dis[i] = -1
	}

	return up.bfs(s)
}

func (up *USSSPath) bfs(s int) error {
	queue := []int{s}
	up.visited[s] = true
	up.pre[s] = s
	up.dis[s] = 0

	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		vertexes, err := up.g.Adj(v)
		if err != nil {
			return err
		}

		for w := range vertexes {
			if up.visited[w] == false {
				up.visited[w] = true
				up.pre[w] = v
				up.dis[w] = up.dis[v] + 1
			}
		}
	}
	return nil
}

// IsConnectedTo 判断是否连接
func (up *USSSPath) IsConnectedTo(t int) bool {
	if up.g.ValidateVertex(t) != nil {
		return false
	}
	return up.visited[t]
}

// Dis 距离
func (up *USSSPath) Dis(t int) int {
	if up.g.ValidateVertex(t) != nil {
		return -1
	}
	return up.dis[t]
}

// Path 路径
func (up *USSSPath) Path(t int) []int {
	res := []int{}
	if up.IsConnectedTo(t) == false {
		return res
	}

	cur := t
	for cur != up.s {
		res = append(res, cur)
		cur = up.pre[cur]
	}
	res = append(res, up.s)

	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return res
}
