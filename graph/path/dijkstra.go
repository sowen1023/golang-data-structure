package path

import (
	"math"

	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/util"
)

// Dijkstra 算法实现求解有权图单源最小路径
type Dijkstra struct {
	// 图
	graph.WeightedGraph
	// S 起点
	S int
	// 起点 S 至各顶点的距离
	Dis []int
	// 标记顶点是否已经求得最短距离
	Visited []bool
	// 从起点到顶点的最短路径上的前一个顶点
	Pre []int
}

func NewDijkstra(g graph.WeightedGraph, s int) *Dijkstra {
	g.ValidateVertex(s)

	dis := make([]int, g.V())

	// 初始化：令各顶点的距离为正无穷，这里用 MaxInt32 表示正无穷
	for v := 0; v < g.V(); v++ {
		dis[v] = math.MaxInt32
	}
	// 起点的距离到自己为 0
	dis[s] = 0
	visited := make([]bool, g.V())
	pre := make([]int, g.V())

	// 使用最小堆存储各顶点与起点间的距离
	pq := util.NewPriorityQueue(graph.LessNode)
	pq.Push(&graph.Node{V: s, Dis: dis[s]})

	for !pq.IsEmpty() {
		cur := pq.Pop().V

		// 跳过已经求得出短距离的顶点
		if visited[cur] {
			continue
		}

		// 堆中取的第一个顶点，当前计算出的该顶点与起点间的距离即为起点到它的最短距离；
		// 因为如果从其他顶点回到该顶点求得的距离肯定大于当前计算结果，不可能得到更小的值了
		visited[cur] = true

		for _, w := range g.Adj(cur) {
			if dis[cur]+g.GetWeight(cur, w) < dis[w] {
				dis[w] = dis[cur] + g.GetWeight(cur, w)
				pq.Push(&graph.Node{V: w, Dis: dis[w]})
				pre[w] = cur
			}
		}
	}

	return &Dijkstra{
		WeightedGraph: g,
		S:             s,
		Dis:           dis,
		Visited:       visited,
		Pre:           pre,
	}
}

func (d *Dijkstra) IsConnectedTo(v int) bool {
	d.ValidateVertex(v)
	return d.Visited[v]
}

func (d *Dijkstra) DistTo(v int) int {
	d.ValidateVertex(v)
	return d.Dis[v]
}

func (d *Dijkstra) Path(t int) []int {
	if !d.IsConnectedTo(t) {
		return nil
	}

	var paths []int
	for cur := t; cur != d.S; cur = d.Pre[cur] {
		paths = append(paths, cur)
	}
	paths = append(paths, d.S)

	// reverse slice
	for i := 0; i < len(paths)/2; i++ {
		paths[i], paths[len(paths)-i-1] = paths[len(paths)-i-1], paths[i]
	}
	return paths
}
