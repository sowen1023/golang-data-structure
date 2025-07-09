package path

import (
	"io.github.sowen.datastructure/graph"
	"log"
	"testing"
)

const NegCycleGraphText = `5 8
							0 1 -1
							0 2 2
							1 2 1
							1 3 2
							1 4 3
							2 3 4
							2 4 5
							3 4 1`

func TestFloyed(t *testing.T) {
	texts := []string{DijkstraGraphText}
	for _, text := range texts {
		g := graph.TextAsWeightedGraph(text)
		f := NewFloyed(g)
		if f.HasNegativeCycle() {
			log.Fatal("Negative cycle exist")
		}
		for v := 0; v < g.V(); v++ {
			for w := 0; w < g.V(); w++ {
				t.Logf("Distance %d -> %d is %d\n", v, w, f.DistTo(v, w))
			}
		}
	}
}

func TestFloyed_HasNegativeCycle(t *testing.T) {
	g := graph.TextAsWeightedGraph(NegCycleGraphText)
	f := NewFloyed(g)
	if !f.HasNegativeCycle() {
		t.Fatal("Negative cycle not identified")
	}
}
