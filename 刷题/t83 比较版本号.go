package main

import "fmt"

func toint(str string) int {
	sum := 0
	for i := len(str) - 1; i >= 0; i-- {
		sum *= 10
		sum += int(str[i] - '0')
	}
	return sum
}

func solution(version1 string, version2 string) int {
	tag1 := 0
	tag2 := 0
	s1 := ""
	s2 := ""
	for {
		// Find the next segment in version1
		for tag1 < len(version1) && version1[tag1] != '.' {
			s1 += string(version1[tag1])
			tag1++
		}
		// Find the next segment in version2
		for tag2 < len(version2) && version2[tag2] != '.' {
			s2 += string(version2[tag2])
			tag2++
		}
		// Compare the segments
		if len(s1) > 0 && len(s2) > 0 {
			if i := toint(s1); i > toint(s2) {
				return 1
			} else if i < toint(s2) {
				return -1
			}
		}
		// Reset for the next segment
		s1 = ""
		s2 = ""
		tag1++
		tag2++
		// If we've reached the end of both versions, they are equal
		if tag1 == len(version1) && tag2 == len(version2) {
			return 0
		}
		// If one version is a prefix of the other, the shorter one is smaller
		if tag1 == len(version1) || tag2 == len(version2) {
			// Check if the remaining part of the longer version contains a non-zero number
			if tag1 == len(version1) {
				for tag2 < len(version2) && version2[tag2] == '.' {
					tag2++
				}
				if tag2 < len(version2) {
					return -1
				}
			} else {
				for tag1 < len(version1) && version1[tag1] == '.' {
					tag1++
				}
				if tag1 < len(version1) {
					return 1
				}
			}
			return 0
		}
	}
}

func main() {
	// Add your test cases here
	fmt.Println(solution("0.1", "1.1") == -1)
	fmt.Println(solution("1.0.1", "1") == 1)
	fmt.Println(solution("7.5.2.4", "7.5.3") == -1)
	fmt.Println(solution("1.0", "1.0.0") == 0)
}
