package collection

// Item is the element in the heap
type Item interface {
	Less(other Item) bool
}

// Heap is Heap data structure
type Heap struct {
	items []Item
	size  int
}

// NewHeap is the constructor of Heap
func NewHeap() *Heap {
	return &Heap{
		items: make([]Item, 0),
		size:  0,
	}
}

// Add pushs the item to the heap and adjust elements
func (heap *Heap) Add(item Item) {
	heap.items = append(heap.items, item)
	idx := heap.size
	heap.size++

	for idx > 0 {
		parentIdx := (idx+1)/2 - 1
		parent := heap.items[parentIdx]
		if item.Less(parent) {
			heap.swap(idx, parentIdx)
			idx = parentIdx
		} else {
			break
		}
	}
}

// Pop removes and returns the first element of the heap
func (heap *Heap) Pop() (Item, bool) {
	if heap.size == 0 {
		return nil, false
	}
	top := heap.items[0]
	heap.swap(0, heap.size-1)
	heap.items = heap.items[:heap.size-1]
	heap.size--

	idx := 0

	for idx < heap.size {
		current := heap.items[idx]
		leftChildIdx, rightChildIdx := idx<<1+1, idx<<1+2
		if leftChildIdx >= heap.size {
			break
		}
		leftChild := heap.items[leftChildIdx]
		child := leftChild
		childIdx := leftChildIdx

		if rightChildIdx < heap.size {
			rightChild := heap.items[rightChildIdx]
			if rightChild.Less(leftChild) {
				child = rightChild
				childIdx = rightChildIdx
			}
		}

		if current.Less(child) {
			break
		}
		heap.swap(idx, childIdx)
		idx = childIdx
	}

	return top, true
}

// Top returns the first element of heap
func (heap *Heap) Top() (Item, bool) {
	if heap.size > 0 {
		return heap.items[0], true
	}
	return nil, false
}

// Size returns the total elements in the heap
func (heap *Heap) Size() int {
	return heap.size
}

// func (heap *Heap) resize() {
// 	newItems := make([]Item, 0, 2*heap.size)
// 	for i, item := range heap.items {
// 		newItems[i] = item
// 	}
// 	heap.items = newItems
// }

func (heap *Heap) isFull() bool {
	return len(heap.items) == heap.size
}

func (heap *Heap) swap(i, j int) {
	temp := heap.items[i]
	heap.items[i] = heap.items[j]
	heap.items[j] = temp
}
