package main

import "fmt"

func subarraySum(nums []int, k int) int {
	//使用hash的方法
	res, sum := 0, 0
	numMap := make(map[int]int, len(nums))
	numMap[0] = 1
	for _, v := range nums {
		sum += v
		res += numMap[sum-k]
		numMap[sum]++
	}
	return res
	//两次遍历
	//res := 0
	//n := len(nums)
	//left := 0
	//l := 0
	//for i := 0; i < n; {
	//	left += nums[i]
	//	if left == k {
	//		res++
	//	}
	//	if l < n && i == n-1 {
	//		i = l
	//		left = 0
	//		l++
	//	}
	//	i++
	//}
	//return res
}

func main() {
	nums := []int{
		-1, -1, 1,
	}
	k := 0
	fmt.Println(subarraySum(nums, k))
}
