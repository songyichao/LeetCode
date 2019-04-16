package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	str := ""
	strL := len(strs)
	if strL == 0 {
		return str
	}
	n := len(strs[0])
	for _, k := range strs {
		l := len(k)
		if n > l {
			n = l
		}
	}
	if n > 0 {
		for i := 0; i < n; i++ {
			a := string(strs[0][i])
			for l := 0; l < strL; l++ {
				if string(strs[l][i]) != a {
					a = string(strs[l][i])
					break
				}
			}
			if a != string(strs[0][i]) {
				break
			} else {
				str += a
			}
		}
	}
	return str
}

func main() {
	strs := []string{
		"c", "acc", "ccc",
	}
	str := longestCommonPrefix(strs)
	fmt.Println(str)
}
