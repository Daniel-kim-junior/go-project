package godatastruct

type HeapNode[T comparable[T]] struct {
	val T
}
type comparable[T any] interface {
	compare(o T) bool
}
type Heap[T comparable[T]] struct {
	Array   []HeapNode[T]
	Size    int
	TailIdx int
}

func (h Heap[T]) Insert(value T) {
	node := &HeapNode[T]{val: value}

	var tail int
	if h.Size%2 == 0 {
		h.Array[h.Size+1] = *node
		tail = h.Size + 1
	} else {
		h.Array[2*(h.Size-1)] = *node
		tail = 2 * (h.Size - 1)
	}
	h.Size++
	h.HeapifyUp(tail)
	h.TailIdx = tail
}

func (h Heap[T]) HeapifyUp(tail int) {
	cur := tail
	tailVal := h.Array[tail]
	for ; cur/2 != 0 && tailVal.val.compare(h.Array[cur].val); cur /= 2 {
		h.Array[cur/2], tailVal = tailVal, h.Array[cur]
	}
}

func (h Heap[T]) Pop() *HeapNode[T] {
	node := h.Array[1]
	h.Array[1] = h.Array[h.TailIdx]
	h.HeapifyDown()
	return &node
}

func (h Heap[T]) HeapifyDown() {
	curIdx := 1
	for true {
		if !h.Array[curIdx*2].val.compare(h.Array[curIdx].val) && !h.Array[curIdx*2+1].val.compare(h.Array[curIdx].val) {
			break
		}
		if h.Array[curIdx*2].val.compare(h.Array[curIdx].val) {
			h.Array[curIdx*2], h.Array[curIdx] = h.Array[curIdx], h.Array[curIdx*2]
			curIdx *= 2
			continue
		}

		if h.Array[curIdx*2+1].val.compare(h.Array[curIdx].val) {
			h.Array[curIdx*2+1], h.Array[curIdx] = h.Array[curIdx], h.Array[curIdx*2+1]
			curIdx = curIdx*2 + 1
		}

	}
	h.TailIdx = curIdx
}
