package path

import (
	"fmt"
	"io.github.sowen.datastructure/graph"
	"testing"
)

func TestBellmanFord(t *testing.T) {
	graphTexts := []string{DijkstraGraphText}
	for _, text := range graphTexts {
		g := graph.TextAsWeightedGraph(text)
		bf := NewBellmanFord(g, 0)
		for v := 0; v < g.V(); v++ {
			fmt.Printf("Distance %d -> %d is %d, ", bf.S, v, bf.DistTo(v))
			fmt.Printf("paths are %d\n", bf.Path(v))
		}
		fmt.Println()
	}
}
