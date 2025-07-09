package tree

import (
	"fmt"
	"io.github.sowen.datastructure/queue"
	"io.github.sowen.datastructure/stack"
	"math"
)

type LazySegmentTree struct {

	// tree 节点数据
	tree []int

	// lazy 坐标与值的关系，值不为 0 表示左右子结点待更新
	lazy []int

	// size 原始数据大小
	size int

	// merger 合并函数
	merger func(int, int) int
}

func (st *LazySegmentTree) String() string {
	return fmt.Sprint(st.tree)
}

// Java method
// Height of segment tree
// int x = (int) (Math.ceil(Math.log(n) / Math.log(2)));
// Maximum size of segment tree
// int max_size = 2 * (int) Math.pow(2, x) - 1;

// NewLazySegmentTree creates a lazy updatable segment tree
func NewLazySegmentTree(data []int) *LazySegmentTree {
	h := math.Ceil(math.Log2(float64(len(data))))
	sz := 2*int(math.Pow(2, h)) - 1

	st := LazySegmentTree{
		tree: make([]int, sz),
		lazy: make([]int, sz),
		size: len(data),
		merger: func(a int, b int) int {
			return a + b
		},
	}
	st.buildLazySegmentTree(0, data, 0, len(data)-1)
	return &st
}

func (st *LazySegmentTree) buildLazySegmentTree(treeIndex int, data []int, l, r int) {
	if l == r {
		st.tree[treeIndex] = data[l]
		return
	}

	leftTree, rightTree := st.leftChild(treeIndex), st.rightChild(treeIndex)

	mid := l + (r-l)/2
	st.buildLazySegmentTree(leftTree, data, l, mid)
	st.buildLazySegmentTree(rightTree, data, mid+1, r)

	st.tree[treeIndex] = st.merger(st.tree[leftTree], st.tree[rightTree])
}

// 给定指定节点下标，计算其左子树节点下标
func (st *LazySegmentTree) leftChild(index int) int {
	return index*2 + 1
}

// 给定指定节点下标，计算其右子树节点下标
func (st *LazySegmentTree) rightChild(index int) int {
	return index*2 + 2
}

// Query 查询区间范围 [queryL, queryR] 的值
func (st *LazySegmentTree) Query(queryL, queryR int) int {
	if queryL > queryR || queryL < 0 || queryR >= st.size {
		panic(fmt.Sprintf("illegal query0 range [%v, %v]", queryL, queryR))
	}
	return st.query(0, 0, st.size-1, queryL, queryR)
}

//func (st *LazySegmentTree) query0(root, l, r, queryL, queryR int) int {
//	// 区间 [queryL, queryR] 正好相等
//	if l == queryL && r == queryR {
//		return st.tree[root]
//	}
//
//	m := l + (r-l)/2
//
//	left := st.leftChild(root)
//	right := st.rightChild(root)
//
//	if st.lazy[root] != 0 {
//		st.tree[left] += (m - l + 1) * st.lazy[root]
//		st.tree[right] += (r - m) * st.lazy[root]
//		st.lazy[left] += st.lazy[root]
//		st.lazy[right] += st.lazy[root]
//		st.lazy[root] = 0
//	}
//
//	if queryR <= m {
//		return st.query0(left, l, m, queryL, queryR)
//	}
//
//	if queryL > m {
//		return st.query0(right, m+1, r, queryL, queryR)
//	}
//
//	lRet := st.query0(left, l, m, queryL, m)
//	rRet := st.query0(right, m+1, r, m+1, queryR)
//
//	return st.merger(lRet, rRet)
//}

// UpdateRange update segment tree from range [updateL, updateR], every element need to add diff
func (st *LazySegmentTree) UpdateRange(updateL, updateR int, diff int) {
	if updateL < 0 || updateR >= st.size || updateR < updateL {
		panic(fmt.Sprintf("update range [%v, %v] is illegal", updateL, updateR))
	}
	st.updateRange(0, 0, st.size-1, updateL, updateR, diff)
}

func (st *LazySegmentTree) lazyUpdate(root int, l int, r int) {
	if st.lazy[root] != 0 {

		// update self
		st.tree[root] += (r - l + 1) * st.lazy[root]

		// not a leaf node
		if l != r {
			st.lazy[st.leftChild(root)] += st.lazy[root]
			st.lazy[st.rightChild(root)] += st.lazy[root]
		}

		st.lazy[root] = 0
	}
}

