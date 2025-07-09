package path

import (
	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/graph/dfs"
	"io.github.sowen.datastructure/stack"
)

// EulerLoop - 欧拉回路
type EulerLoop struct {
	graph.Graph
}

func NewEulerLoop(g graph.Graph) *EulerLoop {
	if g.IsDirected() {
		panic("EulerLoop only works on undirected graph")
	}
	return &EulerLoop{
		Graph: g,
	}
}

// HasEulerLoop - 检查图是否存在欧拉回路
func (e *EulerLoop) HasEulerLoop() bool {
	cc := dfs.NewCC(e)
	if cc.Count() > 1 {
		return false
	}
	for v := 0; v < e.V(); v++ {
		if len(e.Adj(v))%2 != 0 {
			return false
		}
	}
	return true
}

// Result - Hierholzer 算法求解欧拉回路
func (e *EulerLoop) Result() (loop []int) {
	if !e.HasEulerLoop() {
		return nil
	}

	g := e.Graph.Clone()
	s := stack.New[int]()
	var v int
	s.Push(v)
	for !s.IsEmpty() {
		if g.Degree(v) != 0 {
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
