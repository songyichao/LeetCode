package main

import "fmt"

func isValid(s string) bool {
	res := false
	sMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	n := len(s)
	rMap := make(map[int]string)
	m := 0
	for i := 0; i < n; i++ {
		a := string(s[i])
		if _, ok := sMap[a]; ok {
			rMap[m] = sMap[a]
			m++
		} else if m > 0 && rMap[m-1] == a {
			delete(rMap, m-1)
			m--
		} else {
			return false
		}
	}
	if m == 0 {
		res = true
	}
	return res
}

func main() {
	s := "[({(())}[()])]"
	res := isValid(s)
	fmt.Println(res)
}
