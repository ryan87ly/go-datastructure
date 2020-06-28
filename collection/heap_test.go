package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testItem int

func (item testItem) Less(other Item) bool {
	omi := other.(testItem)
	return item < omi
}

func TestAdd(t *testing.T) {
	heap := NewHeap()

	heap.Add(testItem(2))
	heap.Add(testItem(1))

	assert.Equal(t, 2, heap.Size(), "Heap size should be 2")
}

func TestTop(t *testing.T) {
	heap := NewHeap()

	heap.Add(testItem(2))
	heap.Add(testItem(4))
	heap.Add(testItem(3))
	heap.Add(testItem(1))
	heap.Add(testItem(5))

	result, err := heap.Top()
	assert.True(t, err, "Top item should exist")
	assert.Equal(t, testItem(1), result, "Top item should be 1")
}

func TestPop(t *testing.T) {
	heap := NewHeap()

	heap.Add(testItem(2))
	heap.Add(testItem(4))
	heap.Add(testItem(3))
	heap.Add(testItem(1))
	heap.Add(testItem(5))

	assert.Equal(t, 5, heap.Size(), "Size should be 5")

	result, _ := heap.Pop()
	assert.Equal(t, testItem(1), result, "Top item should be 1")
	result, _ = heap.Pop()
	assert.Equal(t, testItem(2), result, "Top item should be 2")
	result, _ = heap.Pop()
	assert.Equal(t, testItem(3), result, "Top item should be 3")
	result, _ = heap.Pop()
	assert.Equal(t, testItem(4), result, "Top item should be 4")

	heap.Add(testItem(4))
	result, _ = heap.Pop()
	assert.Equal(t, testItem(4), result, "Top item should be 4")
	result, _ = heap.Pop()
	assert.Equal(t, testItem(5), result, "Top item should be 5")

	heap.Add(testItem(10))
	result, _ = heap.Pop()
	assert.Equal(t, testItem(10), result, "Top item should be 10")
}

func TestAbleToHandleDuplications(t *testing.T) {
	heap := NewHeap()

	heap.Add(testItem(2))
	heap.Add(testItem(4))
	heap.Add(testItem(3))
	heap.Add(testItem(1))
	heap.Add(testItem(2))

	assert.Equal(t, 5, heap.Size(), "Size should be 5")

	result, _ := heap.Pop()
	assert.Equal(t, testItem(1), result, "Top item should be 1")
	result, _ = heap.Pop()
	assert.Equal(t, testItem(2), result, "Top item should be 2")
	result, _ = heap.Pop()
	assert.Equal(t, testItem(2), result, "Top item should be 2")
	result, _ = heap.Pop()
	assert.Equal(t, testItem(3), result, "Top item should be 3")

}
