package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var n int
var m int
var r int

func main() {
	read := bufio.NewReader(os.Stdin)
	line, _, _ := read.ReadLine()
	input := strings.Split(string(line), " ")
	n, _ = strconv.Atoi(input[0])
	m, _ = strconv.Atoi(input[1])
	r, _ = strconv.Atoi(input[2])
	items := make([]int, n)
	line, _, _ = read.ReadLine()
	input = strings.Split(string(line), " ")
	var item int
	for i := 0; i < n; i++ {
		item, _ = strconv.Atoi(input[i])
		items[i] = item
	}
	var from int
	var to int
	var cost int
	graph := make([][]int, n+1)
	row := make([]int, (n+1)*(n+1))
	for i := 1; i <= n; i++ {
		graph[i] = row[i*n : (i+1)*n]
	}

	for i := 0; i < r; i++ {
		in, _, _ := read.ReadLine()
		input = strings.Split(string(in), " ")
		from, _ = strconv.Atoi(input[0])
		to, _ = strconv.Atoi(input[1])
		cost, _ = strconv.Atoi(input[2])
		graph[from] = []int{to, cost}

	}

}
