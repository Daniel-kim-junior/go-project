package singlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]
	assert.Nil(t, l.rootNode)
	assert.Nil(t, l.tailNode)
	l.PushBack(1)

	assert.NotNil(t, l.rootNode)
	assert.NotNil(t, l.tailNode)
	assert.Equal(t, 1, l.Front().val)
	assert.Equal(t, 1, l.Back().val)

	l.PushBack(2)
	assert.Equal(t, 1, l.Front().val)
	assert.Equal(t, 2, l.Back().val)
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]
	l.PushFront(1)
	assert.NotNil(t, l.rootNode)
	assert.NotNil(t, l.tailNode)
	assert.Equal(t, 1, l.Front().val)
	assert.Equal(t, 1, l.Back().val)
	l.PushFront(2)
	assert.Equal(t, 2, l.Front().val)
	assert.Equal(t, 1, l.Back().val)
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
	assert.Equal(t, 1, l.GetAt(0).val)
	assert.Nil(t, l.GetAt(1))
}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.InsertAfter(l.GetAt(1), 7)
	assert.Equal(t, 4, l.Size2())
	assert.Equal(t, 7, l.GetAt(2).val)
	assert.Equal(t, 3, l.tailNode.val)

	notExist := &Node[int]{val: 100}
	l.InsertAfter(notExist, 8)
	assert.Equal(t, 4, l.Size2())

}
