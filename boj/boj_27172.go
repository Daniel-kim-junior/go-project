package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read N
	in, _, _ := reader.ReadLine()
	N, _ := strconv.Atoi(string(in))

	// Read player input
	in, _, _ = reader.ReadLine()
	playerStr := strings.Split(string(in), " ")

	// Initialize arrays
	SIZE := 1000001
	player := make([]int, N)
	card := make([]bool, SIZE)
	score := make([]int, SIZE)

	// Convert player input to integers and populate arrays
	for i, s := range playerStr {
		player[i], _ = strconv.Atoi(s)
		card[player[i]] = true
	}

	// Process scores
	for _, i := range player {
		for j := i * 2; j < SIZE; j += i {
			if card[j] {
				score[i]++
				score[j]--
			}
		}
	}

	// Build the output string
	var sb strings.Builder
	for _, num := range player {
		sb.WriteString(strconv.Itoa(score[num]) + " ")
	}

	// Print the result
	fmt.Println(sb.String())
}
