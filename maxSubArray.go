package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	sum, max := 0, nums[0]

	for _, v := range nums {
		if sum < 1 {
			sum = 0
		}
		sum += v
		if sum > max {
			max = sum
		}
	}
	return max
	//一个比较慢的版本
	//res := math.MinInt64
	//sum := 0
	//for _, v := range nums {
	//	sum = int(math.Max(float64(sum+v), float64(v)))
	//	res = int(math.Max(float64(res), float64(sum)))
	//}
	//
	//return res
}

func main() {
	nums := []int{
		1,
	}
	fmt.Println(maxSubArray(nums))
}
