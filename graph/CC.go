package graph

import "fmt"

// CC 图 联通分量
type CC struct {
	g       *Graph
	visited []int
	cccount int
}

// Init 初始化，图深度优先遍历（DFS）
func (cc *CC) Init(g *Graph) error {
	// 赋值 初始化
	cc.g = g
	cc.visited = make([]int, g.V())
	for k := range cc.visited {
		cc.visited[k] = -1
	}
	cc.cccount = 0

	// 遍历所有顶点
	var err error
	for v := 0; v < g.V(); v++ {
		if cc.visited[v] == -1 {
			err = cc.dfs(v, cc.cccount)
			if err != nil {
				return err
			}
			cc.cccount++
		}
	}
	return nil
}

func (cc *CC) dfs(v int, ccid int) error {
	cc.visited[v] = ccid
	vertexs, err := cc.g.Adj(v)
	if err != nil {
		return err
	}
	for vertex := range vertexs {
		if cc.visited[vertex] == -1 {
			cc.dfs(vertex, ccid)
		}
	}
	return nil
}

// Count 获取联通分量数量
func (cc *CC) Count() int {
	return cc.cccount
}

// IsConnected 判断两点之间连接
func (cc *CC) IsConnected(v, w int) bool {
	if cc.g.validateVertex(v) != nil {
		return false
	}
	if cc.g.validateVertex(w) != nil {
		return false
	}
	return cc.visited[v] == cc.visited[w]
}

// String
func (cc *CC) String() string {
	res := ""

	res += fmt.Sprintf("CC: V = %d, E = %d\n", cc.g.V(), cc.g.E())
	res += fmt.Sprintf("CCount: %d\n", cc.cccount)

	return res
}
