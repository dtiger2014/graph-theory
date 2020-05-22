package graph

import "fmt"

// SingleSourcePath 图 单路径
type SingleSourcePath struct {
	g       *Graph
	s       int
	visited []bool
	pre     []int
}

// Init 初始化，图深度优先遍历（DFS）
func (ss *SingleSourcePath) Init(g *Graph, s int) error {
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
	err = ss.dfs(s, s)
	if err != nil {
		return err
	}
	return nil
}

func (ss *SingleSourcePath) dfs(v int, parent int) error {
	ss.visited[v] = true
	ss.pre[v] = parent

	vertexs, err := ss.g.Adj(v)
	if err != nil {
		return err
	}
	for vertex := range vertexs {
		if ss.visited[vertex] == false {
			err = ss.dfs(vertex, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// IsConnectedTo 判断顶点是否与source 连接
func (ss *SingleSourcePath) IsConnectedTo(t int) bool {
	if ss.g.validateVertex(t) != nil {
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

// String
func (ss *SingleSourcePath) String() string {
	res := ""

	res += fmt.Sprintf("SingleSourcePath: V = %d, E = %d\n", ss.g.V(), ss.g.E())
	return res
}
