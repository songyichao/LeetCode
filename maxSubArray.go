package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	res := math.MaxInt64
	sum := 0
	for _, v := range nums {
		sum = int(math.Max(float64(sum+v), float64(v)))
		res = int(math.Max(float64(res), float64(sum)))
	}

	return res
}

func main() {
	nums := []int{
		-2, 1, -3, 4, -1, 2, 1, -5, 4,
	}
	fmt.Println(maxSubArray(nums))
}
