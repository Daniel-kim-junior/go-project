package godatastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CompareInt struct {
	Val int
}

func (c CompareInt) Compare(o CompareInt) bool {
	return o.Val < c.Val
}

func TestHeap(t *testing.T) {
	heap := &Heap[CompareInt]{Size: 0}
	assert.Nil(t, heap.Pop())
	heap.Insert(CompareInt{10})
	assert.Equal(t, 1, heap.Size)
	assert.Equal(t, CompareInt{10}, heap.Pop().Val)
	heap.Insert(CompareInt{1})
	heap.Insert(CompareInt{7})
	heap.Insert(CompareInt{6})
	heap.Insert(CompareInt{1})
	assert.Equal(t, CompareInt{1}, heap.Pop().Val)
	assert.Equal(t, CompareInt{1}, heap.Pop().Val)
	assert.Equal(t, CompareInt{6}, heap.Pop().Val)
	assert.Equal(t, CompareInt{7}, heap.Pop().Val)
}
