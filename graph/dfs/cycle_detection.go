package dfs

import "io.github.sowen.datastructure/graph"

type CycleDetection struct {
	graph.Graph
	visited  []bool
	hasCycle bool
}

func NewCycleDetection(g graph.Graph) *CycleDetection {
	if g.IsDirected() {
		panic("CycleDetection only works in undirected graph.")
	}

	res := &CycleDetection{
		Graph:   g,
		visited: make([]bool, g.V()),
	}
	for v := 0; v < g.V(); v++ {
		if !res.visited[v] && res.dfs(v, v) {
			res.hasCycle = true
			break
		}
	}
	return res
}

// dfs 从顶点 v 开始判断是否有环
// 从 0 → 1 后再从 1 → 0 不能算环
func (c *CycleDetection) dfs(v int, parent int) bool {
	c.visited[v] = true
	for _, w := range c.Graph.Adj(v) {
		if !c.visited[w] {
			if c.dfs(w, v) {
				return true
			}
		} else if w != parent {
			return true
		}
	}
	return false
}

func (c *CycleDetection) HasCycle() bool {
	return c.hasCycle
}
