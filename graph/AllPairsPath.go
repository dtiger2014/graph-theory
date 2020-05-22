package graph

// AllPairsPath 创建所有顶点路径
type AllPairsPath struct {
	g     *Graph
	paths []*SingleSourcePath
}

// Init 初始化
func (ap *AllPairsPath) Init(g *Graph) error {
	ap.g = g
	ap.paths = make([]*SingleSourcePath, 0)
	for v := 0; v < ap.g.V(); v++ {
		ss := new(SingleSourcePath)
		err := ss.Init(ap.g, v)
		if err != nil {
			return err
		}
		ap.paths = append(ap.paths, ss)
	}
	return nil
}

// IsConnectedTo 判断两点是否连接
func (ap *AllPairsPath) IsConnectedTo(s int, t int) bool {
	if ap.g.validateVertex(s) != nil {
		return false
	}
	return ap.paths[s].IsConnectedTo(t)
}

// Path 返回两点 路径
func (ap *AllPairsPath) Path(s int, t int) []int {
	if ap.g.validateVertex(s) != nil {
		return nil
	}
	return ap.paths[s].Path(t)
}
