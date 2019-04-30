package main

import "fmt"

func moveZeroes(nums []int) {
	//如果当前值非零就是当前值进行交换
	f := 0
	t := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			t = nums[f]
			nums[f] = nums[i]
			f++
			nums[i] = t
		}
	}
	//所有非零的都前移。剩下的补充0
	//o := 0
	//for i := 0; i < len(nums); i++ {
	//	v := nums[i]
	//	if v == 0 {
	//		continue
	//	} else {
	//		nums[o] = v
	//		o++
	//	}
	//}
	//for i := o; i < len(nums); i++ {
	//	nums[i] = 0
	//}
	//遍历两次  n*n
	//l := len(nums)
	//for i := 0; i < l-1; i++ {
	//	for i := 0; i < l-1; i++ {
	//		if nums[i] == 0 {
	//			a := nums[i]
	//			nums[i] = nums[i+1]
	//			nums[i+1] = a
	//		}
	//	}
	//}
}

func main() {
	nums := []int{
		1, 0, 1, 0, 3, 0, 12, 0,
	}
	moveZeroes(nums)
	fmt.Println(nums)
}
