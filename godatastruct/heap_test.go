package godatastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CompareInt struct {
	val int
}

func (c CompareInt) compare(o CompareInt) bool {
	return o.val < c.val
}

func TestHeap(t *testing.T) {
	heap := &Heap[CompareInt]{make([]HeapNode[CompareInt], 10), 0, 0}
	assert.Equal(t, CompareInt{0}, heap.Pop().val)

	// assert.Equal(t, CompareInt(10), heap.Pop().val)
}
