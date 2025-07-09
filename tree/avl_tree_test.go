package tree

import (
	"fmt"
	"sort"
	"testing"
	"time"

	"io.github.sowen.datastructure/util"
)

func TestTreesPerformance(t *testing.T) {
	words := util.GetPrideAndPrejudiceWords()

	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	bst := &BST[string, int]{}
	for _, word := range words {
		bst.Add(word, 1)
	}

	start := time.Now()
	for _, word := range words {
		bst.Contains(word)
	}
	fmt.Printf("[BST] cost time:%v\n", time.Since(start))

	avl := &AVLTree[string, int]{}
	for _, word := range words {
		avl.Add(word, 1)
	}
	fmt.Printf("[AVL] is valid BST: %v\n", avl.IsValidBST())
	fmt.Printf("[AVL] is balanced: %v\n", avl.IsBalanced())

	start = time.Now()
	for _, word := range words {
		avl.Contains(word)
	}
	fmt.Printf("[AVL] cost time:%v\n", time.Since(start))

	m := make(map[string]int)
	for _, word := range words {
		m[word] = m[word] + 1
	}

	start = time.Now()
	for _, word := range words {
		_ = m[word]
	}
	fmt.Printf("[MAP] cost time:%v\n", time.Since(start))
}
