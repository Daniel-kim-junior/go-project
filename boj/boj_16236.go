package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	ds "github.com/Daniel-kim-junior/go-project/godatasturct"
)

type NodeList []ds.Node

func (n NodeList) Len() int {
	return len(n)
}

func (n NodeList) Less(i, j int) bool {
	if n[i].Data.([]int)[2] == n[j].Data.([]int)[2] {
		if n[i].Data.([]int)[0] == n[j].Data.([]int)[0] {
			return n[i].Data.([]int)[1] < n[j].Data.([]int)[1]
		}
		return n[i].Data.([]int)[0] < n[j].Data.([]int)[0]
	}

	return n[i].Data.([]int)[2] < n[j].Data.([]int)[2]
}
func (n NodeList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
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
	fmt.Println(rst)
}

func BFS(miniMap [][]int, size int) {
	shark := []int{x, y}
	hashMap := make(map[int]int)
	dir := [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
	queue := &ds.DoubleLinkedList{}
	queue.PushFront(shark)
	sharkSize := 2
	cur := queue.RootNode
	var xy []int
	var visited [][]bool
	var row []bool
	for ; cur != nil; cur = cur.NextNode {
		xy = NtoShark(cur)
		willEat := 0
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if miniMap[i][j] != 0 && miniMap[i][j] != 9 && sharkSize > miniMap[i][j] {
					willEat++
				}
			}
		}

		if willEat == 0 {
			return
		} else {
			visited = make([][]bool, size)
			row = make([]bool, size*size)
			for i := 0; i < size; i++ {
				visited[i] = row[i*size : (i+1)*size]
			}
			tmpList := NodeList{}
			innerQueue := &ds.DoubleLinkedList{}
			innerQueue.PushBack([]int{xy[0], xy[1], 0})
			visited[xy[0]][xy[1]] = true
			innerCur := innerQueue.RootNode

			for ; innerCur != nil; innerCur = innerCur.NextNode {
				l := NtoShark(innerCur)
				for k := 0; k < 4; k++ {
					lx := l[0] + dir[k][0]
					ly := l[1] + dir[k][1]
					if lx < 0 || ly < 0 || lx >= size || ly >= size || miniMap[lx][ly] > sharkSize || visited[lx][ly] {
						continue
					}
					visited[lx][ly] = true
					if miniMap[lx][ly] != 0 && miniMap[lx][ly] < sharkSize {
						tmpList = append(tmpList, ds.Node{[]int{lx, ly, l[2] + 1}, nil, nil})
					}
					innerQueue.PushBack([]int{lx, ly, l[2] + 1})
				}
			}
			if tmpList.Len() == 0 {
				return
			}
			sort.Sort(tmpList)
			id := tmpList[0].Data.([]int)
			miniMap[xy[0]][xy[1]] = 0
			miniMap[id[0]][id[1]] = 9
			queue.PushBack([]int{id[0], id[1]})
			rst += id[2]
		}
		hashMap[sharkSize] = hashMap[sharkSize] + 1
		if sharkSize == hashMap[sharkSize] {
			sharkSize++
		}
	}

}

func NtoShark(node *ds.Node) []int {
	return node.Data.([]int)
}
