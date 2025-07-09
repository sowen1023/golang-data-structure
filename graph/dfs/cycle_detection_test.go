package dfs

import (
	"testing"

	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/util"
)

func TestCycleDetection(t *testing.T) {
	g1 := graph.FileAsGraph(util.GetFileAbsolutePath(TestGraphPath))
	t.Log("Graph has cycle?", NewCycleDetection(g1).HasCycle())

	text := `7 5
			 0 1
			 0 2
			 1 3
			 1 4
			 2 6`
	g2 := graph.TextAsGraph(text)
	t.Log("Graph has cycle?", NewCycleDetection(g2).HasCycle())

	text = `4 4
			0 1
			0 3
			1 2
			2 3`
	g3 := graph.TextAsGraph(text)
	t.Log("Graph has cycle?", NewCycleDetection(g3).HasCycle())
}
