package path

import (
	"io.github.sowen.datastructure/graph"
	"testing"
)

const DijkstraGraphText = `5 8
					0 1 4
					0 2 2
					1 2 1
					1 3 2
					1 4 3
					2 3 4
					2 4 5
					3 4 1`

func TestDijkstra(t *testing.T) {
	graphTexts := []string{DijkstraGraphText}
	for _, text := range graphTexts {
		g := graph.TextAsWeightedGraph(text)
		dij := NewDijkstra(g, 0)
		for v := 0; v < g.V(); v++ {
			t.Logf("Distance %d -> %d is %d, paths: %v ", dij.S, v, dij.DistTo(v), dij.Path(v))
		}
	}
}
