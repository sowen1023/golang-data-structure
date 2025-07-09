package tree

import (
	"errors"

	"io.github.sowen.datastructure/list"
	"io.github.sowen.datastructure/types"
)

type MinHeap[E types.Comparable] struct {
	arr *list.ArrayList[E]
}

func NewMinHeap[E types.Comparable]() *MinHeap[E] {
	return &MinHeap[E]{arr: list.New[E]()}
}

func (h *MinHeap[E]) get(i int) E {
	return h.arr.Get(i)
}

func (h *MinHeap[E]) set(i int, e E) {
	h.arr.Set(i, e)
}

func (h *MinHeap[E]) IsEmpty() bool {
	return h.arr.IsEmpty()
}

func (h *MinHeap[E]) Size() int {
	return h.arr.Size()
}

func (h *MinHeap[E]) swap(i, j int) {
	t := h.get(i)
	h.set(i, h.get(j))
	h.set(j, t)
}

func (h *MinHeap[E]) FindMin() (E, error) {
	if h.IsEmpty() {
		var e E
		return e, errors.New("can't extract min value from empty array")
	}
	return h.arr.GetFirst(), nil
}

func (h *MinHeap[E]) Add(e E) {
	h.arr.AddLast(e)
	h.siftUp(h.arr.Size() - 1)
}

func (h *MinHeap[E]) siftUp(i int) {
	for i > 0 && h.get(i) < h.get(parent(i)) {
		h.swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MinHeap[E]) siftDown(i int) {
	for leftChild(i) < h.arr.Size() {
		j := leftChild(i)

		if r := rightChild(i); r < h.Size() && h.get(r) < h.get(j) {
			j = r
		}

		if h.get(j) > h.get(i) {
			break
		}

		h.swap(i, j)
		i = j
	}
}

func (h *MinHeap[E]) ExtractMin() (E, error) {
	min, err := h.FindMin()
	if err != nil {
		var e E
		return e, err
	}

	h.swap(0, h.Size()-1)
	h.arr.RemoveLast()
	h.siftDown(0)

	return min, nil
}
