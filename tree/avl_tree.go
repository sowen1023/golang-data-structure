package tree

import (
	"math"

	"io.github.sowen.datastructure/types"
)

// AVLTree implementation
type AVLTree[K types.Comparable, V any] struct {
	// The root of the tree
	root *Node[K, V]
	// The tree size
	size int
}

// Node the node element of the AVL
type Node[K types.Comparable, V any] struct {
	// The value
	key K
	// The value
	val V
	// The left child
	left *Node[K, V]
	// The right child
	right *Node[K, V]
	// The tree height
	height int
}

func (n *Node[K, V]) recalculateHeight() int {
	l, r := getHeight(n.left), getHeight(n.right)
	if l >= r {
		return 1 + l
	}
	return 1 + r
}

//func (t *AVLTree) rightRotate(y *Node) *Node {
//	x := y.left
//	t3 := x.right
//	x.right = y
//	y.left = t3
//	return x
//}

// LL - Right Rotation

// If a tree becomes unbalanced, when a node is inserted into the left subtree of the left subtree,
// then we perform a single right rotation.

// T1 < z < T2 < x < T3 < y < T4
//
//	   y
//	  / \
//	  x  T4          向右旋转                   x
//	 / \          - - - - - - >              /  \
//	 z  T3                                  z    y
//	/ \                                   / \   / \
// T1  T2                               T1  T2 T3 T4

func (t *AVLTree[K, V]) rightRotate(y *Node[K, V]) *Node[K, V] {
	x := y.left
	x.right, y.left = y, x.right

	// update height
	y.height = y.recalculateHeight()
	x.height = x.recalculateHeight()

	return x
}

// RR - Left Rotation

// If a tree becomes unbalanced, when a node is inserted into the right subtree of the right subtree,
// then we perform a single left rotation.

//  T1 > z > T2 > x > T3 > y > T4
//
//	   y
//	  / \
//	 T4   x              向左旋转                   x
//	     / \          - - - - - - >              /   \
//	    T3  z                                   y     z
//	       / \                                 / \   / \
//        T2  T1                              T4 T3 T2 T1

func (t *AVLTree[K, V]) leftRotate(y *Node[K, V]) *Node[K, V] {
	x := y.right
	x.left, y.right = y, x.left

	// update height
	y.height = y.recalculateHeight()
	x.height = x.recalculateHeight()
	return x
}

// RL - Right Left Rotation

// If a tree becomes unbalanced, when a node is inserted into the left subtree of the right subtree.
// This makes [y], an unbalanced node with balance factor 2.
//
// 1. First, we perform the right rotation along [x] node, make [x] the right subtree of its own left subtree [z].
// Now, [z] becomes the right subtree of [y]. Node [y] is still unbalanced because of the right subtree of its
// right subtree and requires a left rotation.
//
// 2. A left rotation is performed by making [z] the new root node of the subtree. [y] becomes the left subtree of its right subtree [z].

//  T1 < y < T2 < z < T3 < x < T4
//
//	   y                            	 y
//	  / \								/ \
//	 T1   x              向右旋转      T1    z         向左旋转  			 z
//	     / \          - - - - - - >   	   / \      - - - - - - >       /   \
//	    z   T4                            T2   x    			       y     x
//	   / \                                    / \					  / \   / \
//    T2  T3                              	T3  T4 				     T1 T2 T3 T4

func (t *AVLTree[K, V]) rightLeftRotate(y *Node[K, V]) *Node[K, V] {
	y.right = t.rightRotate(y.right)
	return t.leftRotate(y)
}

// LR - Left Right Rotation

// If a tree becomes unbalanced, when a node is inserted into the right subtree of the left subtree.
// This makes [y], an unbalanced node with balance factor 2.
//
// 1. First, we perform the left rotation along [x] node, make [x] left node of its own right subtree [z].
// Now, [z] becomes the left subtree of [y]. Node [y] is still unbalanced because of the left subtree of its
// left subtree and requires a right rotation.
//
// 2. A right rotation is performed by making [z] the new root node of the subtree. [y] becomes the right subtree of its left subtree [z].

//  T1 < x < T2 < z < T3 < y < T4
//
//	      y                            	     y
//	     / \								/ \
//	    x  T4          向左旋转             z   T4     向右旋转  		    z
//	   / \          - - - - - - >   	  / \      - - - - - - >        /  \
//	  T1  z                               x  T3     			       x     y
//	     / \                             / \					      / \   / \
//      T2 T3                           T1 T2  	 				     T1 T2 T3 T4

func (t *AVLTree[K, V]) leftRightRotate(y *Node[K, V]) *Node[K, V] {
	y.left = t.leftRotate(y.left)
	return t.rightRotate(y)
}

func getHeight[K types.Comparable, V any](n *Node[K, V]) int {
	if n == nil {
		return 0
	}
	return n.height
}

func getBalanceFactor[K types.Comparable, V any](n *Node[K, V]) int {
	return getHeight(n.left) - getHeight(n.right)
}

func (t *AVLTree[K, V]) IsValidBST() bool {
	return isValidBST(t.root)
}

func isValidBST[K types.Comparable, V any](n *Node[K, V]) bool {
	if n == nil {
		return true
	}
	if n.left != nil {
		return n.key > n.left.key
	}
	if n.right != nil {
		return n.key < n.right.key
	}
	return isValidBST(n.left) && isValidBST(n.right)
}

func (t *AVLTree[K, V]) IsBalanced() bool {
	return isBalanced(t.root)
}

func isBalanced[K types.Comparable, V any](n *Node[K, V]) bool {
	if n == nil {
		return true
	}
	f := getBalanceFactor(n)
	if math.Abs(float64(f)) > 1 {
		return false
	}
	return isBalanced(n.left) && isBalanced(n.right)
}

