package heap

import (
	"fmt"
	"sync"
)

type BinaryHeap struct {
	data []int
	lock sync.RWMutex
}

var ErrorValueNotFound = fmt.Errorf("key not found")
var ErrorEmptyHeap = fmt.Errorf("empty heap")

func (h *BinaryHeap) up(idx int) {
	for idx > 0 {
		parentIdx := parent(idx)
		if parentIdx >= 0 && h.data[idx] < h.data[parentIdx] {
			h.data[parentIdx], h.data[idx] = h.data[idx], h.data[parentIdx]
			idx = parentIdx
		} else {
			return
		}
	}
}

func (h *BinaryHeap) down(idx int) {
	for idx < len(h.data)/2 {
		minIdx := idx
		leftChildIdx := leftChild(idx)
		rightChildIdx := rightChild(idx)

		if leftChildIdx < len(h.data) && h.data[minIdx] > h.data[leftChildIdx] {
			minIdx = leftChildIdx
		}

		if rightChildIdx < len(h.data) && h.data[minIdx] > h.data[rightChildIdx] {
			minIdx = rightChildIdx
		}

		if idx == minIdx {
			return
		}
		h.data[idx], h.data[minIdx] = h.data[minIdx], h.data[idx]
		idx = minIdx
	}
}

// Length return the number of items in the Heap
func (h *BinaryHeap) Length() int {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return len(h.data)
}

// Insert insert an item into a Heap
func (h *BinaryHeap) Insert(val int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.data = append(h.data, val)
	h.up(len(h.data) - 1)
}

// Extract return the smallest item
// and remove it from a Heap
func (h *BinaryHeap) Extract() (int, error) {
	h.lock.Lock()
	defer h.lock.Unlock()

	if len(h.data) == 0 {
		return 0, ErrorEmptyHeap
	}

	item := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.down(0)

	return item, nil
}

// Top return to top item of the Heap without remove it
func (h *BinaryHeap) Top() (int, error) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	if len(h.data) == 0 {
		return 0, ErrorEmptyHeap
	}

	return h.data[0], nil
}

// Delete will delete item with value 'val' from the Heap
func (h *BinaryHeap) Delete(val int) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	if len(h.data) == 0 {
		return ErrorEmptyHeap
	}

	deleteIndex := -1

	for i := 0; i < len(h.data); i++ {
		if val == h.data[i] {
			deleteIndex = i
			break
		}
	}

	if deleteIndex == -1 {
		return ErrorValueNotFound
	}

	h.data[deleteIndex] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	if deleteIndex > 0 && h.data[deleteIndex] < h.data[parent(deleteIndex)] {
		h.up(deleteIndex)
	} else {
		h.down(deleteIndex)
	}

	return nil
}

func NewBinaryHeap(arr []int) *BinaryHeap {
	res := BinaryHeap{
		data: arr,
	}

	start := len(res.data)/2 - 1
	for start >= 0 {
		res.down(start)
		start -= 1
	}

	return &res
}
