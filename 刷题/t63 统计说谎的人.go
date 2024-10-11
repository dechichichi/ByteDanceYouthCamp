package main

import (
	"fmt"
	"sort"
)

func solution(A []int) int {
	// Edit your code here
	sort.Ints(A)
	ans := len(A) / 2
	mid := len(A)/2 + len(A)&1
	midnum := A[ans]
	for A[mid-1] == midnum {
		ans++
		if ans == len(A) {
			break
		}
		mid--
	}
	return ans
}

//1 30 30 30

// 1 2 3
func main() {
	// Add your test cases here
	fmt.Println(solution([]int{100, 100, 100}) == 3)
	fmt.Println(solution([]int{2, 1, 3}) == 2)
	fmt.Println(solution([]int{30, 1, 30, 30}) == 3)
	fmt.Println(solution([]int{19, 27, 73, 55, 88}) == 3)
	fmt.Println(solution([]int{19, 27, 73, 55, 88, 88, 2, 17, 22}) == 5)
}
