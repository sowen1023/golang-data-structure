package mst

import (
	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/graph/dfs"
	"io.github.sowen.datastructure/tree"
	"sort"
)

// Kruskal 算法实现最小生成树 - Minimum Tree Spanning
func Kruskal(g graph.WeightedGraph) ([]graph.WeightedEdge, bool) {
	// 图中所有顶点必须是连通的
	cc := dfs.NewCC(g)
	if cc.Count() > 1 {
		return nil, false
	}

	var edges []*graph.WeightedEdge
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if v < w {
				edges = append(edges, &graph.WeightedEdge{V: v, W: w, Weight: g.GetWeight(v, w)})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return graph.LessWeightedEdge(edges[i], edges[j])
	})

	var mst []graph.WeightedEdge
	// 并查集
	uf := tree.NewUnionFind5(g.V())
	for _, edge := range edges {
		if !uf.IsConnected(edge.V, edge.W) {
			mst = append(mst, *edge)
			uf.UnionElements(edge.V, edge.W)
		}
	}
	return mst, true
}
