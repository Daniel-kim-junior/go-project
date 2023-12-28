package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node[T any] struct {
	x T
	y T
}

type Item interface {
}

type Queue struct {
	items []Item
	len   int
}

func (q *Queue) enqueue(item Item) {
	q.items = append(q.items, item)
	q.len++
}

func (q Queue) dequeue() Item {
	if q.isEmpty() {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	q.len--
	return item
}

func (q Queue) isEmpty() bool {
	if len(q.items) == 0 {
		return true
	}
	return false
}

var x int
var y int
var rst int = 0

func main() {
	in := bufio.NewReader(os.Stdin)
	n, _, _ := in.ReadLine()

	N, _ := strconv.Atoi(string(n))

	miniMap := make([][]int, N)
	row := make([]int, N*N)
	/**
		2차원 배열 생성시 Cache locality를 증가시켜 Frequently Cache hit를 위해
		1차원 배열의 슬라이싱의 메모리를 분배
	**/
	for i := 0; i < N; i++ {
		miniMap[i] = row[i*N : (i+1)*N]
	}

	for i := 0; i < N; i++ {
		str, _, _ := in.ReadLine()
		strArr := strings.Split(string(str), " ")
		for j := 0; j < N; j++ {
			s, _ := strconv.Atoi(strArr[j])
			miniMap[i][j] = s
			if s == 9 {
				x = i
				y = j
			}
		}
	}
	BFS(miniMap, N)
	fmt.Print(rst)
}

func BFS(miniMap [][]int, size int) {
	shark := &Node[int]{x, y}
	queue := &Queue{}
	queue.enqueue(shark)

}
