package main

import "fmt"

func strStr(haystack string, needle string) int {
	num := -1
	l := len(haystack)
	m := len(needle)
	n := 0
	a := -1
	for i := 0; i < l && n < m; i++ {
		if string(haystack[i]) == string(needle[n]) {
			if a == -1 {
				a = i
			}
			n++
		} else if a != -1 && m != n && string(haystack[i]) != string(needle[n]) {
			i = a
			a = -1
			n = 0
		}
	}
	if m == n {
		num = a
	}
	if m == 0 {
		num = 0
	}
	return num
}

func main() {
	haystack := "aaaaa"
	needle := "bba"
	a := strStr(haystack, needle)
	fmt.Println(a)
}
