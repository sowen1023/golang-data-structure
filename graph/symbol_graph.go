package graph

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"io.github.sowen.datastructure/tree"
)

type SymbolGraph struct {
	Graph                // Graph 图
	st    map[string]int // st 符号名 -> 索引
	keys  []string       // keys 索引 -> 符号名
}

func (g *SymbolGraph) G() Graph {
	return g.Graph
}

func (g *SymbolGraph) Contains(s string) bool {
	_, ok := g.st[s]
	return ok
}

func (g *SymbolGraph) Name(v int) string {
	return g.keys[v]
}

func (g *SymbolGraph) Index(s string) int {
	i, ok := g.st[s]
	if ok {
		return i
	}
	panic(fmt.Sprintf("Vertex of name %v not found", s))
}

func ReadSymbolGraph(src io.Reader, delim string) *SymbolGraph {
	if src == nil {
		panic("invalid io.Reader")
	}

	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanLines)

	// 将数据暂存，后面需要循环两次
	var data [][]string
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), delim)
		if len(tokens) > 0 {
			data = append(data, tokens)
		}
	}

	// 构建名称 -> 顶点的映射
	st := map[string]int{}
	for _, tokens := range data {
		for _, name := range tokens {
			if _, ok := st[name]; !ok {
				st[name] = len(st)
			}
		}
	}

	// 构建顶点 -> 名称映射
	keys := make([]string, len(st))
	for name, v := range st {
		keys[v] = name
	}

	// 构建图
	g := &AdjMap{
		v:   len(st),
		adj: tree.NewTreeMaps(len(st)),
	}

	// 向图中添加边
	for _, tokens := range data {
		v := st[tokens[0]]
		for i := 1; i < len(tokens); i++ {
			g.AddEdge(v, st[tokens[i]])
		}
	}

	return &SymbolGraph{
		Graph: g,
		st:    st,
		keys:  keys,
	}
}

func FileAsSymbolGraph(file string, delim string) *SymbolGraph {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Print(err)
		}
	}()
	return ReadSymbolGraph(f, delim)
}

func URLAsSymbolGraph(rawURL string, delim string) *SymbolGraph {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	if u.Scheme == "file" {
		return FileAsSymbolGraph(u.Path, delim)
	}
	if u.Scheme == "http" || u.Scheme == "https" {
		return HTTPUrlAsSymbolGraph(rawURL, delim, 10*time.Second)
	}
	panic("Unsupported URL scheme: " + u.Scheme)
}

func HTTPUrlAsSymbolGraph(rawURL string, delim string, timeout time.Duration) *SymbolGraph {
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest(http.MethodGet, rawURL, nil)
	if err != nil {
		panic(err)
	}
	response, err := client.Do(request)
	if response != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Print(err)
			}
		}(response.Body)
	}
	if err != nil {
		panic(err)
	}

	contentType := response.Header["Content-Type"]
	if len(contentType) > 0 && strings.HasSuffix(contentType[0], "text/plain") {
		panic("Unsupported Content-Type: " + contentType[0])
	}

	return ReadSymbolGraph(response.Body, delim)
}
