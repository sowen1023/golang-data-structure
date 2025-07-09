package tree

import "io.github.sowen.datastructure/types"

// BST Binary search tree implementation.
type BST[K types.Comparable, V types.Comparable] struct {
	// The root of the tree
	root *BSTNode[K, V]
	// The tree size
	size int
}

// BSTNode The node element of the BST.
type BSTNode[K types.Comparable, V types.Comparable] struct {
	// The value
	key K
	// The value
	val V
	// The left child
	left *BSTNode[K, V]
	// The right child
	right *BSTNode[K, V]
}

// Contains returns true if the tree contains element.
func (t *BST[K, V]) Contains(key K) bool {
	return t.contains(t.root, key)
}

func (t *BST[K, V]) contains(node *BSTNode[K, V], key K) bool {
	if node == nil {
		return false
	}
	if node.key == key {
		return true
	} else if node.key > key {
		return t.contains(node.left, key)
	}
	return t.contains(node.right, key)
}

func (t *BST[K, V]) Get(k K) V {
	return t.get(t.root, k)
}

func (t *BST[K, V]) get(node *BSTNode[K, V], k K) V {
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

// Add the element to the tree.
func (t *BST[K, V]) Add(k K, v V) {
	t.root = t.add(t.root, k, v)
}

func (t *BST[K, V]) add(node *BSTNode[K, V], k K, v V) *BSTNode[K, V] {
	if node == nil {
		t.size++
		return &BSTNode[K, V]{k, v, nil, nil}
	}
	if node.key == k {
		node.val = v
		return node
	} else if node.key > k {
		node.left = t.add(node.left, k, v)
		return node
	} else {
		node.right = t.add(node.right, k, v)
		return node
	}
}

// Remove the element from the tree.
func (t *BST[K, V]) Remove(k K) {
	t.root = t.remove(t.root, k)
}

func (t *BST[K, V]) remove(node *BSTNode[K, V], k K) *BSTNode[K, V] {
	if node == nil {
		return node
	}
	if node.key > k {
		node.left = t.remove(node.left, k)
		return node
	} else if node.key < k {
		node.right = t.remove(node.right, k)
		return node
	} else {
		if node.left == nil {
			right := node.right
			node.right = nil
			t.size--
			return right
		} else if node.right == nil {
			left := node.left
			node.left = nil
			t.size--
			return left
		} else {
			successor := t.minNode(node.right)
			//successor.right = t.remove(node.right, successor.val)
			successor.right, successor.left = t.removeMin(node.right), node.left
			node.left, node.right = nil, nil
			return successor
		}
	}
}

// IsEmpty returns the if the tree is empty.
func (t *BST[K, V]) IsEmpty() bool {
	return t.size == 0
}

func (t *BST[K, V]) Size() int {
	return t.size
}

// Returns min value in the tree.
func (t *BST[K, V]) minimum() K {
	if t.root == nil {
		panic("tree is empty")
	}
	return t.minNode(t.root).key
}

func (t *BST[K, V]) minNode(node *BSTNode[K, V]) *BSTNode[K, V] {
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
func (t *BST[K, V]) maximum() K {
	if t.root == nil {
		panic("tree is empty")
	}
	return t.maxNode(t.root).key
}

func (t *BST[K, V]) maxNode(node *BSTNode[K, V]) *BSTNode[K, V] {
	if node == nil {
		return nil
	}
	max := node
	for max.right != nil {
		max = max.right
	}
	return max
}

func (t *BST[K, V]) RemoveMin() K {
	k := t.minimum()
	t.root = t.removeMin(t.root)
	return k
}

func (t *BST[K, V]) removeMin(node *BSTNode[K, V]) *BSTNode[K, V] {
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

func (t *BST[K, V]) RemoveMax() K {
	k := t.maximum()
	t.root = t.removeMax(t.root)
	return k
}

func (t *BST[K, V]) removeMax(node *BSTNode[K, V]) *BSTNode[K, V] {
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
