package path

import (
	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/util"
	"math"
)

// BellmanFord 算法实现求解带负权边的最短路径问题
type BellmanFord struct {
	// 图
	graph.WeightedGraph
	// S 起点
	S int
	// 起点 S 至各顶点的距离
	Dis []int
	// 从起点到顶点的最短路径上的前一个顶点
	Pre []int
	// 不能有负权环
	HasNegCycle bool
}

func NewBellmanFord(g graph.WeightedGraph, s int) *BellmanFord {
	g.ValidateVertex(s)

	dis := make([]int, g.V())

	// 初始化：令各顶点的距离为正无穷，这里用 MaxInt32 表示正无穷
	for v := 0; v < g.V(); v++ {
		dis[v] = math.MaxInt32
	}
	// 起点的距离到自己为 0
	dis[s] = 0

	pre := make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		pre[v] = -1
	}

	var hasNegCycle bool

	// 最多做 V-1 轮松弛（Relaxation）操作，但可能在没有做完 V-1 次松弛操作就已经求得所有点的最短路
	// 对所有的边做一次松驰操作，则求出了所有的点，经过的边数最多为 1 的最短路
	// 对所有的边再做一次松驰操作，则求出了所有的点，经过的边数最多为 2 的最短路
	// 对所有的边做 V-1 次松驰操作，则求出了所有的点，经过的边数最多为 V-1 的最短路
	for pass := 1; pass < g.V(); pass++ {
		for v := 0; v < g.V(); v++ {
			for _, w := range g.Adj(v) {
				if dis[v] != math.MaxInt32 && dis[v]+g.GetWeight(v, w) < dis[w] {
					dis[w] = dis[v] + g.GetWeight(v, w)
					pre[w] = v
				}
			}
		}
	}

	// 再做一轮弛操作，如果某个点的最短路还能变得更小，则说明存在负权环
	// 对存在负权环的图求解最短路是没有意义的
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			if dis[v] != math.MaxInt32 && dis[v]+g.GetWeight(v, w) < dis[w] {
				hasNegCycle = true
			}
		}
	}

	return &BellmanFord{
		WeightedGraph: g,
		S:             s,
		Dis:           dis,
		Pre:           pre,
		HasNegCycle:   hasNegCycle,
	}
}

func (d *BellmanFord) IsConnectedTo(v int) bool {
	d.ValidateVertex(v)
	return d.Dis[v] != math.MaxInt32
}

func (d *BellmanFord) DistTo(v int) int {
	d.ValidateVertex(v)
	return d.Dis[v]
}

func (d *BellmanFord) Path(t int) []int {
	if !d.IsConnectedTo(t) {
		return nil
	}

	var paths []int
	for cur := t; cur != d.S; cur = d.Pre[cur] {
		paths = append(paths, cur)
	}
	paths = append(paths, d.S)

	util.ReverseSlice(paths)
	return paths
}

func (d *BellmanFord) HasNegativeCycle() bool {
	return d.HasNegCycle
}
