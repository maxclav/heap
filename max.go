package heap

import "fmt"

// Max is a max heap.
type Max struct {
	items []int
}

// NewMax return a new max heap.
func NewMax() *Max {
	return &Max{
		items: make([]int, 0),
	}
}

// Insert inserts `val` in the max heap.
func (h *Max) Insert(val int) {
	h.items = append(h.items, val)
	h.heapifyUp(h.Size() - 1)
}

// Get returns the maximum value of the max heap
// (value on top of the heao) without removing it.
func (h *Max) Get() int {
	if h.IsEmpty() {
		return -1
	}
	return h.items[0]
}

// Extract removes the maximum value of the max heap
// (value on top of the heao) and returns it.
func (h *Max) Extract() int {
	if h.IsEmpty() {
		return -1
	}
	max := h.items[0]
	h.items[0] = h.items[h.Size()-1]
	h.items = h.items[:h.Size()-1]
	h.heapifyDown(0)
	return max
}

func (h *Max) heapifyUp(pos int) {
	for h.hasParentOf(pos) {
		parentPos := IndexParent(pos)
		if h.items[pos] > h.parentOf(pos) {
			h.swap(pos, parentPos)
		}
		pos = parentPos
	}
}

func (h *Max) heapifyDown(pos int) {
	if pos < h.Size() {
		if h.hasBothChildrenOf(pos) {
			if h.leftChildOf(pos) > h.rightChildOf(pos) {
				h.swapLeftChildIfBigger(pos)
			} else {
				h.swapRightChildIfBigger(pos)
			}
		} else if h.hasLeftChildOf(pos) {
			h.swapLeftChildIfBigger(pos)
		} else if h.hasRightChildOf(pos) {
			h.swapRightChildIfBigger(pos)
		}
	}
}

func (h *Max) swapLeftChildIfBigger(pos int) {
	if pos < h.Size() && h.hasLeftChildOf(pos) {
		current, leftChild := h.items[pos], h.leftChildOf(pos)
		if current < leftChild {
			leftChildPos := IndexLeftChild(pos)
			h.swap(leftChildPos, pos)
			h.heapifyDown(leftChildPos)
		}
	}
}

func (h *Max) swapRightChildIfBigger(pos int) {
	if pos < h.Size() && h.hasRightChildOf(pos) {
		current, rightChild := h.items[pos], h.rightChildOf(pos)
		if current < rightChild {
			rightChildPos := IndexRightChild(pos)
			h.swap(rightChildPos, pos)
			h.heapifyDown(rightChildPos)
		}
	}
}

func (h *Max) parentOf(pos int) int {
	if !h.hasParentOf(pos) {
		return -1
	}

	return h.items[IndexParent(pos)]
}

func (h *Max) leftChildOf(pos int) int {
	if !h.hasLeftChildOf(pos) {
		return -1
	}
	return h.items[IndexLeftChild(pos)]
}

func (h *Max) rightChildOf(pos int) int {
	if !h.hasRightChildOf(pos) {
		return -1
	}
	return h.items[IndexRightChild(pos)]
}

func (h *Max) hasParentOf(pos int) bool {
	return pos > 0
}

func (h *Max) hasBothChildrenOf(pos int) bool {
	return h.hasLeftChildOf(pos) && h.hasRightChildOf(pos)
}

func (h *Max) hasLeftChildOf(pos int) bool {
	idx := IndexLeftChild(pos)
	return idx < h.Size()
}

func (h *Max) hasRightChildOf(pos int) bool {
	idx := IndexRightChild(pos)
	return idx < h.Size()
}

func (h *Max) swap(x, y int) {
	h.items[x], h.items[y] = h.items[y], h.items[x]
}

// Size returns the size of the max heap.
func (h *Max) Size() int {
	return len(h.items)
}

// IsEmpty returns whether the max heap
// is empty (true) or not (false).
func (h *Max) IsEmpty() bool {
	return h.Size() == 0
}

// Println prints the max heap.
// TODO: Print it as a tree.
func (h *Max) Println() {
	fmt.Printf("Max Heap: %v\n", h.items)
}
