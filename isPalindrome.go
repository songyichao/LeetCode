package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	a := false
	if x >= 0 && x < 10 {
		a = true
	}
	if x >= 10 {
		var arr []int
		//b := 0
		b := x
		for b >= 10 {
			arr = append(arr, b%10)
			b = b / 10
			if b < 10 {
				arr = append(arr, b)
			}
		}
		c := 0
		l := len(arr) - 1
		for k, v := range arr {
			c += v * int(math.Pow10(l-k))
		}
		if c == x {
			a = true
		}
	}
	return a
}

func main() {
	x := 123
	d := isPalindrome(x)
	fmt.Println(d)
}
