package dfs

import (
	"fmt"
	"strings"

	"io.github.sowen.datastructure/tree"

	"io.github.sowen.datastructure/graph"
)

type FindCutPoints struct {
	graph.Graph
	visited []bool
	count   int
	order   []int
	low     []int
	res     tree.TreeSet
}

func NewFindCutPoints(g graph.Graph) *FindCutPoints {
	return NewFindCutPointsWithLog(g, false)
}

func NewFindCutPointsWithLog(g graph.Graph, verbose bool) *FindCutPoints {
	f := &FindCutPoints{
		Graph:   g,
		visited: make([]bool, g.V()),
		order:   make([]int, g.V()),
		low:     make([]int, g.V()),
		res:     tree.NewTreeSet(),
	}
	for v := 0; v < g.V(); v++ {
		if !f.visited[v] {
			if verbose {
				f.dfsWithLog(v, v, 0)
			} else {
				f.dfs(v, v)
			}
		}
	}
	return f
}

func (f *FindCutPoints) dfs(v int, parent int) {
	f.visited[v] = true
	f.order[v] = f.count
	f.low[v] = f.order[v]
	f.count++
	var child int
	for _, w := range f.Graph.Adj(v) {
		if !f.visited[w] {
			f.dfs(w, v)
			f.low[v] = min(f.low[v], f.low[w])
			// vertex v is a cut point because its child can't reach its parents
			if v != parent && f.low[w] >= f.order[v] {
				f.res.Add(v)
			}
			child++
			// the start vertex is a cut point that has more than 1 children
			if v == parent && child > 1 {
				f.res.Add(v)
			}
		} else if w != parent {
			// pay attention: compare with w's order[w] not low[w]
			f.low[v] = min(f.low[v], f.order[w])
		}
	}
}

func (f *FindCutPoints) Result() []int {
	if f.res.Size() == 0 {
		return nil
	}
	return f.res.Keys()
}

// dfsWithLog print some log
func (f *FindCutPoints) dfsWithLog(v int, parent int, level int) {
	logPrefix := getLogPrefix(level)
	fmt.Printf("%v=> [p:%d -> v:%d]\n", logPrefix, parent, v)

	f.visited[v] = true
	f.order[v] = f.count
	f.low[v] = f.order[v]
	f.count++
	var child int
	for _, w := range f.Graph.Adj(v) {
		if !f.visited[w] {
			fmt.Printf("%v=> [p:%d -> v:%d -> w:%d]\n", logPrefix, parent, v, w)
			f.dfsWithLog(w, v, level+1)
			fmt.Printf("%v<= [p:%d -> v:%d -> w:%d]\n", logPrefix, parent, v, w)

			f.low[v] = min(f.low[v], f.low[w])
			if v != parent && f.low[w] >= f.order[v] {
				f.res.Add(v)
				fmt.Printf("%v √ Find cut point:[%d]\n", logPrefix, v)
			}
			child++
			if v == parent && child > 1 {
				f.res.Add(v)
				fmt.Printf("%v √ Find cut point:[%d]\n", logPrefix, v)
			}
		} else if w != parent {
			f.low[v] = min(f.low[v], f.order[w])
		}
		fmt.Printf("%v low[%d] = %d\n", logPrefix, v, f.low[v])
	}
	fmt.Printf("%v<= [p:%d -> v:%d]\n", logPrefix, parent, v)
}

func getLogPrefix(level int) string {
	var res strings.Builder
	for i := 0; i < level; i++ {
		res.WriteString("  ")
	}
	return res.String()
}