//func (st *LazySegmentTree) updateRange0(root int, l int, r int, updateL int, updateR int, diff int) {
//	fmt.Println(fmt.Sprintf(">>> update range [%v, %v] = %v", updateL, updateR, diff))
//	if updateL == l && updateR == r {
//		// update self
//		st.tree[root] += (r - l + 1) * diff
//		st.lazy[root] += diff
//		return
//	}
//
//	mid := l + (r-l)/2
//
//	left, right := st.leftChild(root), st.rightChild(root)
//
//	// 更新子节点的值并传递懒惰标记值
//	if st.lazy[root] != 0 && l != r {
//		st.tree[left] += (mid - l + 1) * st.lazy[root]
//		st.tree[right] += (r - mid) * st.lazy[root]
//
//		st.lazy[left] += st.lazy[root]
//		st.lazy[right] += st.lazy[root]
//
//		st.lazy[root] = 0
//	}
//
//	if updateR <= mid {
//		st.updateRange0(left, l, mid, updateL, updateR, diff)
//	} else if updateL > mid+1 {
//		st.updateRange0(right, mid+1, r, updateL, updateR, diff)
//	} else {
//		st.updateRange0(left, l, mid, updateL, mid, diff)
//		st.updateRange0(right, mid+1, r, mid+1, updateR, diff)
//	}
//	st.tree[root] = st.merger(st.tree[left], st.tree[right])
//}

func (st *LazySegmentTree) updateRange(root int, l int, r int, updateL int, updateR int, diff int) {
	if updateL <= l && updateR >= r {
		st.tree[root] += (r - l + 1) * diff
		st.lazy[root] += diff
		return
	}

	mid := l + (r-l)/2

	left, right := st.leftChild(root), st.rightChild(root)

	// 更新子节点的值并传递懒惰标记值
	if st.lazy[root] != 0 && l != r {
		st.tree[left] += (mid - l + 1) * st.lazy[root]
		st.tree[right] += (r - mid) * st.lazy[root]
		st.lazy[left] += st.lazy[root]
		st.lazy[right] += st.lazy[root]
		st.lazy[root] = 0
	}

	if updateL <= mid {
		st.updateRange(left, l, mid, updateL, updateR, diff)
	}
	if updateR > mid {
		st.updateRange(right, mid+1, r, updateL, updateR, diff)
	}
	st.tree[root] = st.merger(st.tree[left], st.tree[right])
}

func (st *LazySegmentTree) query(root, l, r, queryL, queryR int) int {
	// [l, r] 正好是 [queryL, queryR] 的子集
	if queryL <= l && queryR >= r {
		return st.tree[root]
	}

	m := l + (r-l)/2

	left := st.leftChild(root)
	right := st.rightChild(root)

	if st.lazy[root] != 0 {
		st.tree[left] += (m - l + 1) * st.lazy[root]
		st.tree[right] += (r - m) * st.lazy[root]
		st.lazy[left] += st.lazy[root]
		st.lazy[right] += st.lazy[root]
		st.lazy[root] = 0
	}

	var lRet, rRet int
	if queryL <= m {
		lRet = st.query(left, l, m, queryL, queryR)
	}
	if queryR > m {
		rRet = st.query(right, m+1, r, queryL, queryR)
	}
	return st.merger(lRet, rRet)
}

type stNode struct {
	i int
	v int
	l int
	r int
}

func rootSTNode(st *LazySegmentTree) stNode {
	return stNode{
		i: 0,
		v: st.Query(0, st.size-1),
		l: 0,
		r: st.size - 1,
	}
}

func (node *stNode) getChildren(st *LazySegmentTree) (bool, *stNode, *stNode) {
	if node.l == node.r {
		return false, nil, nil
	}
	mid := node.l + ((node.r - node.l) >> 1)
	return true,
		&stNode{i: node.i*2 + 1, v: st.Query(node.l, mid), l: node.l, r: mid},
		&stNode{i: node.i*2 + 2, v: st.Query(mid+1, node.r), l: mid + 1, r: node.r}
}

func (st *LazySegmentTree) RankVisit(fn func(index, v int, l, r int)) {
	q := queue.NewArrayQueue[stNode]()
	q.Enqueue(rootSTNode(st))
	st.rankVisit(q, fn)
}

func (st *LazySegmentTree) rankVisit(q queue.Interface[stNode], fn func(index, v int, l, r int)) {
	for !q.IsEmpty() {
		node := q.Dequeue()
		fn(node.i, node.v, node.l, node.r)
		ok, left, right := node.getChildren(st)
		if ok {
			q.Enqueue(*left)
			q.Enqueue(*right)
		}
	}
}

func (st *LazySegmentTree) ReversedRankVisit(fn func(index, v int, l, r int)) {
	q := queue.NewArrayQueue[stNode]()
	q.Enqueue(rootSTNode(st))
	st.reversedRankVisit(q, fn)
}

func (st *LazySegmentTree) reversedRankVisit(q queue.Interface[stNode], fn func(index, v int, l, r int)) {
	s := stack.New[stNode]()
	for !q.IsEmpty() {
		node := q.Dequeue()
		s.Push(node)

		ok, left, right := node.getChildren(st)
		if ok {
			q.Enqueue(*right)
			q.Enqueue(*left)
		}
	}

	for !s.IsEmpty() {
		node := s.Pop()
		fn(node.i, node.v, node.l, node.r)
	}
}
