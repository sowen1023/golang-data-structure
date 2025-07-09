package mst

import (
	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/graph/dfs"
	"io.github.sowen.datastructure/util"
)

// Prime 算法实现最小生成树 - Minimum Tree Spanning
func Prime(g graph.WeightedGraph) ([]*graph.WeightedEdge, bool) {
	// 图中所有顶点必须是连通的
	cc := dfs.NewCC(g)
	if cc.Count() > 1 {
		return nil, false
	}

	var mst []*graph.WeightedEdge
	visited := make([]bool, g.V())
	visited[0] = true

	pq := util.NewPriorityQueue[*graph.WeightedEdge](graph.LessWeightedEdge)
	for _, w := range g.Adj(0) {
		pq.Push(&graph.WeightedEdge{
			V:      0,
			W:      w,
			Weight: g.GetWeight(0, w),
		})
	}
	for !pq.IsEmpty() {
		edge := pq.Pop()
		// 两个顶点 w & w 属于同一切分
		if visited[edge.W] && visited[edge.V] {
			continue
		}
		var nv = edge.W
		if visited[edge.W] {
			nv = edge.V
		}
		for _, w := range g.Adj(nv) {
			pq.Push(&graph.WeightedEdge{
				V: nv, W: w,
				Weight: g.GetWeight(nv, w),
			})
		}

		visited[edge.W] = true
		visited[edge.V] = true
		mst = append(mst, edge)
	}
	return mst, true
}
