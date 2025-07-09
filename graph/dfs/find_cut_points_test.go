package dfs

import (
	"testing"

	"io.github.sowen.datastructure/graph"
)

func TestFindCutPoints(t *testing.T) {
	graphTexts := []string{Graph1Text, Graph2Text, Graph3Text, Graph4Text}
	for _, text := range graphTexts {
		g := graph.TextAsGraph(text)
		t.Log("Graph's cut points?", NewFindCutPoints(g).Result())
	}
}

func TestFindCutPointsVerbose(t *testing.T) {
	graphTexts := []string{Graph3Text}
	for _, text := range graphTexts {
		g := graph.TextAsGraph(text)
		t.Log("Graph's cut points?", NewFindCutPointsWithLog(g, true).Result())
	}
}
