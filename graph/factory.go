package graph

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"io.github.sowen.datastructure/tree"
)

func FileAsGraph(file string) Graph {
	return FileAsWeightedGraph(file, false, false)
}

func FileAsWeightedGraph(file string, weighted bool, directed bool) *AdjMap {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Print(err)
		}
	}()
	return buildAdjMap(f, weighted, directed)
}

func TextAsGraph(src string) Graph {
	return ParseGraph(src, false, false)
}

func TextAsWeightedGraph(src string) WeightedGraph {
	return ParseGraph(src, true, false)
}

func ParseGraph(src string, weighted bool, directed bool) *AdjMap {
	return buildAdjMap(bytes.NewReader([]byte(src)), weighted, directed)
}

func buildAdjMap(src io.Reader, weighted bool, directed bool) *AdjMap {
	if src == nil {
		panic("invalid io.Reader")
	}

	reg := regexp.MustCompile(`-?\d+`)
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanLines)

	var V, E int
	var adj []tree.TreeMap
	var indegrees []int
	var outdegrees []int

	if scanner.Scan() {
		tokens := reg.FindAllString(scanner.Text(), -1)
		if len(tokens) != 2 {
			panic(fmt.Sprintf("read graph V & E faild at first row; required 2 tokens but got %d count", len(tokens)))
		}
		V, _ = strconv.Atoi(tokens[0])
		E, _ = strconv.Atoi(tokens[1])
		adj = make([]tree.TreeMap, V)
		indegrees = make([]int, V)
		outdegrees = make([]int, V)
		for i := 0; i < len(adj); i++ {
			adj[i] = tree.NewTreeMap()
		}
	}
	for row := 1; scanner.Scan(); row++ {
		tokens := reg.FindAllString(scanner.Text(), -1)
		cnt := len(tokens)
		if weighted && cnt != 3 {
			panic(fmt.Sprintf("read weighted edge failed at row: %v; got %d tokens", row, cnt))
		} else if !weighted && cnt != 2 {
			panic(fmt.Sprintf("read unweighted edge failed at row: %v; got %d tokens", row, cnt))
		}

		v, _ := strconv.Atoi(tokens[0])
		validateVertex(v, V)
		w, _ := strconv.Atoi(tokens[1])
		validateVertex(w, V)
		if v == w {
			panic("Self path is detected!")
		}
		if adj[v].Contains(w) {
			panic("Parallel edges detected!")
		}
		weight := 1
		if weighted {
			weight, _ = strconv.Atoi(tokens[2])
		}
		adj[v].Put(w, weight)
		if directed {
			outdegrees[v]++
			indegrees[w]++
		} else {
			adj[w].Put(v, weight)
		}
	}
	return &AdjMap{
		v:          V,
		e:          E,
		adj:        adj,
		directed:   directed,
		indegrees:  indegrees,
		outdegrees: outdegrees,
	}
}
