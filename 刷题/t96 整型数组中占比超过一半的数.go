package main

import (
	"fmt"
	"sort"
)

func solution(array []int) int {
	// Edit your code here
	sort.Ints(array)
	return array[len(array)/2]
}

func main() {
	// Add your test cases here

	fmt.Println(solution([]int{1, 3, 8, 2, 3, 1, 3, 3, 3}) == 3)
}
