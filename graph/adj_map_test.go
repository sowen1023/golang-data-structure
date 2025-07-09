package graph

import (
	"fmt"
	"testing"

	"io.github.sowen.datastructure/util"
)

func TestAdjSet(t *testing.T) {
	graph := FileAsGraph(util.GetFileAbsolutePath("/data/g.txt"))
	fmt.Println(graph)
}
