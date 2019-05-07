package main

import "fmt"

func canJump(nums []int) bool {
	//遍历一次，每次获取可到达的最大点
	n := len(nums)
	if n == 0 {
		return true
	}
	max := 0
	for i := 0; i < n; i++ {
		if i > max {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return true
	//从后往前查看
	//res := true
	//n := len(nums)
	//for i := n - 2; i >= 0; i-- {
	//	if nums[i] == 0 {
	//		can := -1
	//		for l := i - 1; l >= 0; l-- {
	//			if i-l < nums[l] {
	//				can = l
	//				break
	//			}
	//		}
	//		if can == -1 {
	//			res = false
	//			break
	//		}
	//		i = can
	//	}
	//}
	//return res
}
func main() {
	nums := []int{
		2, 0, 0,
	}
	fmt.Println(canJump(nums))
}
