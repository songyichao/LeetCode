package main

import "fmt"

func arrangeCoins(n int) int {
	//二分查找
	if n == 0 {
		return 0
	}
	if n <0 {
		return -1
	}
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2
		sum := mid * (1+mid)/2
		if sum > n {
			right = mid-1
		}
		if sum < n {
			left = mid +1
		}
		if sum == n {
			return mid
		}
	}

	return left-1
	//递减就完事了
	//res := 0
	//for i := 1; n >= i; i++ {
	//	n -= i
	//	res = i
	//}
	//return res
}

func main() {
	n := 1
	fmt.Println(arrangeCoins(n))
}
