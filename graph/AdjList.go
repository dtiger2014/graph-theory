package graph

import (
	"errors"
	"container/list"
	"fmt"
	"graph-theory/utils"
	"strconv"
	"strings"
)
/*
空间复杂度：
	- O(V+E):顶点数+边数
时间复杂图：
	- 建图：O(E*V)
	- 查看亮点是否相邻：O(degree(v))
	- 求一个点的相邻节点：O(degree(v))
优化：
	- 需要解决，建图、求一个点的相邻节点 操作性能
	- 不实用链表（LinkedList）
	- 使用哈希表HashSet O(1) : go语言直接使用map
	- 使用红黑树TreeSet O(logV) 
*/
type AdjList struct {
	v   int
	e   int
	adj []*list.List
}

// validateVertex 验证 顶点合法性
func (adj *AdjList) validateVertex(v int) error {
	if v < 0 || v >= adj.v {
		return errors.New(fmt.Sprintf("vertex %v is invalid", v))
	}
	return nil
}

// contains adj[v]链表中，包含w
func (adj *AdjList) contains(v, w int) bool {
	err := adj.validateVertex(v)
	if err != nil {
		return false
	}
	for i:= adj.adj[v].Front(); i != nil; i = i.Next() {
		if i.Value == w {
			return false
		}
	}
	return true
}

// Init 初始化 图
func (adj *AdjList) Init(filepath string) error {
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
	adj.adj = make([]*list.List, 0)
	for i := 0; i < adj.v; i ++ {
		adj.adj = append(adj.adj, list.New())
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

		if adj.contains(a, b) == false {
			return errors.New("Parallel Edges are Detected!")
		}

		adj.adj[a].PushBack(b)
		adj.adj[b].PushBack(a)
	}

	return nil
}

// V 返回 图 顶点数（vertex）
func (adj *AdjList) V() int {
	return adj.v
}

// E 返回 图 边数(edge)
func (adj *AdjList) E() int {
	return adj.e
}

// HasEdge 判断是否有边
func (adj *AdjList) HasEdge(v, w int) (bool, error) {
	err := adj.validateVertex(v)
	if err != nil {
		return false, err
	}
	err = adj.validateVertex(w)
	if err != nil {
		return false, err
	}

	return adj.contains(v, w), nil
}

// Adj 返回Adj
func (adj *AdjList) Adj(v int) (*list.List, error) {
	err := adj.validateVertex(v)
	if err != nil {
		return nil, err
	}

	return adj.adj[v], nil
}

// Degree 返回顶点的度
func (adj *AdjList) Degree(v int) (int, error) {
	vAdj, err := adj.Adj(v)
	if err != nil {
		return 0, err
	}
	return vAdj.Len(), nil
}

func (adj *AdjList) String() string {
	res := ""

	// 第一行
	res += fmt.Sprintf("V = %d, E = %d\n", adj.v, adj.e)
	// 其他行
	for v := 0; v < adj.v; v++ {
		res += fmt.Sprintf("%d : ", v)
		for w := adj.adj[v].Front(); w != nil; w = w.Next() {
			res += fmt.Sprintf("%d ", w.Value)
		}
		res += "\n"
	}
	return res
}
