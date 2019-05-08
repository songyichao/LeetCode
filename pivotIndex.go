package main

import "fmt"

func pivotIndex(nums []int) int {
	//一个更快的版本
	res := -1
	n := len(nums)
	var sum, left, right = 0, 0, 0
	for _, v := range nums {
		sum += v
	}
	for i := 0; i < n; i++ {
		right = sum - left - nums[i]
		if left == right {
			res = i
			break
		}
		left += nums[i]
	}
	return res

	//类似不包含本身的乘积
	//res := -1
	//var left, right = 0, 0
	//n := len(nums)
	//newNums := make([]int, n)
	//for i := 0; i < n; i++ {
	//	newNums[i] += left
	//	newNums[n-i-1] -= right
	//	left += nums[i]
	//	right += nums[n-i-1]
	//}
	//for k, v := range newNums {
	//	if v == 0 {
	//		res = k
	//		break
	//	}
	//}
	//fmt.Println(newNums)
	//return res
}

func main() {
	nums := []int{
		-1,-1,0,0,-1,-1,
	}
	fmt.Println(pivotIndex(nums))
}
