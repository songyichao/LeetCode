package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	res := math.MinInt64
	sum := 0
	for _, v := range nums {
	sum = int(math.Max(float64(sum+v), float64(v)))
	res = int(math.Max(float64(res), float64(sum)))
}

	return res
}

	func main() {
	nums := []int{
		1,
	}
	fmt.Println(maxSubArray(nums))
}
