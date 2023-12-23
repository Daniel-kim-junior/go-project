package singlelinkedlist

type Node[T any] struct {
	next *Node[T]
	val  T
}

type LinkedList[T any] struct {
	rootNode *Node[T]
	tailNode *Node[T]
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{nil, value}
	if l.rootNode == nil {
		l.rootNode = node
		l.tailNode = node
		return
	}

	l.tailNode.next = node
	l.tailNode = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{val: value}
	if l.rootNode == nil {
		l.rootNode = node
		l.tailNode = node
		return
	}
	node.next = l.rootNode
	l.rootNode = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.rootNode
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tailNode
}
