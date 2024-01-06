package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Juewly struct {
	weight int
	value  int
	index  int // 힙 내에서의 인덱스
}

type JPQ []*Juewly
type MINJPQ []*Juewly

func (jpq MINJPQ) Len() int {
	return len(jpq)
}

func (pq MINJPQ) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}
func (pq MINJPQ) PeekValue() int {
	return pq[0].value
}
func (pq MINJPQ) PeekWeight() int {
	return pq[0].weight
}

// Swap은 i번째와 j번째의 Item을 교환합니다.
func (pq MINJPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push는 PriorityQueue에 Item을 추가합니다.
func (pq *MINJPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Juewly)
	item.index = n
	*pq = append(*pq, item)
}

// Pop은 PriorityQueue에서 가장 우선순위가 높은 Item을 제거하고 반환합니다.
func (pq *MINJPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1 // 마지막으로 이동한 요소의 index를 초기화
	*pq = old[0 : n-1]
	return item
}

func (jpq JPQ) Len() int {
	return len(jpq)
}

func (pq JPQ) Less(i, j int) bool {
	return pq[i].value > pq[j].value
}
func (pq JPQ) PeekValue() int {
	return pq[0].value
}
func (pq JPQ) PeekWeight() int {
	return pq[0].weight
}

// Swap은 i번째와 j번째의 Item을 교환합니다.
func (pq JPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push는 PriorityQueue에 Item을 추가합니다.
func (pq *JPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Juewly)
	item.index = n
	*pq = append(*pq, item)
}

// Pop은 PriorityQueue에서 가장 우선순위가 높은 Item을 제거하고 반환합니다.
func (pq *JPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1 // 마지막으로 이동한 요소의 index를 초기화
	*pq = old[0 : n-1]
	return item
}

type JSlice []int

func (js JSlice) Len() int {
	return len(js)
}

func (js JSlice) Less(a int, b int) bool {

	return js[a] < js[b]
}

func (s JSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	in, _, _ := reader.ReadLine()
	s := strings.Split(string(in), " ")
	N, _ := strconv.Atoi(s[0])
	K, _ := strconv.Atoi(s[1])
	arr := make(MINJPQ, 0)
	heap.Init(&arr)

	rst := int64(0)
	for i := 0; i < N; i++ {
		in, _, _ = reader.ReadLine()
		s = strings.Split(string(in), " ")
		M, _ := strconv.Atoi(string(s[0]))
		V, _ := strconv.Atoi(string(s[1]))
		heap.Push(&arr, &Juewly{weight: M, value: V})
	}
	begs := make(JSlice, K)

	for i := 0; i < K; i++ {
		in, _, _ = reader.ReadLine()
		L, _ := strconv.Atoi(string(in))
		begs[i] = L
	}
	sort.Sort(begs)
	begHeap := make(JPQ, 0)
	heap.Init(&begHeap)
	for i := 0; i < len(begs); i++ {

		for arr.Len() > 0 && arr.PeekWeight() <= begs[i] {
			l := heap.Pop(&arr).(*Juewly)
			heap.Push(&begHeap, l)
		}
		if begHeap.Len() > 0 {
			l := heap.Pop(&begHeap).(*Juewly)
			rst += int64(l.value)
		} else if arr.Len() == 0 {
			break
		}
	}
	fmt.Fprintln(writer, rst)
}
