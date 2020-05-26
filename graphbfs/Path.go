package graphbfs

import "graph-theory/graph"

// Path 图 路径
type Path struct {
	g       *graph.Graph
	s       int
	t       int
	visited []bool
	pre     []int
}

// Init 初始化，图广度优先遍历（BFS）
func (path *Path) Init(g *graph.Graph, s, t int) error {
	// 赋值 初始化
	path.g = g
	path.s = s // 源source
	path.t = t
	path.visited = make([]bool, path.g.V())
	path.pre = make([]int, path.g.V())
	for k := range path.pre {
		path.pre[k] = -1
	}

	// 遍历所有顶点
	var err error
	err = path.bfs()
	if err != nil {
		return err
	}
	return nil
}

func (path *Path) bfs() error {
	queue := []int{path.s}
	path.visited[path.s] = true
	path.pre[path.s] = path.s
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		vertexs, err := path.g.Adj(v)
		if err != nil {
			return err
		}

		for w := range vertexs {
			if path.visited[w] == false {
				queue = append(queue, w)
				path.visited[w] = true
				path.pre[w] = v

				if w == path.t {
					return nil
				}
			}
		}
	}
	return nil
}

// IsConnectedTo 判断顶点是否与source 连接
func (path *Path) IsConnected() bool {
	return path.visited[path.t]
}

// Path 返回 source 与目标顶点t 路径
func (path *Path) Path() []int {
	res := []int{}
	if path.IsConnected() == false {
		return res
	}

	cur := path.t
	for cur != path.s {
		res = append(res, cur)
		cur = path.pre[cur]
	}
	res = append(res, path.s)

	// 翻转res
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}
	return res
}
