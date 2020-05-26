package graphbfs

import "graph-theory/graph"

type AllPairsPath struct {
	g     *graph.Graph
	paths []*SingleSourcePath
}

// Init 初始化
func (ap *AllPairsPath) Init(g *graph.Graph) error {
	ap.g = g
	ap.paths = make([]*SingleSourcePath, ap.g.V())
	for v := 0; v < ap.g.V(); v++ {
		path := new(SingleSourcePath)
		err := path.Init(ap.g, v)
		if err != nil {
			return err
		}
		ap.paths[v] = path
	}

	return nil
}

// IsConntectedTo 是否连接？
func (ap *AllPairsPath) IsConntectedTo(s, t int) bool {
	if ap.g.ValidateVertex(s) != nil {
		return false
	}
	return ap.paths[s].IsConnectedTo(t)
}

// Path 返回 s->t 路径
func (ap *AllPairsPath) Path(s, t int) []int {
	if ap.g.ValidateVertex(s) != nil {
		return nil
	}
	return ap.paths[s].Path(t)
}
