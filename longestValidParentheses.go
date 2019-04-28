package main

import "fmt"

func longestValidParentheses(s string) int {
	var  num =  0
	n := len(s)
	if n < 2 {
		return num
	}
	var rArr []int
	m := 0
	for i := 0; i < n; i++ {
		a := string(s[i])
		if a == "(" {
			rArr = append(rArr, i)
		} else {
			if len(rArr) == 0 {
				m = i + 1
				continue
			}
			rArr = rArr[:len(rArr)-1]
			if len(rArr) == 0 {
				if i-m+1 > num {
					num = i - m + 1
				}
			} else {
				if i-rArr[len(rArr)-1]-1 > num {
					num = i - rArr[len(rArr)-1]
				}
			}
		}
	}
	return num
}
func main() {
	s := "(()()("
	fmt.Println(longestValidParentheses(s))
}
