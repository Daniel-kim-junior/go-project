package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	ds "github.com/Daniel-kim-junior/go-project/godatasturct"
)

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
	return g.itemCost > o.itemCost
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
	graph := make([][]GraphNode, n+1)
	row := make([]GraphNode, (n+1)*(n+1))
	for i := 1; i <= n; i++ {
		graph[i] = row[i*n : (i+1)*n]
	}

	for i := 0; i < r; i++ {
		in, _, _ := read.ReadLine()
		input = strings.Split(string(in), " ")
		from, _ = strconv.Atoi(input[0])
		to, _ = strconv.Atoi(input[1])
		cost, _ = strconv.Atoi(input[2])
		graph[from] = append(graph[from], GraphNode{to, cost, 0})
		graph[to] = append(graph[to], GraphNode{from, cost, 0})
	}

	for i := 1; i <= n; i++ {
		Dijkstra(i, graph, items)
	}
	fmt.Print(rst)

}
func Dijkstra(cur int, graph [][]GraphNode, items []int) {
	heap := &ds.Heap[GraphNode]{make([]ds.HeapNode[GraphNode], 0), 0}
	DP := make([]int, n+1)
	for i := 0; i <= n; i++ {
		DP[i] = 1234567890
	}
	DP[cur] = 0
	heap.Insert(GraphNode{cur, 0, items[cur]})
	var NList []GraphNode
	var Nnode GraphNode
	for heap.Size != 0 {
		node := *heap.Pop()
		NList = graph[node.Val.idx]
		if node.Val.cost > m {
			continue
		}
		if rst < node.Val.itemCost {
			rst = node.Val.itemCost
		}
		for i := 0; i < len(graph[node.Val.idx]); i++ {
			Nnode = NList[i]
			if DP[Nnode.idx] > node.Val.cost+Nnode.cost {
				DP[Nnode.idx] = node.Val.cost + Nnode.cost
				heap.Insert(GraphNode{Nnode.idx, DP[Nnode.idx], node.Val.itemCost + items[Nnode.idx]})
			}
		}
	}

}
