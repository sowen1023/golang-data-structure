package path

import (
	"testing"

	"io.github.sowen.datastructure/graph"
)

func TestDirectedEulerLoop(t *testing.T) {
	texts := []string{
		`5 5
			0 1
			1 2
			1 3
			2 4
			3 2`,
		`5 8
			0 1
			1 2
			1 3
			2 0
			2 4
			3 1
			3 2
			4 3`}
	for _, text := range texts {
		g := graph.ParseGraph(text, false, true)
		eulerLoop := NewDirectedEulerLoop(g)
		t.Log("Graph has euler path?", eulerLoop.HasEulerLoop())
		if eulerLoop.HasEulerLoop() {
			t.Log("Graph's euler path: ", eulerLoop.Result())
		}
	}
}
