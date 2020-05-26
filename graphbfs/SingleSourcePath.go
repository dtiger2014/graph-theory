package graphbfs

import "graph-theory/graph"

// SingleSourcePath 图 单路径
type SingleSourcePath struct {
	g       *graph.Graph
	s       int
	visited []bool
	pre     []int
}

// Init 初始化，图广度优先遍历（BFS）
func (ss *SingleSourcePath) Init(g *graph.Graph, s int) error {
	// 赋值 初始化
	ss.g = g
	ss.s = s // 源source
	ss.visited = make([]bool, ss.g.V())
	ss.pre = make([]int, ss.g.V())
	for k := range ss.pre {
		ss.pre[k] = -1
	}

	// 遍历所有顶点
	var err error
	err = ss.bfs(s)
	if err != nil {
		return err
	}
	return nil
}

func (ss *SingleSourcePath) bfs(s int) error {
	queue := []int{s}
	ss.visited[s] = true
	ss.pre[s] = s
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		vertexs, err := ss.g.Adj(v)
		if err != nil {
			return err
		}

		for w := range vertexs {
			if ss.visited[w] == false {
				queue = append(queue, w)
				ss.visited[w] = true
				ss.pre[w] = v
			}
		}
	}
	return nil
}

// IsConnectedTo 判断顶点是否与source 连接
func (ss *SingleSourcePath) IsConnectedTo(t int) bool {
	if ss.g.ValidateVertex(t) != nil {
		return false
	}
	return ss.visited[t]
}

// Path 返回 source 与目标顶点t 路径
func (ss *SingleSourcePath) Path(t int) []int {
	res := []int{}
	if ss.IsConnectedTo(t) == false {
		return res
	}

	cur := t
	for cur != ss.s {
		res = append(res, cur)
		cur = ss.pre[cur]
	}
	res = append(res, ss.s)

	// 翻转res
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return res
}
