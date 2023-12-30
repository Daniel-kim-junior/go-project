package godatastruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type CompareInt int

func (i CompareInt) compare(o interface{}) bool {
	return int(i) < o.(int)
}

func TestHeap(t *testing.T) {
	heap := &Heap[CompareInt]{make([]HeapNode[CompareInt], 10), 0}
	heap.Insert(9)
	assert.Equal(t, CompareInt(9), heap.Pop().val)
}
