package mst

import (
	"fmt"
	"io.github.sowen.datastructure/graph"
	"testing"
)

const MstGText = `7 12
				0 1 2
				0 3 7
				0 5 2
				1 2 1
				1 3 4
				1 4 3
				1 5 5
				2 4 4
				2 5 4
				3 4 1
				3 6 5
				4 6 7`

func TestKruskal(t *testing.T) {
	g := graph.TextAsWeightedGraph(MstGText)
	fmt.Println(g)
	edges, ok := Kruskal(g)
	if !ok {
		t.Fatalf("Minimum tree spanning failed")
	}
	fmt.Println("Minimum Tree Spanning by Kruskal Algorithm:", edges)
}

func TestPrime(t *testing.T) {
	g := graph.TextAsWeightedGraph(MstGText)
	fmt.Println(g)
	edges, ok := Prime(g)
	if !ok {
		t.Fatalf("Minimum tree spanning failed")
	}
	fmt.Println("Minimum Tree Spanning by Prime Algorithm:", edges)
}
