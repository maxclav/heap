package heap

import "fmt"

// Min is a min heap.
type Min struct {
	items []int
}

// NewMax return a new min heap.
func NewMin() *Min {
	return &Min{
		items: make([]int, 0),
	}
}

// Insert inserts `val` in the min heap.
func (h *Min) Insert(val int) {
	h.items = append(h.items, val)
	h.heapifyUp(h.Size() - 1)
}

// Get returns the minimum value of the min heap
// (value on top of the heao) without removing it.
func (h *Min) Get() int {
	if h.IsEmpty() {
		return -1
	}
	return h.items[0]
}

// Extract removes the minimum value of the min heap
// (value on top of the heao) and returns it.
func (h *Min) Extract() int {
	if h.IsEmpty() {
		return -1
	}
	max := h.items[0]
	h.items[0] = h.items[h.Size()-1]
	h.items = h.items[:h.Size()-1]
	h.heapifyDown(0)
	return max
}

func (h *Min) heapifyUp(pos int) {
	for h.hasParentOf(pos) {
		parentPos := IndexParent(pos)
		if h.items[pos] < h.parentOf(pos) {
			h.swap(pos, parentPos)
		}
		pos = parentPos
	}
}

func (h *Min) heapifyDown(pos int) {
	if pos < h.Size() {
		if h.hasBothChildrenOf(pos) {
			if h.leftChildOf(pos) <= h.rightChildOf(pos) {
				h.swapLeftChildIfSmaller(pos)
			} else {
				h.swapRightChildIfSmaller(pos)
			}
		} else if h.hasLeftChildOf(pos) {
			h.swapLeftChildIfSmaller(pos)
		} else if h.hasRightChildOf(pos) {
			h.swapRightChildIfSmaller(pos)
		}
	}
}

func (h *Min) swapLeftChildIfSmaller(pos int) {
	if pos < h.Size() && h.hasLeftChildOf(pos) {
		current, leftChild := h.items[pos], h.leftChildOf(pos)
		if current > leftChild {
			leftChildPos := IndexLeftChild(pos)
			h.swap(leftChildPos, pos)
			h.heapifyDown(leftChildPos)
		}
	}
}

func (h *Min) swapRightChildIfSmaller(pos int) {
	if pos < h.Size() && h.hasRightChildOf(pos) {
		current, rightChild := h.items[pos], h.rightChildOf(pos)
		if current > rightChild {
			rightChildPos := IndexRightChild(pos)
			h.swap(rightChildPos, pos)
			h.heapifyDown(rightChildPos)
		}
	}
}

func (h *Min) parentOf(pos int) int {
	if !h.hasParentOf(pos) {
		return -1
	}

	return h.items[IndexParent(pos)]
}

func (h *Min) leftChildOf(pos int) int {
	if !h.hasLeftChildOf(pos) {
		return -1
	}
	return h.items[IndexLeftChild(pos)]
}

func (h *Min) rightChildOf(pos int) int {
	if !h.hasRightChildOf(pos) {
		return -1
	}
	return h.items[IndexRightChild(pos)]
}

func (h *Min) hasParentOf(pos int) bool {
	return pos > 0
}

func (h *Min) hasBothChildrenOf(pos int) bool {
	return h.hasLeftChildOf(pos) && h.hasRightChildOf(pos)
}

func (h *Min) hasLeftChildOf(pos int) bool {
	idx := IndexLeftChild(pos)
	return idx < h.Size()
}

func (h *Min) hasRightChildOf(pos int) bool {
	idx := IndexRightChild(pos)
	return idx < h.Size()
}

func (h *Min) swap(x, y int) {
	h.items[x], h.items[y] = h.items[y], h.items[x]
}

// Size returns the size of the min heap.
func (h *Min) Size() int {
	return len(h.items)
}

// IsEmpty returns whether the min heap
// is empty (true) or not (false).
func (h *Min) IsEmpty() bool {
	return h.Size() == 0
}

// Println prints the min heap.
// TODO: Print it as a tree.
func (h *Min) Println() {
	fmt.Printf("Min Heap: %v\n", h.items)
}
