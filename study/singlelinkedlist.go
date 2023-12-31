package study

type Node[T any] struct {
	next *Node[T]
	Val  T
}

type LinkedList[T any] struct {
	RootNode *Node[T]
	TailNode *Node[T]
	count    int
}

func (l *LinkedList[T]) PushBack(Value T) {
	node := &Node[T]{nil, Value}
	l.count++
	if l.RootNode == nil {
		l.RootNode = node
		l.TailNode = node
		return
	}

	l.TailNode.next = node
	l.TailNode = node
}

func (l *LinkedList[T]) PushFront(Value T) {
	node := &Node[T]{Val: Value}
	l.count++
	if l.RootNode == nil {
		l.RootNode = node
		l.TailNode = node
		return
	}
	node.next = l.RootNode
	l.RootNode = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.RootNode
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.TailNode
}

func (l *LinkedList[T]) Size() int {
	node := l.RootNode
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
	for node := l.RootNode; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], Value T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{nil, Value}
	node.next, newNode.next = newNode, node.next
	l.count++
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.RootNode

	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}

	return false
}

func (l *LinkedList[T]) insertBefore(node *Node[T], Value T) {
	if !l.isIncluded(node) {
		return
	}

}
