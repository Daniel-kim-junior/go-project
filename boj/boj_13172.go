package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const mod int = 1000000007

func main() {
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	M := string(bytes)
	dices, _ := strconv.Atoi(M)
	var data []string
	var N int
	var S int
	var rst int = 0
	for i := 1; i <= dices; i++ {
		bytes, _, _ = reader.ReadLine()
		data = strings.Split(string(bytes), " ")
		N, _ = strconv.Atoi(data[0])
		S, _ = strconv.Atoi(data[1])
		t := calModRev(N, mod-2)
		rst = (rst + (t*S)%mod) % mod
	}
	fmt.Print(rst)

}

func calModRev(n int, r int) int {
	if r == 0 {
		return 1
	}
	var rst int
	if r%2 == 0 {
		rst = calModRev(n, r>>1) % mod
		return (rst * rst) % mod
	} else {
		rst = calModRev(n, (r-1)>>1) % mod
		return (((rst * rst) % mod) * n) % mod
	}
}
