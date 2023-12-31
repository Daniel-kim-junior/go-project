package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

var n int
var m int
var r int
var rst int

type GraphNode struct {
	idx      int
	cost     int
	itemCost int
}

func (g GraphNode) Compare(o GraphNode) bool {
	return g.cost < o.cost
}

func main() {
	read := bufio.NewReader(os.Stdin)
	line, _, _ := read.ReadLine()
	input := strings.Split(string(line), " ")
	n, _ = strconv.Atoi(input[0])
	m, _ = strconv.Atoi(input[1])
	r, _ = strconv.Atoi(input[2])
	items := make([]int, n+1)
	line, _, _ = read.ReadLine()
	input = strings.Split(string(line), " ")
	var item int
	for i := 1; i <= n; i++ {
		item, _ = strconv.Atoi(input[i-1])
		items[i] = item
	}
	var from int
	var to int
	var cost int
	graph := make([][]int, n+1)
	row := make([]int, (n+1)*(n+1))
	for i := 1; i <= n; i++ {
		graph[i] = row[i*(n+1) : (i+1)*(n+1)]
	}

	for i := 0; i < r; i++ {
		in, _, _ := read.ReadLine()
		input = strings.Split(string(in), " ")
		from, _ = strconv.Atoi(input[0])
		to, _ = strconv.Atoi(input[1])
		cost, _ = strconv.Atoi(input[2])
		graph[from][to] = cost
		graph[to][from] = cost
	}
	for i := 1; i <= n; i++ {
		BFS(i, graph, items)
	}

	fmt.Print(rst)

}
func BFS(cur int, graph [][]int, items []int) {
	heap := &Heap[GraphNode]{Size: 0}
	DP := make([]int, n+1)
	var INF int = 123456789
	for i := 1; i <= n; i++ {
		DP[i] = INF
	}
	DP[cur] = 0
	l := 0
	heap.Insert(GraphNode{cur, 0, 0})
	for heap.Size != 0 {
		node := heap.Pop()

		for i := 1; i <= n; i++ {
			if node.Val.cost+graph[node.Val.idx][i] > m {
				continue
			}
			if DP[i] > graph[node.Val.idx][i]+node.Val.cost && graph[node.Val.idx][i] != 0 {
				DP[i] = graph[node.Val.idx][i] + node.Val.cost
				heap.Insert(GraphNode{i, node.Val.cost + graph[node.Val.idx][i], 0})
			}
		}
	}
	for i := 1; i <= n; i++ {
		if DP[i] == INF {
			continue
		}
		l += items[i]
	}
	if l > rst {
		rst = l
	}
}
