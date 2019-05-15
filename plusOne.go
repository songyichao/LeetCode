package main

import "fmt"

func plusOne(digits []int) []int {
	var tmp []int
	l := len(digits)
	o := 0
	for i := l - 1; i >= 0; i-- {
		s := digits[i] + o
		if i == l-1 {
			s++
		}
		q := s % 10
		tmp = append(tmp, q)
		o = (s - q) / 10
	}
	if o != 0 {
		tmp = append(tmp, o)
	}
	t := len(tmp)
	res := make([]int, t)
	for i := t - 1; i >= 0; i-- {
		res[t-1-i] = tmp[i]
	}
	return res
}
func main() {
	digits := []int{
		9, 9, 9,
	}
	fmt.Println(plusOne(digits))
}
