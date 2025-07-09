package sort

import (
	"testing"

	"io.github.sowen.datastructure/graph"
)

func TestTopoSort(t *testing.T) {
	texts := []string{
		`5 5
			0 1
			1 2
			1 3
			2 4
			3 2`}
	for _, text := range texts {
		g := graph.ParseGraph(text, false, true)
		res, hasCycle := TopoSort(g)
		t.Log("Graph has a cycle?", hasCycle)
		if !hasCycle {
			t.Log("Graph's topological sort result: ", res)
		}
	}
}
