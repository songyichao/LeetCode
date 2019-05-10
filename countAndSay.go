package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	last := "1"
	next := ""
	for i := 0; i < n; i++ {
		str := ""
		cnt := 0
		for _, v := range last {
			sv := string(v)
			if cnt == 0 && str == "" {
				str = sv
				cnt++
			} else if str == sv {
				cnt++
			} else {
				next += strconv.Itoa(cnt) + str
				str = sv
				cnt = 1
			}
		}
		if cnt != 0 && str != "" {
			next += strconv.Itoa(cnt) + str
		}
		res = last
		last = next
		next = ""
	}
	return res
}
func main() {
	fmt.Println(countAndSay(5))
}
