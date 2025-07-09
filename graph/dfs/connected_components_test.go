package dfs

import (
	"testing"

	"io.github.sowen.datastructure/graph"
	"io.github.sowen.datastructure/util"
)

const TestGraphPath = "/data/g.txt"

func TestCC(t *testing.T) {
	g := graph.FileAsGraph(util.GetFileAbsolutePath(TestGraphPath))
	cc := NewCC(g)
	t.Log("cc count:", cc.Count())
}

func TestCCDistributed(t *testing.T) {
	g := graph.FileAsGraph(util.GetFileAbsolutePath(TestGraphPath))
	cc := NewCCDistributed(g)
	t.Log("connected components count:", cc.Count())
	t.Logf("%v and %v is connected? %v", 1, 6, cc.IsConnected(1, 6))
	t.Logf("%v and %v is connected? %v", 2, 5, cc.IsConnected(2, 5))
	t.Log("connected components:", cc.Components())
}
