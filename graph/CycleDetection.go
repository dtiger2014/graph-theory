package graph

// CycleDetection 无向图 检测环
type CycleDetection struct {
	g        *Graph
	visited  []bool
	hasCycle bool
}

// Init 初始化
func (cd *CycleDetection) Init(g *Graph) error {
	cd.g = g
	cd.visited = make([]bool, cd.g.V())

	for v := 0; v < cd.g.V(); v++ {
		if cd.visited[v] == false {
			res, err := cd.dfs(v, v)
			if err != nil {
				return err
			}
			if res == true {
				cd.hasCycle = true
				break
			}
		}
	}
	return nil
}

func (cd *CycleDetection) dfs(v, parent int) (bool, error) {
	cd.visited[v] = true

	vertexs, err := cd.g.Adj(v)
	if err != nil {
		return false, err
	}

	for vertex := range vertexs {
		if cd.visited[vertex] == false {
			res, err := cd.dfs(vertex, v)
			if err != nil {
				return false, err
			}
			if res == true {
				return true, nil
			}
		} else if vertex != parent {
			return true, nil
		}
	}
	return false, nil
}

// HasCycle 检测是否有环
func (cd *CycleDetection) HasCycle() bool {
	return cd.hasCycle
}
