package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	newS := ""
	n := len(s)
	c := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += c {
			newS += string(s[j+i])
			if i != 0 && i != numRows-1 && j+c-i < n {
				newS += string(s[j+c-i])
			}
		}
	}
	return newS
}

func main() {
	s := "LEETCODEISHIRINGh"
	a := convert(s, 3)
	fmt.Println(a)
}
