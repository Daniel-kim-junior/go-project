package godatastruct

type HeapNode[T Comparable[T]] struct {
	Val T
}
type Comparable[T any] interface {
	Compare(o T) bool
}
type Heap[T Comparable[T]] struct {
	Array []HeapNode[T]
	Size  int
}

func (h *Heap[T]) Insert(Value T) {
	node := HeapNode[T]{Val: Value}
	if h.Array == nil {
		h.Array = make([]HeapNode[T], 0)
	}
	h.Array = append(h.Array, node)
	h.Size = h.Size + 1
	h.HeapifyUp(h.Size)
}

func (h *Heap[T]) HeapifyUp(tail int) {
	cur := tail
	for ; cur/2 != 0 && h.Array[cur/2-1].Val.Compare(h.Array[cur-1].Val); cur /= 2 {
		h.Array[cur/2-1], h.Array[cur-1] = h.Array[cur-1], h.Array[cur/2-1]
	}
}

func Zero[T Comparable[T]]() T {
	return *new(T)
}

func (h *Heap[T]) Pop() *HeapNode[T] {
	if h.Size == 0 {
		return nil
	}
	node := h.Array[0]
	h.Size--
	if h.Size == 0 {
		h.Array = h.Array[1:]
	} else {
		h.Array[0] = h.Array[len(h.Array)-1]
		h.Array = h.Array[:len(h.Array)-1]
		h.HeapifyDown()
	}
	return &node
}

func (h *Heap[T]) HeapifyDown() {
	for curIdx := 1; curIdx*2 < h.Size; {
		if !h.Array[curIdx-1].Val.Compare(h.Array[curIdx*2-1].Val) && !h.Array[curIdx-1].Val.Compare(h.Array[curIdx*2].Val) {
			break
		}
		if h.Array[curIdx-1].Val.Compare(h.Array[curIdx*2-1].Val) {
			h.Array[curIdx-1], h.Array[curIdx*2-1] = h.Array[curIdx*2-1], h.Array[curIdx-1]
			curIdx *= 2
			curIdx -= 1
		} else {
			h.Array[curIdx-1], h.Array[curIdx*2] = h.Array[curIdx*2], h.Array[curIdx-1]
			curIdx *= 2
		}
	}
}
