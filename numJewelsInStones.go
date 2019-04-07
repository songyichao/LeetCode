package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	mapJ := make(map[string]int)
	for _, v := range J {
		mapJ[string(v)] = 1
	}
	n := 0
	for _, v := range S {
		if _, ok := mapJ[string(v)]; ok {
			n++
		}
	}
	return n
}

func main() {
	J := "aA"
	S := "aAAbbbb"
	n := numJewelsInStones(J, S)
	fmt.Println(n)
}
