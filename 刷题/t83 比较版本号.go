package main

import "fmt"

func toint(str string) int {
	sum := 0
	for i := 0; i < len(str); i++ {
		sum = sum*10 + int(str[i]-'0')
	}
	return sum
}

func solution(version1 string, version2 string) int {
	v1Len, v2Len := len(version1), len(version2)
	tag1, tag2 := 0, 0

	for {
		// Skip dots
		for tag1 < v1Len && version1[tag1] == '.' {
			tag1++
		}
		for tag2 < v2Len && version2[tag2] == '.' {
			tag2++
		}

		// Extract numbers
		var s1, s2 string
		for tag1 < v1Len && version1[tag1] != '.' {
			s1 += string(version1[tag1])
			tag1++
		}
		for tag2 < v2Len && version2[tag2] != '.' {
			s2 += string(version2[tag2])
			tag2++
		}

		// Convert strings to integers
		num1, num2 := toint(s1), toint(s2)

		// Compare numbers
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}

		// If both have reached the end, they are equal
		if tag1 == v1Len && tag2 == v2Len {
			return 0
		}

		// If one has reached the end, but not the other, check if the remaining part is all zeros
		if tag1 == v1Len {
			// Check if the remaining part of version2 is all zeros
			for tag2 < v2Len {
				if version2[tag2] != '.' && version2[tag2] != '0' {
					return -1
				}
				tag2++
			}
			return 0
		}
		if tag2 == v2Len {
			// Check if the remaining part of version1 is all zeros
			for tag1 < v1Len {
				if version1[tag1] != '.' && version1[tag1] != '0' {
					return 1
				}
				tag1++
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
	fmt.Println(solution("1", "1.0") == 0)     // "1" and "1.0" should be equal
	fmt.Println(solution("1.0.0.0", "1") == 0) // "1.0.0.0" and "1" should be equal
}
