package singlelinkedlist

type Node[T any] struct {
	next *Node[T]
	val  T
}

type LinkedList[T any] struct {
	rootNode *Node[T]
	tailNode *Node[T]
	count    int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{nil, value}
	l.count++
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
	l.count++
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

func (l *LinkedList[T]) Size() int {
	node := l.rootNode
	var count int = 0

	for ; node != nil; node = node.next {
		count++
	}

	return count
}

func (l *LinkedList[T]) Size2() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Size2() {
		return nil
	}

	i := 0
	for node := l.rootNode; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{nil, value}
	node.next, newNode.next = newNode, node.next
	l.count++
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.rootNode

	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}

	return false
}
