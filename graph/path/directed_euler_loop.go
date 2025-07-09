package path

import (
	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/graph/dfs"
	"io.github.sowen.datastructure/stack"
)

// EulerLoop - 欧拉回路
type DirectedEulerLoop struct {
	graph.Graph
}

func NewDirectedEulerLoop(g graph.Graph) *DirectedEulerLoop {
	if !g.IsDirected() {
		panic("DirectedEulerLoop only works on directed graph")
	}
	return &DirectedEulerLoop{
		Graph: g,
	}
}

// HasEulerLoop - 检查图是否存在欧拉回路
func (e *DirectedEulerLoop) HasEulerLoop() bool {
	cc := dfs.NewCC(e)
	if cc.Count() > 1 {
		return false
	}
	for v := 0; v < e.V(); v++ {
		if e.Indegree(v) != e.Outdegree(v) {
			return false
		}
	}
	return true
}

// Result - Hierholzer 算法求解欧拉回路
func (e *DirectedEulerLoop) Result() (loop []int) {
	if !e.HasEulerLoop() {
		return nil
	}

	g := e.Graph.Clone()
	s := stack.New[int]()
	var v int
	s.Push(v)
	for !s.IsEmpty() {
		if g.Outdegree(v) != 0 {
			s.Push(v)
			w := g.Adj(v)[0]
			g.RemoveEdge(v, w)
			v = w
		} else {
			loop = append(loop, v)
			v = s.Pop()
		}
	}
	return loop
}
