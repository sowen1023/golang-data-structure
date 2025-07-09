package dfs

import (
	"testing"

	"io.github.sowen.datastructure/graph"
)

func TestDirctedCycleDetection(t *testing.T) {
	texts := []string{
		`5 5
			0 1
			1 2
			1 3
			2 4
			3 2`,
		`5 6
			0 1
			1 2
			1 3
			2 4
			3 2
			3 0`,
	}
	for _, text := range texts {
		// directed graph
		g := graph.ParseGraph(text, false, true)
		t.Logf("Graph has cycle? %v", DirectedCycleDetection(g))
	}
}
