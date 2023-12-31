package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	nums := []int{5, 1, 2, 4, 5, 6, 7, 2, 3, 4, 5, 7, 8, 0}
	rst := make([]int, len(nums))
	records := [11]int{}

	students := []struct {
		Name   string
		Height float64
	}{
		{
			Name:   "Keyle",
			Height: 173.4,
		},
		{
			Name:   "Ken",
			Height: 164.5,
		},
		{
			Name:   "Ryu",
			Height: 178.8,
		},
		{
			Name:   "Honda",
			Height: 154.2,
		},
		{
			Name:   "Hwarang",
			Height: 188.8,
		},
		{
			Name:   "Lebron",
			Height: 209.9,
		},
		{
			Name:   "Hodong",
			Height: 197.7,
		},
		{
			Name:   "Tom",
			Height: 164.8,
		},
	}
	var HeightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		HeightMap[idx] = append(HeightMap[idx], students[i].Name)
	}

	for i := 0; i < len(nums); i++ {
		records[nums[i]]++
	}

	for i := 1; i < len(records); i++ {
		records[i] += records[i-1]
	}

	for i := 0; i < len(nums); i++ {
		rst[records[nums[i]]-1] = nums[i]
		records[nums[i]]--
	}

	fmt.Println(rst)

	strArr := []string{"a", "a", "b", "c", "d", "k", "k", "k"}
	lenArr := []int{100, 200, 170, 150, 174, 156, 137, 177, 178, 191}
	findFrequentAlphabetFromArray(&strArr)
	// fmt.Print(c)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)
	from, _ := strconv.Atoi(words[0])
	to, _ := strconv.Atoi(words[1])
	r := findSpecificLenStudents(from, to, &lenArr)
	fmt.Print(*r)
}

func findFrequentAlphabetFromArray(arr *[]string) string {
	count := make([]int, 26)
	maxValue := 0
	frequent := ""
	for i := 0; i < len(*arr); i++ {
		n, _ := utf8.DecodeRuneInString((*arr)[i])
		count[n-97]++
		if maxValue < count[n-97] {
			maxValue = count[n-97]
			frequent = (*arr)[i]
		}
	}
	return frequent
}

func findSpecificLenStudents(minLen int, maxLen int, arr *[]int) *[]int {
	rst := make([]int, 0, len(*arr))
	for _, Val := range *arr {
		if maxLen >= Val && minLen <= Val {
			rst = append(rst, Val)
		}
	}
	return &rst
}
