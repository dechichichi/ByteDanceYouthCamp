package main

import (
	"fmt"
	"strconv"
)

// tick 函数用于判断一个整数是否为完美整数
func tick1(n int, len int) int {
	// 首先将数字转换为字符串
	strN := strconv.Itoa(n)
	// 记录第一个数字
	firstDigit := int(strN[0] - '0')
	tk := 0
	for ; len > 0; len-- {
		tk *= 10
		tk += firstDigit
	}
	if n <= tk {
		return 10 - firstDigit
	} else {
		return 9 - firstDigit
	}

}
func tick2(n int, len int) int {
	// 首先将数字转换为字符串
	strN := strconv.Itoa(n)
	// 记录第一个数字
	firstDigit := int(strN[0] - '0')
	tk := 0
	for ; len > 0; len-- {
		tk *= 10
		tk += int(firstDigit)
	}
	// 10 11
	if n >= tk {
		return firstDigit
	} else {
		return firstDigit-1
	}
}

// solution 函数计算区间 [x, y] 中完美整数的数量
func solution(x, y int) int {
	// 计算 x 和 y 的数字长度
	len1 := 0
	len2 := 0
	tmp1 := x
	tmp2 := y
	for tmp1 > 0 {
		tmp1 /= 10
		len1++
	}
	for tmp2 > 0 {
		tmp2 /= 10
		len2++
	}
	ans := 0
	// 计算中间长度的区间中完美整数的数量
	ans += (len2 - len1 - 1) * 9
	// 判断 x 和 y 是否为完美整数并进行计数
	ans += tick1(x, len1)
	ans += tick2(y, len2)
	return ans
}

func main() {
	// 您可以在这里添加测试用例
	fmt.Println(solution(1, 10)==9)
	fmt.Println(solution(2, 22)==10)
}

// 2 22
