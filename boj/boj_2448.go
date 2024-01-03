package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var starMap [][]rune
var row []rune

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, _, _ := reader.ReadLine()
	star, _ := strconv.Atoi(string(n))

	rst := drawStar(star)
	for i := 0; i < len(rst); i++ {
		fmt.Println(rst[i])
	}

}

func drawStar(n int) []string {
	if n == 3 {
		return []string{"  *  ", " * * ", "*****"}
	}
	half := drawStar(n / 2)
	blank := strings.Repeat(" ", n/2)
	bBlank := " "
	roof := half
	bottom := make([]string, 0, 0)
	for i := range half {
		roof = append(roof, blank+half[i]+blank)
		bottom = append(bottom, half[i]+bBlank+half[i])
	}

	roof = append(roof, bottom...)

	return roof[n/2:]

}
