package graphbfs

import "graph-theory/graph"

type CC struct {
	g       *graph.Graph
	visited []int
	cccount int
}

func (cc *CC) Init(g *graph.Graph) error {
	cc.g = g
	cc.visited = make([]int, cc.g.V())
	for idx := range cc.visited {
		cc.visited[idx] = -1
	}
	cc.cccount = 0

	for v := 0; v < cc.g.V(); v++ {
		if cc.visited[v] == -1 {
			err := cc.bfs(v, cc.cccount)
			if err != nil {
				return err
			}
			cc.cccount++
		}
	}
	return nil
}

func (cc *CC) bfs(v, ccid int) error {
	queue := []int{v}
	cc.visited[v] = ccid
	for len(queue) != 0 {
		v := queue[0]
		queue = queue[1:]

		vertexes, err := cc.g.Adj(v)
		if err != nil {
			return err
		}
		for w := range vertexes {
			if cc.visited[w] == -1 {
				queue = append(queue, w)
				cc.visited[w] = ccid
			}
		}
	}
	return nil
}

func (cc *CC) Count() int {
	return cc.cccount
}

func (cc *CC) IsConnected(v, w int) bool {
	if cc.g.ValidateVertex(v) != nil {
		return false
	}
	if cc.g.ValidateVertex(w) != nil {
		return false
	}
	return cc.visited[v] == cc.visited[w]
}

func (cc *CC) Components() [][]int {
	result := make([][]int, cc.cccount)
	for i := 0; i < cc.cccount; i++ {
		result[i] = make([]int, 0)
	}

	for v := 0; v < cc.g.V(); v++ {
		result[cc.visited[v]] = append(result[cc.visited[v]], v)
	}
	return result
}

