package dfs

import (
	"testing"

	"io.github.sowen.datastructure/graph"
)

func TestBipartitionDetection(t *testing.T) {
	texts := []string{
		`7 6
			0 1
			0 2
			1 3
			1 4
			2 3
			2 6`,
		`4 6
			0 1
			0 2
			0 3
			1 2
			1 3
			2 3`,
		`4 4
			0 1
			0 3
			1 2
			2 3`,
	}
	for _, text := range texts {
		g := graph.TextAsGraph(text)
		t.Log("Graph is bipartite?", NewBipartitionDetection(g).IsBipartite())
	}
}
