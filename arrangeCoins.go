package main

import "fmt"

func arrangeCoins(n int) int {
	//递减就完事了
	res := 0
	for i := 1; n >= i; i++ {
		n -= i
		res = i
	}
	return res
}

func main() {
	n := 1
	fmt.Println(arrangeCoins(n))
}
