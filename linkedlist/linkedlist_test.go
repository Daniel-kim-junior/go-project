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
