package tree

import (
	"fmt"

	"io.github.sowen.datastructure/types"
)

type Merger[E types.Comparable] func(l, r E) E

// example 1
//      [1, 6]
//      /    \
//   [1, 3]   [4, 6]
//    /  \     /   \
//  [1,2] [3] [4,5] [6]
//	/  \       /  \
// [1] [2]    [4] [5]

// example 2
//      [0, 7]
//      /     \
//   [0, 3]    [4, 7]
//    / \       /   \
//[0,1] [2,3] [4,5] [6,7]
//	/\    /\    /\    /\
//[0][1][2][3][4][5][6][7]

// SegmentTree O(longN)
type SegmentTree[E types.Comparable] struct {
	tree   []E
	data   []E
	merger func(E, E) E
}

func (st *SegmentTree[E]) String() string {
	return fmt.Sprint(st.tree)
}

func NewSegmentTree[E types.Comparable](data []E, merger func(E, E) E) *SegmentTree[E] {
	st := SegmentTree[E]{
		data:   make([]E, len(data)),
		tree:   make([]E, 4*len(data)),
		merger: merger,
	}
	copy(st.data, data)
	st.buildSegmentTree(0, 0, len(st.data)-1)
	return &st
}

func (st *SegmentTree[E]) buildSegmentTree(treeIndex int, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}

	leftTreeIndex, rightTreeIndex := st.leftChild(treeIndex), st.rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)
	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}

func (st *SegmentTree[E]) Get(index int) E {
	if index < 0 || index >= len(st.data) {
		panic(fmt.Sprintf("index %v is illegal", index))
	}
	return st.data[index]
}

// 给定指定节点下标，计算其左子树节点下标
func (st *SegmentTree[E]) leftChild(index int) int {
	return index*2 + 1
}

// 给定指定节点下标，计算其右子树节点下标
func (st *SegmentTree[E]) rightChild(index int) int {
	return index*2 + 2
}

// 数组下标以 0 开始作为根结点的二叉树左、右、父节点下标计算方式
//
// 给定指定节点下标，计算其父节点下标
// func (st *SegmentTree[E]) parent(index int) int {
// 	return (index - 1) / 2
// }

// Query 查询区间范围 [queryL, queryR] 的值
func (st *SegmentTree[E]) Query(queryL, queryR int) E {
	if queryL > queryR || queryL < 0 || queryR >= len(st.data) {
		panic(fmt.Sprintf("illegal query0 range [%v, %v]", queryL, queryR))
	}
	return st.query(0, 0, len(st.data)-1, queryL, queryR)
}

// 在以 treeIndex 为根节点的线段树中 [l,r] 的范围里，搜索区间 [queryL, queryR] 的值
func (st *SegmentTree[E]) query(treeIndex, l, r, queryL, queryR int) E {
	// 区间 [queryL, queryR] 正好相等
	if l == queryL && r == queryR {
		return st.tree[treeIndex]
	}

	m := l + (r-l)/2

	leftChildIndex := st.leftChild(treeIndex)
	rightChildIndex := st.rightChild(treeIndex)

	// 区间 [queryL, queryR] 全位于左子树上
	if queryR <= m {
		return st.query(leftChildIndex, l, m, queryL, queryR)
	}

	// 区间 [queryL, queryR] 全位于右子树上
	if queryL > m {
		return st.query(rightChildIndex, m+1, r, queryL, queryR)
	}

	// 区间 [queryL, queryR] 部分们于左子树，部分位于右子树

	leftRet := st.query(leftChildIndex, l, m, queryL, m)
	rightRet := st.query(rightChildIndex, m+1, r, m+1, queryR)

	return st.merger(leftRet, rightRet)
}

func (st *SegmentTree[E]) Set(index int, e E) {
	if index < 0 || index >= len(st.data) {
		panic(fmt.Sprintf("index %v is illegal", index))
	}
	st.set(0, 0, len(st.data)-1, index, e)
}

// 在以 treeIndex 为根的线段树中更新 index 的值为 e
func (st *SegmentTree[E]) set(treeIndex, l, r, index int, e E) {
	if l == r {
		st.tree[treeIndex] = e
		return
	}

	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	m := l + (r - l) + 1
	if index <= m {
		st.set(leftTreeIndex, l, m, index, e)
	} else {
		st.set(rightTreeIndex, m+1, r, index, e)
	}

	st.tree[treeIndex] = st.merger(st.tree[leftTreeIndex], st.tree[rightTreeIndex])
}
