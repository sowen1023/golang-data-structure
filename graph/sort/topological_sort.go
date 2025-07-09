package sort

import "io.github.sowen.datastructure/graph"

// TopoSort - 返回拓扑排序结果 & 是否有环
func TopoSort(g graph.Graph) ([]int, bool) {
	if !g.IsDirected() {
		panic("TopoSort only works on directed graph")
	}

	var res []int

	indegrees := make([]int, g.V())
	var q []int
	for v := 0; v < g.V(); v++ {
		indegrees[v] = g.Indegree(v)
		if indegrees[v] == 0 {
			q = append(q, v)
		}
	}

	for len(q) > 0 {
		cur := q[0]
		res = append(res, cur)
		q = q[1:]
		for _, next := range g.Adj(cur) {
			indegrees[next]--
			if indegrees[next] == 0 {
				q = append(q, next)
			}
		}
	}

	if len(res) != g.V() {
		return nil, true
	}

	return res, false
}
