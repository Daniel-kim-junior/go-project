package study

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]
	assert.Nil(t, l.RootNode)
	assert.Nil(t, l.TailNode)
	l.PushBack(1)

	assert.NotNil(t, l.RootNode)
	assert.NotNil(t, l.TailNode)
	assert.Equal(t, 1, l.Front().Val)
	assert.Equal(t, 1, l.Back().Val)

	l.PushBack(2)
	assert.Equal(t, 1, l.Front().Val)
	assert.Equal(t, 2, l.Back().Val)
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]
	l.PushFront(1)
	assert.NotNil(t, l.RootNode)
	assert.NotNil(t, l.TailNode)
	assert.Equal(t, 1, l.Front().Val)
	assert.Equal(t, 1, l.Back().Val)
	l.PushFront(2)
	assert.Equal(t, 2, l.Front().Val)
	assert.Equal(t, 1, l.Back().Val)
}

func TestSize(t *testing.T) {
	var l LinkedList[int]
	l.PushFront(1)
	l.PushBack(1)
	l.PushFront(2)
	assert.Equal(t, 3, l.Size())
}

func TestSize2(t *testing.T) {
	var l LinkedList[int]
	l.PushFront(1)
	assert.Equal(t, 1, l.Size2())
	assert.Equal(t, 1, l.GetAt(0).Val)
	assert.Nil(t, l.GetAt(1))
}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.InsertAfter(l.GetAt(1), 7)
	assert.Equal(t, 4, l.Size2())
	assert.Equal(t, 7, l.GetAt(2).Val)
	assert.Equal(t, 3, l.TailNode.Val)

	notExist := &Node[int]{Val: 100}
	l.InsertAfter(notExist, 8)
	assert.Equal(t, 4, l.Size2())
}
