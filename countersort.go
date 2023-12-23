package main

import "fmt"

func main() {
	nums := []int{5, 1, 2, 4, 5, 6, 7, 2, 3, 4, 5, 7, 8, 0}
	rst := make([]int, len(nums))
	records := [11]int{}

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
}
