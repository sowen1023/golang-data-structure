package tree

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"io.github.sowen.datastructure/util"
)

func wordsData() []string {
	filePath := util.GetFileAbsolutePath("/data/pride-and-prejudice.txt")
	words := util.ReadFile(filePath)
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	fmt.Printf("[Pride-And-Prejudice] words count: %d\n", len(words))
	return words
}

func TestTreePerformance(t *testing.T) {
	words := wordsData()

	//sort.Slice(words, func(i, j int) bool {
	//	return words[i] < words[j]
	//})

	bst := &BST[string, int]{}
	start := time.Now()
	testBST(bst, words)
	fmt.Printf("[BST] cost time:%v\n", time.Since(start))
	fmt.Printf("Total different words:%v\n", bst.Size())
	fmt.Printf("Frequency of PRIDE: %v\n", bst.Get("pride"))
	fmt.Printf("Frequency of PREJUDICE: %v\n", bst.Get("prejudice"))

	avl := &AVLTree[string, int]{}
	start = time.Now()
	testAVLTree(avl, words)
	fmt.Printf("[AVL] cost time:%v\n", time.Since(start))
	fmt.Printf("Total different words:%v\n", avl.Size())
	fmt.Printf("Frequency of PRIDE: %v\n", avl.Get("pride"))
	fmt.Printf("Frequency of PREJUDICE: %v\n", avl.Get("prejudice"))

	if !avl.IsValidBST() {
		t.Fatal("[AVL] isn't valid BST")
	}
	if !avl.IsBalanced() {
		t.Fatal("[AVL] isn't balanced")
	}

	for _, word := range words {
		avl.Remove(word)
		if !avl.IsValidBST() {
			t.Fatal("[AVL] isn't valid BST")
		}
		if !avl.IsBalanced() {
			t.Fatal("[AVL] isn't balanced")
		}
	}

	m := make(map[string]int)
	start = time.Now()
	testMap(m, words)
	fmt.Printf("[MAP] cost time:%v\n", time.Since(start))
	fmt.Printf("Total different words:%v\n", len(m))
	fmt.Printf("Frequency of PRIDE: %v\n", m["pride"])
	fmt.Printf("Frequency of PREJUDICE: %v\n", m["prejudice"])
}

func testBST(bst *BST[string, int], words []string) {
	for _, word := range words {
		if bst.Contains(word) {
			bst.Add(word, bst.Get(word)+1)
		} else {
			bst.Add(word, 1)
		}
	}
}

func testAVLTree(avl *AVLTree[string, int], words []string) {
	for _, word := range words {
		if avl.Contains(word) {
			avl.Add(word, avl.Get(word)+1)
		} else {
			avl.Add(word, 1)
		}
	}
}

func testMap(m map[string]int, words []string) {
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word] = m[word] + 1
		} else {
			m[word] = 1
		}
	}
}

func TestTrie(t *testing.T) {
	words := wordsData()

	bst := &BST[string, int]{}

	start := time.Now()
	for _, word := range words {
		bst.Add(word, 1)
	}
	for _, word := range words {
		bst.Contains(word)
	}
	fmt.Printf("[BST] cost time:%v\n", time.Since(start))
	fmt.Printf("Total different words:%v\n", bst.Size())

	trie := NewTrie()

	start = time.Now()
	for _, word := range words {
		trie.Add(word)
	}
	for _, word := range words {
		trie.Contains(word)
	}
	fmt.Printf("[Trie] cost time:%v\n", time.Since(start))
	fmt.Printf("Total different words:%v\n", trie.Size())
}
