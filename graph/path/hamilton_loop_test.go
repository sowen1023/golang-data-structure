package path

import (
	"io.github.sowen.datastructure/graph"
	"testing"
)

const HamiltonG1Text = `4 5
					0 1
					0 2
					0 3
					1 2
					1 3`
const HamiltonG2Text = `20 30
					0 1
					0 4
					0 13
					1 2
					1 11
					2 9
					2 3
					3 4
					3 7
					4 5
					5 6
					5 14
					6 7
					6 16
					7 8
					8 9
					8 17
					9 10
					10 11
					10 18
					11 12
					12 13
					12 19
					13 14
					14 15
					15 16
					15 19
					16 17
					17 18
					18 19`

func TestHamiltonLoop(t *testing.T) {
	graphTexts := []string{HamiltonG1Text, HamiltonG2Text}
	for _, text := range graphTexts {
		g := graph.TextAsGraph(text)
		hl := NewHamiltonLoop(g)
		t.Log("Graph has hamilton path?", hl.HasHamiltonLoop())
		if hl.HasHamiltonLoop() {
			t.Log("Graph's hamilton path: ", hl.Result())
		}
	}
}
