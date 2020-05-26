package graphbfs

import "graph-theory/graph"

// CycleDetection 图 环检测
type CycleDetection struct {
	g        *graph.Graph
	visited  []bool
	pre      []int
	hasCycle bool
}

// Init 初始化
func (cd *CycleDetection) Init(g *graph.Graph) error {
	cd.g = g
	cd.visited = make([]bool, cd.g.V())
	cd.pre = make([]int, cd.g.V())
	for i := 0; i < cd.g.V(); i++ {
		cd.pre[i] = -1
	}
	cd.hasCycle = false

	for v := 0; v < cd.g.V(); v++ {
		if cd.visited[v] == false {
			if cd.bfs(v) == true {
				cd.hasCycle = true;
				break
			}
		}
	}
	return nil
}

func (cd *CycleDetection) bfs(s int) bool {
	queue := []int{s}
	cd.visited[s] = true
	cd.pre[s] = s
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		vertexes, err := cd.g.Adj(v)
		if err != nil {
			return false
		}

		for w := range vertexes {
			if cd.visited[w] == false {
				queue = append(queue,w)
				cd.visited[w] = true
				cd.pre[w] = v
			} else if cd.pre[v] != w {
				return true
			}
  		}
	}
	return false
}

// HasCycle 检测是否有环
func (cd *CycleDetection) HasCycle() bool {
	return cd.hasCycle
}