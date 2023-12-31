package godatastruct

type HeapNode[T Comparable[T]] struct {
	val T
}
type Comparable[T any] interface {
	Compare(o T) bool
}
type Heap[T Comparable[T]] struct {
	Array []HeapNode[T]
	Size  int
}

func (h *Heap[T]) Insert(value T) {
	node := &HeapNode[T]{val: value}
	h.Array[h.Size+1] = *node
	h.Size = h.Size + 1
	h.HeapifyUp(h.Size + 1)
}

func (h *Heap[T]) HeapifyUp(tail int) {
	cur := tail
	tailVal := h.Array[tail]
	for ; cur/2 != 0 && tailVal.val.Compare(h.Array[cur].val); cur /= 2 {
		h.Array[cur/2], tailVal = tailVal, h.Array[cur]
	}
}

func Zero[T Comparable[T]]() T {
	return *new(T)
}

func (h *Heap[T]) Pop() *HeapNode[T] {
	if h.Size == 0 {
		return nil
	}
	node := h.Array[1]
	h.Array[1], h.Array[h.Size].val = h.Array[h.Size], Zero[T]()
	h.Size--
	h.HeapifyDown()
	return &node
}

func (h *Heap[T]) HeapifyDown() {
	for curIdx := 1; curIdx*2 < h.Size; {
		if !h.Array[curIdx].val.Compare(h.Array[curIdx*2].val) && !h.Array[curIdx].val.Compare(h.Array[curIdx*2+1].val) {
			break
		}
		if h.Array[curIdx].val.Compare(h.Array[curIdx*2].val) {
			h.Array[curIdx], h.Array[curIdx*2] = h.Array[curIdx*2], h.Array[curIdx]
			curIdx *= 2
		} else {
			h.Array[curIdx], h.Array[curIdx*2+1] = h.Array[curIdx*2+1], h.Array[curIdx]
			curIdx = curIdx*2 + 1
		}
	}
}
