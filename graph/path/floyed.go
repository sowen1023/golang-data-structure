package path

import (
	"io.github.sowen.datastructure/graph"
	"math"
)

// Floyed 算法求解图任意两点的最短距离
type Floyed struct {
	// 图
	graph.WeightedGraph
	// 起点 S 至各顶点的距离
	Dis [][]int
	// 不能有负权环
	hasNegCycle bool
}

func NewFloyed(g graph.WeightedGraph) *Floyed {
	// 初始化：令各顶点的距离为正无穷，这里用 MaxInt32 表示正无穷
	dis := make([][]int, g.V())
	for v := 0; v < g.V(); v++ {
		dis[v] = make([]int, g.V())
		for w := 0; w < g.V(); w++ {
			dis[v][w] = math.MaxInt32
		}
	}

	for v := 0; v < g.V(); v++ {
		dis[v][v] = 0
		for _, w := range g.Adj(v) {
			dis[v][w] = g.GetWeight(v, w)
		}
	}

	for t := 0; t < g.V(); t++ {
		for v := 0; v < g.V(); v++ {
			for w := 0; w < g.V(); w++ {
				if dis[v][t] != math.MaxInt32 && dis[t][w] != math.MaxInt32 &&
					dis[v][t]+dis[t][w] < dis[v][w] {
					dis[v][w] = dis[v][t] + dis[t][w]
				}
			}
		}
	}

	var hasNegCycle bool
	for v := 0; v < g.V(); v++ {
		if dis[v][v] < 0 {
			hasNegCycle = true
			break
		}
	}

	return &Floyed{
		WeightedGraph: g,
		Dis:           dis,
		hasNegCycle:   hasNegCycle,
	}
}

func (f *Floyed) IsConnectedTo(v int, w int) bool {
	f.ValidateVertex(v)
	f.ValidateVertex(w)
	return f.Dis[v][w] != math.MaxInt32
}

func (f *Floyed) DistTo(v int, w int) int {
	f.ValidateVertex(v)
	f.ValidateVertex(w)
	return f.Dis[v][w]
}

func (f *Floyed) HasNegativeCycle() bool {
	return f.hasNegCycle
}