// Contains returns true if the tree contains element
func (t *AVLTree[K, V]) Contains(e K) bool {
	return t.contains(t.root, e)
}

func (t *AVLTree[K, V]) contains(node *Node[K, V], k K) bool {
	if node == nil {
		return false
	}
	if node.key == k {
		return true
	} else if node.key > k {
		return t.contains(node.left, k)
	}
	return t.contains(node.right, k)
}

func (t *AVLTree[K, V]) Get(k K) V {
	return t.get(t.root, k)
}

func (t *AVLTree[K, V]) get(node *Node[K, V], k K) V {
	if node == nil {
		var v V
		return v
	}
	if node.key == k {
		return node.val
	} else if node.key > k {
		return t.get(node.left, k)
	}
	return t.get(node.right, k)
}

// Add the element to the tree
func (t *AVLTree[K, V]) Add(k K, v V) {
	t.root = t.add(t.root, k, v)
}

func (t *AVLTree[K, V]) add(node *Node[K, V], k K, v V) *Node[K, V] {
	if node == nil {
		t.size++
		return &Node[K, V]{k, v, nil, nil, 1}
	}
	if node.key == k {
		node.val = v
	} else if node.key > k {
		node.left = t.add(node.left, k, v)
	} else {
		node.right = t.add(node.right, k, v)
	}

	// if node's height unchanged, rotation is not required
	h := node.recalculateHeight()
	if node.height == h {
		return node
	}
	node.height = h
	return t.rotate(node)
}

func (t *AVLTree[K, V]) rotate(node *Node[K, V]) *Node[K, V] {
	// calculate the balance f of the node
	f := getBalanceFactor(node)

	// LL - Right Rotation
	if f > 1 && getBalanceFactor(node.left) >= 0 {
		return t.rightRotate(node)
	}
	// RR - Left Rotation
	if f < -1 && getBalanceFactor(node.right) <= 0 {
		return t.leftRotate(node)
	}
	// LR - Left Right Rotation
	if f > 1 && getBalanceFactor(node.left) < 0 {
		return t.leftRightRotate(node)
	}
	// RL - Right Left Rotation
	if f < -1 && getBalanceFactor(node.right) > 0 {
		return t.rightLeftRotate(node)
	}
	return node
}

// Remove the element from the tree
func (t *AVLTree[K, V]) Remove(k K) {
	t.root = t.remove(t.root, k)
}

func (t *AVLTree[K, V]) remove(node *Node[K, V], k K) *Node[K, V] {
	if node == nil {
		return nil
	}

	var ret *Node[K, V]
	if node.key > k {
		node.left = t.remove(node.left, k)
		ret = node
	} else if node.key < k {
		node.right = t.remove(node.right, k)
		ret = node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			t.size--
			ret = right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			t.size--
			ret = left
		} else {
			successor := t.minNode(node.right)
			successor.right = t.remove(node.right, successor.key)
			// successor.right, successor.left = t.removeMin(node.right), node.left
			successor.left = node.left
			node.left, node.right = nil, nil
			ret = successor
		}
	}

	if ret == nil {
		return nil
	}

	ret.height = ret.recalculateHeight()
	return t.rotate(ret)
}

// IsEmpty returns the if the tree is empty
func (t *AVLTree[K, V]) IsEmpty() bool {
	return t.size == 0
}

func (t *AVLTree[K, V]) Size() int {
	return t.size
}

// Returns min value in the tree.
func (t *AVLTree[K, V]) minimum() K {
	if t.root == nil {
		panic("tree is empty")
	}
	return t.minNode(t.root).key
}

func (t *AVLTree[K, V]) minNode(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	min := node
	for min.left != nil {
		min = min.left
	}
	return min
}

// Returns max value in the tree.
func (t *AVLTree[K, V]) maximum() K {
	if t.root == nil {
		panic("tree is empty")
	}
	return t.maxNode(t.root).key
}

func (t *AVLTree[K, V]) maxNode(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	max := node
	for max.right != nil {
		max = max.right
	}
	return max
}

func (t *AVLTree[K, V]) RemoveMin() (k K) {
	k = t.minimum()
	t.root = t.removeMin(t.root)
	return
}

func (t *AVLTree[K, V]) removeMin(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.left == nil {
		r := node.right
		node.right = nil
		t.size--
		return r
	}
	node.left = t.removeMin(node.left)
	return node
}

func (t *AVLTree[K, V]) RemoveMax() K {
	k := t.maximum()
	t.root = t.removeMax(t.root)
	return k
}

func (t *AVLTree[K, V]) removeMax(node *Node[K, V]) *Node[K, V] {
	if node == nil {
		return nil
	}
	if node.right == nil {
		l := node.left
		node.left = nil
		t.size--
		return l
	}
	node.right = t.removeMax(node.right)
	return node
}

func (t *AVLTree[K, V]) PreOrderFunc(fn func(k K, v V)) {
	if t.root != nil {
		t.preOrderFunc(t.root, fn)
	}
}

func (t *AVLTree[K, V]) preOrderFunc(node *Node[K, V], fn func(k K, v V)) {
	if node != nil {
		fn(node.key, node.val)
		t.preOrderFunc(node.left, fn)
		t.preOrderFunc(node.right, fn)
	}
}

func (t *AVLTree[K, V]) InOrderFunc(fn func(k K, v V)) {
	t.inOrderFunc(t.root, fn)
}

func (t *AVLTree[K, V]) inOrderFunc(node *Node[K, V], fn func(k K, v V)) {
	if node != nil {
		t.inOrderFunc(node.left, fn)
		fn(node.key, node.val)
		t.inOrderFunc(node.right, fn)
	}
}
