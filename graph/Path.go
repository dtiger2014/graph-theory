package graph

// Path 路径
type Path struct {
	g       *Graph
	s       int
	t       int
	pre     []int
	visited []bool
}

// Init 初始化
func (p *Path) Init(g *Graph, s, t int) error {
	var err error

	err = g.validateVertex(s)
	if err != nil {
		return err
	}
	err = g.validateVertex(t)
	if err != nil {
		return err
	}

	p.g = g
	p.s = s
	p.t = t

	p.pre = make([]int, p.g.V())
	for k := range p.pre {
		p.pre[k] = -1
	}
	p.visited = make([]bool, p.g.V())

	p.dfs(p.s, p.s)

	return nil
}

func (p *Path) dfs(v, parent int) bool {
	p.visited[v] = true
	p.pre[v] = parent

	if v == p.t {
		return true
	}

	vertexs, err := p.g.Adj(v)
	if err != nil {
		return false
	}
	for vertex := range vertexs {
		if p.visited[vertex] == false {
			if p.dfs(vertex, v) == true {
				return true
			}
		}
	}
	return false
}

// IsConnected 判断 s与t 是否连接
func (p *Path) IsConnected() bool {
	return p.visited[p.t]
}

// Path 返回路径
func (p *Path) Path() []int {
	res := []int{}
	if p.IsConnected() == false {
		return res
	}

	cur := p.t
	for cur != p.s {
		res = append(res, cur)
		cur = p.pre[cur]
	}
	res = append(res, p.s)

	// 翻转res
	n := len(res)
	for i := 0; i < n/2; i++ {
		res[i], res[n-1-i] = res[n-1-i], res[i]
	}

	return res
}