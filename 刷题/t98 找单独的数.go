package main

import "fmt"

func solution(inp []int) int {

  // Edit your code here

  tag := make(map[int]int)

  for _,i := range inp {

​    tag[i]++

  }

  for _,i := range inp {

​    if tag[i] == 1 {

​      return i

​    }

  }

  return 0

}

func main() {

  // Add your test cases here

  fmt.Println(solution([]int{1, 1, 2, 2, 3, 3, 4, 5, 5}) == 4)

  fmt.Println(solution([]int{0, 1, 0, 1, 2}) == 2)

}