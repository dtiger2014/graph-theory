package graph

import (
	"errors"
	"fmt"
	"graph-theory/utils"
	"strconv"
	"strings"
)

/*
使用 set（map）进行优化

空间复杂度：
	- O(V+E):顶点数+边数
时间复杂图：
	- 建图：O(E*V)
	- 查看亮点是否相邻：O(degree(v))
	- 求一个点的相邻节点：O(degree(v))
*/
type AdjSet struct {
	v   int
	e   int
	adj []map[int]int
}

// validateVertex 验证 顶点合法性
func (adj *AdjSet) validateVertex(v int) error {
	if v < 0 || v >= adj.v {
		return errors.New(fmt.Sprintf("vertex %v is invalid", v))
	}
	return nil
}

// Init 初始化 图
func (adj *AdjSet) Init(filepath string) error {
	fContents, err := utils.ReadFile(filepath)
	if err != nil {
		return err
	}

	// 第一行：V，E
	hline := strings.Split(fContents[0], " ")
	vNum, err := strconv.Atoi(hline[0])
	if err != nil {
		return errors.New("V must be Integer")
	}
	adj.v = vNum

	eNum, err := strconv.Atoi(hline[1])
	if err != nil {
		return errors.New("E must be Integer")
	}
	adj.e = eNum

	// 初始化 二维数组
	adj.adj = make([]map[int]int, 0)
	for i := 0; i < adj.v; i++ {
		adj.adj = append(adj.adj, make(map[int]int))
	}

	// 顶点，边
	for i := 1; i < len(fContents); i++ {
		line := strings.Split(fContents[i], " ")
		a, err := strconv.Atoi(line[0])
		if err != nil {
			return errors.New(fmt.Sprintf("file %v line vertex %v must be Integer", i, line[0]))
		}
		err = adj.validateVertex(a)
		if err != nil {
			return err
		}

		b, err := strconv.Atoi(line[1])
		if err != nil {
			return errors.New(fmt.Sprintf("file %v line vertex %v must be Integer", i, line[1]))
		}
		err = adj.validateVertex(b)
		if err != nil {
			return err
		}

		if a == b {
			return errors.New("Self Loop is Detected!")
		}

		if _, ok := adj.adj[a][b]; ok {
			return errors.New("Parallel Edges are Detected!")
		}

		adj.adj[a][b] = 1
		adj.adj[b][a] = 1
	}

	return nil
}

// V 返回 图 顶点数（vertex）
func (adj *AdjSet) V() int {
	return adj.v
}

// E 返回 图 边数(edge)
func (adj *AdjSet) E() int {
	return adj.e
}

// HasEdge 判断是否有边
func (adj *AdjSet) HasEdge(v, w int) (bool, error) {
	err := adj.validateVertex(v)
	if err != nil {
		return false, err
	}
	err = adj.validateVertex(w)
	if err != nil {
		return false, err
	}

	_, ok := adj.adj[v][w]
	return ok, nil
}

// Adj 返回Adj
func (adj *AdjSet) Adj(v int) (map[int]int, error) {
	err := adj.validateVertex(v)
	if err != nil {
		return nil, err
	}

	return adj.adj[v], nil
}

// Degree 返回顶点的度
func (adj *AdjSet) Degree(v int) (int, error) {
	vAdj, err := adj.Adj(v)
	if err != nil {
		return 0, err
	}
	return len(vAdj), nil
}

func (adj *AdjSet) String() string {
	res := ""

	// 第一行
	res += fmt.Sprintf("V = %d, E = %d\n", adj.v, adj.e)
	// 其他行
	for v := 0; v < adj.v; v++ {
		res += fmt.Sprintf("%d : ", v)
		for w, _ := range adj.adj[v] {
			res += fmt.Sprintf("%d ", w)
		}
		res += "\n"
	}
	return res
}
