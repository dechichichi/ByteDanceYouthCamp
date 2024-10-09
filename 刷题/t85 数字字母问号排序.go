package main

import (

  "fmt"

  "sort"

  "unicode"

)

func solution(inp string) string {

  // 初始化切片

  nums := make([]rune, 0)

  lets := make([]rune, 0)

  questionMarks := []rune{}

  // 用于记录数字和字母的索引

  numIndex := make(map[int]int)

  letIndex := make(map[int]int)

  // 遍历输入字符串，分类存储数字和字母，并记录索引

  for index, i := range inp {

​    if unicode.IsDigit(i) {

​      nums = append(nums, i)

​      numIndex[index] = len(nums) - 1

​    } else if unicode.IsLetter(i) {

​      lets = append(lets, i)

​      letIndex[index] = len(lets) - 1

​    } else if i == '?' {

​      questionMarks = append(questionMarks, i)

​    }

  }

  // 对数字和字母进行排序

  sort.SliceStable(nums, func(i, j int) bool { return nums[i] > nums[j] }) // 降序排序

  sort.SliceStable(lets, func(i, j int) bool { return lets[i] < lets[j] }) // 字典序排序

  // 构建结果字符串

  result := make([]rune, len(inp))

  for index, i := range inp {

​    if i == '?' {

​      result[index] = i // 问号位置不变

​    } else if idx, exists := numIndex[index]; exists {

​      // 数字位置不变，但替换为排序后的数字

​      result[index] = nums[idx]

​    } else if idx, exists := letIndex[index]; exists {

​      // 字母位置不变，但替换为排序后的字母

​      result[index] = lets[idx]

​    }

  }

  return string(result)

}

func main() {

  fmt.Println(solution("12A?zc") == "21A?cz")

}