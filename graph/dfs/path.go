package dfs

import (
	"io.github.sowen.datastructure/graph"
)

type SingleSourcePath struct {
	graph.Graph
	// s the source vertex
	s int
	// visited marked whether the vertex had been progressed
	visited []bool
	// pre represents the previous vertex in path
	pre []int
}

func NewSingleSourcePath(g graph.Graph, s int) *SingleSourcePath {
	p := &SingleSourcePath{
		Graph:   g,
		s:       s,
		visited: make([]bool, g.V()),
		pre:     make([]int, g.V()),
	}
	p.dfs(p.s, p.s)
	return p
}

func (p *SingleSourcePath) dfs(v int, parent int) {
	p.visited[v] = true
	p.pre[v] = parent
	for _, w := range p.Graph.Adj(v) {
		if !p.visited[w] {
			p.dfs(w, v)
		}
	}
}

func (p *SingleSourcePath) IsConnectedTo(w int) bool {
	p.ValidateVertex(w)
	return p.visited[w]
}

func (p *SingleSourcePath) Path(w int) []int {
	p.ValidateVertex(w)
	var ret []int

	if !p.IsConnectedTo(w) {
		return ret
	}

	// for x := w; p.pre[x] >= 0 && x != p.s; x = p.pre[x] {
	//	ret = append(ret, x)
	//}

	for x := w; x != p.s; x = p.pre[x] {
		ret = append(ret, x)
	}
	ret = append(ret, p.s)

	// reverse slice
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
	}
	return ret
}
