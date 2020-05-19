package graph

import (
	"errors"
	"fmt"
	"graph-theory/utils"
	"strconv"
	"strings"
)

type AdjMatrix struct {
	v   int
	e   int
	adj [][]int
}

// validateVertex 验证 顶点合法性
func (adj *AdjMatrix) validateVertex(v int) error {
	if v < 0 || v >= adj.v {
		return errors.New(fmt.Sprintf("vertex %v is invalid", v))
	}
	return nil
}

// Init 初始化 图
func (adj *AdjMatrix) Init(filepath string) error {
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
	adj.adj = make([][]int, adj.v)
	for k := range adj.adj {
		adj.adj[k] = make([]int, adj.v)
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

		if adj.adj[a][b] == 1 {
			return errors.New("Parallel Edges are Detected!")
		}

		adj.adj[a][b] = 1
		adj.adj[b][a] = 1
	}

	return nil
}

// V 返回 图 顶点数（vertex）
func (adj *AdjMatrix) V() int {
	return adj.v
}

// E 返回 图 边数(edge)
func (adj *AdjMatrix) E() int {
	return adj.e
}

// HasEdge 判断是否有边
func (adj *AdjMatrix) HasEdge(v, w int) (bool, error) {
	err := adj.validateVertex(v)
	if err != nil {
		return false, err
	}
	err = adj.validateVertex(w)
	if err != nil {
		return false, err
	}

	return adj.adj[v][w] == 1, nil
}

// Adj 返回Adj
func (adj *AdjMatrix) Adj(v int) ([]int, error) {
	err := adj.validateVertex(v)
	if err != nil {
		return nil, err
	}

	res := []int{}
	for k, val := range adj.adj[v] {
		if val == 1 {
			res = append(res, k)
		}
	}

	return res, nil
}

// Degree 返回顶点的度
func (adj *AdjMatrix) Degree(v int) (int, error) {
	vAdj, err := adj.Adj(v)
	if err != nil {
		return 0, err
	}
	return len(vAdj), nil
}

func (adj *AdjMatrix) String() string {
	res := ""

	// 第一行
	res += fmt.Sprintf("V = %d, E = %d\n", adj.v, adj.e)
	for i := 0; i < len(adj.adj); i++ {
		for j := 0; j < len(adj.adj[i]); j++ {
			res += fmt.Sprintf("%d ", adj.adj[i][j])
		}
		res += "\n"
	}
	return res
}
