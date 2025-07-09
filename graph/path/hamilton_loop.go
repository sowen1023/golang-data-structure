package path

import (
	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/graph/dfs"
)

// HamiltonLoop - 哈密尔顿回路
type HamiltonLoop struct {
	graph.Graph
	start   int
	visited map[int]bool
	pre     []int
}

func NewHamiltonLoop(g graph.Graph) *HamiltonLoop {
	pre := make([]int, g.V())
	for v := 0; v < g.V(); v++ {
		pre[v] = -1
	}
	hl := &HamiltonLoop{
		Graph:   g,
		start:   0,
		visited: make(map[int]bool, g.V()),
		pre:     pre,
	}
	cc := dfs.NewCC(g)
	if cc.Count() > 1 {
		return hl
	}
	// hl.visited[hl.start] = true
	hl.dfs(hl.start, hl.start)
	return hl
}

func (h *HamiltonLoop) dfs(v, parent int) bool {
	h.visited[v] = true
	h.pre[v] = parent
	for _, w := range h.Adj(v) {
		if !h.visited[w] {
			if h.dfs(w, v) {
				return true
			}
		} else if w == h.start && len(h.visited) == h.V() {
			h.pre[h.start] = v
			return true
		}
	}
	h.pre[v] = -1
	delete(h.visited, v)
	return false
}

// func (hl *HamiltonLoop) dfs(v int) bool {
//	for _, w := range hl.Adj(v) {
//		if !hl.visited[w] {
//			hl.visited[w] = true
//			hl.pre[w] = v
//			if hl.resolveHamiltonLoop(w) {
//				return true
//			}
//			hl.pre[w] = -1
//			delete(hl.visited, w)
//		} else if w == hl.start && len(hl.visited) == hl.V() {
//			hl.pre[hl.start] = v
//			return true
//		}
//	}
//	return false
//}

// HasHamiltonLoop - 检查图是否存在哈密尔顿回路
func (h *HamiltonLoop) HasHamiltonLoop() bool {
	return h.pre[h.start] != -1
}

func (h *HamiltonLoop) Result() []int {
	if !h.HasHamiltonLoop() {
		return nil
	}
	var res []int
	for cur := h.pre[h.start]; cur != h.start; cur = h.pre[cur] {
		res = append(res, cur)
	}
	res = append(res, h.start)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
