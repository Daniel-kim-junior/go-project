package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
		}
	}
}
