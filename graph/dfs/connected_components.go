package dfs

import "io.github.sowen.datastructure/graph"

// CC connected components count
type CC struct {
	graph.Graph
	visited []bool // visited marked whether the vertex had been progressed
	count   int    // count of the connected components
}

func NewCC(g graph.Graph) *CC {
	cc := &CC{
		Graph:   g,
		visited: make([]bool, g.V()),
	}

	for v := 0; v < g.V(); v++ {
		if !cc.visited[v] {
			cc.dfs(v)
			cc.count++
		}
	}
	return cc
}

// dfs Depth-First-Search
func (c *CC) dfs(v int) {
	c.visited[v] = true
	for _, w := range c.Graph.Adj(v) {
		if !c.visited[w] {
			c.dfs(w)
		}
	}
}

func (c *CC) Count() int {
	return c.count
}

// CCDistributed connected components count
type CCDistributed struct {
	graph.Graph
	visited []int // visited marked the vertex belongs to a connected components
	count   int   // count of the connected components
}

func NewCCDistributed(g graph.Graph) *CCDistributed {
	cc := &CCDistributed{
		Graph:   g,
		visited: make([]int, g.V()),
	}

	for v := 0; v < g.V(); v++ {
		if cc.visited[v] == 0 {
			cc.count++
			cc.dfs(v, cc.count)
		}
	}
	return cc
}

func (c *CCDistributed) dfs(v int, ccid int) {
	c.visited[v] = ccid
	for _, w := range c.Graph.Adj(v) {
		if c.visited[w] == 0 {
			c.dfs(w, ccid)
		}
	}
}

func (c *CCDistributed) Count() int {
	return c.count
}

func (c *CCDistributed) IsConnected(v, w int) bool {
	c.Graph.ValidateVertex(v)
	c.Graph.ValidateVertex(w)
	return c.visited[v] == c.visited[w]
}

func (c *CCDistributed) Components() [][]int {
	res := make([][]int, c.count)
	for v := 0; v < c.Graph.V(); v++ {
		res[c.visited[v]-1] = append(res[c.visited[v]-1], v)
	}
	return res
}
