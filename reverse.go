package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var arr []int
	c := 1
	if x < 0 {
		x *= -1
		c = -1
	}
	a := 0
	for true {
		b := x % 10
		x -= b
		x /= 10
		arr = append(arr, b)
		if x < 10 {
			if x != 0 {
				arr = append(arr, x)
			}
			break
		}
	}
	l := len(arr)
	for _, v := range arr {
		l--
		a += v * int(math.Pow10(l))
	}
	a = a * c
	if a > int(math.Pow(2, 31))-1 || a < -1*int(math.Pow(2, 31)) {
		a = 0
	}
	return a
}

func main() {
	x := 1534236469
	b := reverse(x)
	fmt.Println(b)
}
