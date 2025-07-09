package dfs

import "io.github.sowen.datastructure/graph"

func DirectedCycleDetection(g graph.Graph) bool {
	if !g.IsDirected() {
		panic("CycleDetection only works on directed graph")
	}

	visited := make([]bool, g.V())
	onPath := make([]bool, g.V())

	// 定义 dfs 函数
	var dfs func(int, int) bool
	dfs = func(v int, parent int) bool {
		visited[v] = true
		onPath[v] = true
		for _, w := range g.Adj(v) {
			if !visited[w] {
				if dfs(w, v) {
					return true
				}
			} else if onPath[w] { // can't check w != parent
				return true
			}
		}
		// 回溯
		onPath[v] = false
		return false
	}

	for v := 0; v < g.V(); v++ {
		if !visited[v] && dfs(v, v) {
			return true
		}
	}
	return false
}
