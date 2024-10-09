package main

import (
	"fmt"
)

func good(substr string) bool {
	// 首先检查子串长度，如果小于 3 则直接返回 false，因为题目要求至少 3 个数
	prev := rune(substr[0])
	// 遍历子串，检查相邻字符是否相同
	for _, curr := range substr[1:] {
		if curr == prev {
			return false
		}
		prev = curr
	}
	return true
}

// 优化后的 solution 函数
func solution(inp string) string {
	// 初始化滑动窗口的左右边界
	left, right := 0, 0
	maxLength := 0
	maxSubstr := ""
	// 当右边界未超出输入字符串长度时
	for right < len(inp) {
		// 如果当前窗口满足条件或者窗口长度小于 2，扩展右边界
		if right-left < 3 || good(inp[left:right]) {
			right++
		} else {
			// 如果当前窗口长度大于最大长度，更新最大长度和最大子串
			if right-left > maxLength {
				maxLength = right - left
				maxSubstr = inp[left : right-1]
			}
			// 移动左边界
			left++
		}
	}
	// 检查最后一次窗口
	if right-left > maxLength {
		maxLength = right - left
		maxSubstr = inp[left:]
	}
	return maxSubstr
}

func main() {
	// 测试用例
	fmt.Println(solution("0101011101") == "010101")
}
