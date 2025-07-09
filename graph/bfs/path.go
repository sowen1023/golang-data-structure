package bfs

import (
	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/queue"
)

type BreadFirstPaths struct {
	graph.Graph
	visited []bool
	edgeTo  []int // 到达该顶点的已知路径上的最后一个顶点
	s       int   // 起点
}

func NewBreadFirstPaths(g graph.Graph, s int) *BreadFirstPaths {
	path := &BreadFirstPaths{
		Graph:   g,
		visited: make([]bool, g.V()),
		edgeTo:  make([]int, g.V()),
		s:       s,
	}
	path.bfs(s)
	return path
}

func (p *BreadFirstPaths) bfs(s int) {
	q := queue.NewLoopQueue[int]()
	p.edgeTo[s] = s
	p.visited[s] = true
	q.Enqueue(s)

	for !q.IsEmpty() {
		v := q.Dequeue()
		for _, w := range p.Graph.Adj(v) {
			if !p.visited[w] {
				p.edgeTo[w] = v
				p.visited[w] = true
				q.Enqueue(w)
			}
		}
	}
}

func (p *BreadFirstPaths) IsConnectedTo(w int) bool {
	p.ValidateVertex(w)
	return p.visited[w]
}

func (p *BreadFirstPaths) Path(w int) []int {
	p.ValidateVertex(w)
	var ret []int

	if !p.IsConnectedTo(w) {
		return ret
	}

	for x := w; x != p.s; x = p.edgeTo[x] {
		ret = append(ret, x)
	}
	ret = append(ret, p.s)

	// reverse slice
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}
	return ret
}
