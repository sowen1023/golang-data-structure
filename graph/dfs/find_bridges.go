package dfs

import (
	"io.github.sowen.datastructure/graph"
)

type FindBridges struct {
	graph.Graph
	visited []bool
	count   int
	order   []int
	low     []int
	res     []graph.Edge
}

func NewFindBridges(g graph.Graph) *FindBridges {
	fb := &FindBridges{
		Graph:   g,
		visited: make([]bool, g.V()),
		order:   make([]int, g.V()),
		low:     make([]int, g.V()),
	}
	for v := 0; v < g.V(); v++ {
		if !fb.visited[v] {
			fb.dfs(v, v)
		}
	}
	return fb
}

// dfs Depth-First-Search
func (f *FindBridges) dfs(v int, parent int) {
	f.visited[v] = true
	f.order[v] = f.count
	f.low[v] = f.order[v]
	f.count++
	for _, w := range f.Graph.Adj(v) {
		if !f.visited[w] {
			f.dfs(w, v)
			f.low[v] = min(f.low[v], f.low[w])
			// w 找不到第二条路回到 v
			if f.low[w] > f.order[v] {
				f.res = append(f.res, graph.Edge{V: v, W: w})
			}
		} else if w != parent {
			// 回到了比访问 v 更早的顶点
			f.low[v] = min(f.low[v], f.low[w])
		}
	}
}

func min(v, w int) int {
	if v <= w {
		return v
	}
	return w
}

func (f *FindBridges) Result() []graph.Edge {
	return f.res
}
