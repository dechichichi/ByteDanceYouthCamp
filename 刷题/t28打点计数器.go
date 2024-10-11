package main

import (
	"fmt"
	"sort"
)

func solution(inputArray [][]int) int {
	// 使用sort.Slice进行排序，比较函数根据每个子数组的第二个元素来比较
	sort.Slice(inputArray, func(i, j int) bool {
		return inputArray[i][1] < inputArray[j][1]
	})
	tag := 0
	sum := 0
	for _, i := range inputArray {
		if i[1] >= tag {
			sum += i[1] - max(i[0], tag)
			tag = i[1]
		}
	}
	return sum
}

func main() {
	// 你可以在这里添加更多的测试用例
	testArray1 := [][]int{{1, 4}, {7, 10}, {3, 5}}
	testArray2 := [][]int{{1, 2}, {6, 10}, {11, 15}}

	fmt.Println(solution(testArray1) == 7)
	fmt.Println(solution(testArray2) == 9)
}
