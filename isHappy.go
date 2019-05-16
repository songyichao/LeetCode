package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]int)
	for n != 1 {
		sum := 0
		for n > 0 {
			i := n % 10
			sum += i * i
			n /= 10
		}
		if m[sum] == 0 {
			m[sum]++
			n = sum
		} else {
			return false
		}
	}
	return true
}

func main() {
	n := 12
	fmt.Println(isHappy(n))
}
