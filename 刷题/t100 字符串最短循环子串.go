package main

import (
	"fmt"
)

// 判断一个子串是否为神奇数列
func good(substr string) bool {
	// 检查子串是否满足 1 和 0 没有重复跟随并且至少由 3 个数组成
	if len(substr) < 3 {
		return false
	}
	prev := rune(substr[0])
	for _, curr := range substr[1:] {
		if curr == prev {
			return false
		}
		prev = curr
	}
	return true
}

// 比较两个整数，返回较大的那个
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution(inp string) string {
	// 初始化滑动窗口的左右边界
	left, right := 0, 0
	maxLength := 0
	maxSubstr := ""
	for right < len(inp) {
		// 扩展窗口
		right++
		if !good(inp[left:right]) {
			maxLength = max(maxLength, right-left-1)
			if len(maxSubstr) < maxLength {
				maxSubstr = inp[left : left+maxLength]
			}
			left = right - (right-left) + 1
		}
	}
	maxLength = max(maxLength, right-left)
	if len(maxSubstr) < maxLength {
		maxSubstr = inp[left : left+maxLength]
	}
	return maxSubstr
}

func main() {
	// 测试用例
	fmt.Println(solution("0101011101"))
}