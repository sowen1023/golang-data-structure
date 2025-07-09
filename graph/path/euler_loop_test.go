package path

import (
	"testing"

	"io.github.sowen.datastructure/graph"
)

func TestEulerLoop(t *testing.T) {
	texts := []string{
		`5 6
			0 1
			0 2
			1 2
			2 3
			2 4
			3 4`,
		`11 15
			0 1
			0 3
			1 2
			1 4
			1 5
			2 5
			3 4
			4 5
			4 6
			5 7
			6 7
			7 8
			7 9
			8 10
			9 10`}
	for _, text := range texts {
		g := graph.TextAsGraph(text)
		eulerLoop := NewEulerLoop(g)
		t.Log("Graph has euler path?", eulerLoop.HasEulerLoop())
		if eulerLoop.HasEulerLoop() {
			t.Log("Graph's euler path: ", eulerLoop.Result())
		}
	}
}
