package main

import "fmt"

func mySqrt(x int) int {
	mMap := make(map[int]int)
	res := 0
	s := 1
	e := x
	n := 0
	for n < x {
		m := (s + e) / 2
		n = m * m
		if mMap[m] > 0 {
			break
		}
		mMap[m]++
		if n == x {
			res = m
			break
		} else if n < x {
			res = m
			s = m
		} else {
			e = m
			n = 0
		}
	}
	return res
}

func main() {
	x := 1
	fmt.Println(mySqrt(x))
}
