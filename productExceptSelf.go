package main

import "fmt"

func productExceptSelf(nums []int) []int {
	//
	n := len(nums)
	var left, right = 1, 1
	res := make([]int, n)
	for i := 0; i < n; i++ {
		//if left == 1 && res[i] == 0 {
		//	res[i] = 1
		//}
		//if right == 1 && res[n-i-1] == 0 {
		//	res[n-i-1] = 1
		//}
		res[i] *= left
		res = append(res, left)
		res[n-i-1] *= right
		left *= nums[i]
		right *= nums[n-i-1]
	}
	return res
	//从左到右相乘  后再从右到左乘  只能说佩服 但是还是超时了
	//n := len(nums)
	//res := make([]int, n)
	//var left, right = 1, 1
	//for i := 0; i < n; i++ {
	//	res[i] = left
	//	left *= nums[i]
	//}
	//for l := n - 1; l >= 0; l-- {
	//	res[l] *= right
	//	right *= nums[l]
	//}
	//return res

	//双指针超时了
	//var res []int
	//n := len(nums)
	//for k, _ := range nums {
	//	s := 1
	//	l := 0
	//	m := n - 1
	//	for l <= m {
	//		if (m != k && nums[m] == 0) || (l != k && nums[l] == 0) {
	//			s = 0
	//			break
	//		}
	//		if m != k {
	//			s *= nums[m]
	//		}
	//		if l != k && l != m {
	//			s *= nums[l]
	//		}
	//		l++
	//		m--
	//	}
	//	res = append(res, s)
	//}
	//return res
}

func main() {
	nums := []int{
		1, 2, 3, 4,
	}
	fmt.Println(productExceptSelf(nums))
}
