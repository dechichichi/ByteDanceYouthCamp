package main

import "fmt"

var t []int

func t_init(A int) {
	t = make([]int, A+1)
	t[0] = 0
	t[1] = 1
	t[2] = 2
	for i := 3; i < A; i++ {
		t[i] = t[i-1] + t[i-2]
	}
}

func solution(A int) int {
	// Edit your code here
	return t[A]
}

func main() {
	// Add your test cases here
	t_init(75)
	fmt.Println(solution(5) == 8)
	fmt.Println(solution(1) == 1)
	fmt.Println(solution(15) == 987)
	fmt.Println(solution(50) == 20365011074)
}
