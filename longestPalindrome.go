package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	var arr []string
	l := len(s)
	for k, v := range s {
		arr = append(arr, string(v))
		if l != 1 && k != l-1 {
			arr = append(arr, "")
		}
	}
	n := len(arr)
	b := ""
	for k, v := range arr {
		i := k
		m := n - k
		if m < i {
			i = m
		}
		a := v
		for o := 1; o <= i; o++ {
			if k+o < n && k-o >= 0 && arr[k+o] == arr[k-o] {
				if arr[k+o] != "" {
					a = arr[k-o] + a + arr[k-o]
				}
			} else {
				break
			}
		}
		if len(a) > 1000 {
			break
		}
		if len(b) < len(a) {
			b = a
		}
	}
	return b
}

func main() {
	s := "abcda"
	fmt.Println(len(s))
	a := longestPalindrome(s)
	fmt.Println(len(a))
	fmt.Println(a)
}
