package dfs

import "io.github.sowen.datastructure/graph"

type BipartitionDetection struct {
	graph.Graph
	visited     []bool
	colors      []int
	isBipartite bool
}

func NewBipartitionDetection(g graph.Graph) *BipartitionDetection {
	colors := make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		colors[v] = -1
	}
	res := &BipartitionDetection{
		Graph:       g,
		visited:     make([]bool, g.V()),
		colors:      colors,
		isBipartite: true,
	}
	for v := 0; v < g.V(); v++ {
		if !res.visited[v] {
			if !res.dfs(v, 0) {
				res.isBipartite = false
				break
			}
		}
	}
	return res
}

// dfs  初始化所有 vertex 的 color 为 -1
//
//	用 0 和 1 表示两种 color
func (c *BipartitionDetection) dfs(v int, color int) bool {
	c.visited[v] = true
	c.colors[v] = color
	for _, w := range c.Graph.Adj(v) {
		if !c.visited[w] {
			if !c.dfs(w, 1-color) {
				return false
			}
		} else if c.colors[w] == color {
			return false
		}
	}
	return true
}

func (c *BipartitionDetection) IsBipartite() bool {
	return c.isBipartite
}
