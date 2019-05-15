package main

import "fmt"

func lengthOfLastWord(s string) int {
	res := 0
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if res > 0 {
				break
			} else {
				res--
			}
		}
		res++
	}
	return res
}

func main() {
	s := "   a  "
	fmt.Println(lengthOfLastWord(s))
}
