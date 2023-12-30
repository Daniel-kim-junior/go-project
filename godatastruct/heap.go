package godatastruct

type HeapNode[T comparable] struct {
	val T
}
type comparable interface {
	compare(o interface{}) bool
}
type Heap[T comparable] struct {
	array []HeapNode[T]
	size  int
}

func (h Heap[T]) Insert(value T) {
	node := &HeapNode[T]{val: value}

	var tail int
	if h.size%2 == 0 {
		h.array[h.size+1] = *node
		tail = h.size + 1
	} else {
		h.array[2*(h.size-1)] = *node
		tail = 2 * (h.size - 1)
	}
	h.size++
	h.HeapifyUp(tail)
}

func (h Heap[T]) HeapifyUp(tail int) {
	cur := tail
	tailVal := h.array[tail]
	for ; cur/2 != 0 && tailVal.val.compare(h.array[cur].val); cur /= 2 {
		h.array[cur/2], tailVal = tailVal, h.array[cur]
	}
}

func (h Heap[T]) Pop() *HeapNode[T] {
	return &h.array[1]
}
