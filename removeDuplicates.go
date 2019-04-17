package main

import "fmt"

func removeDuplicates(nums []int) int {
	//双指针
	i := 1
	for l := 1; l < len(nums); l++ {
		if nums[l] != nums[i-1] {
			nums[i] = nums[l]
			i++
		}
	}
	return i
	//for i := 0; i < len(nums); i++ {
	//	if i > 0 && nums[i] == nums[i-1] {
	//		i--
	//		nums = append(nums[:i], nums[i+1:]...)
	//	}
	//}
	//return len(nums)
}

func main() {
	nums := []int{
		0, 0, 1, 1, 1, 2, 2, 3, 3, 4,
	}
	res := removeDuplicates(nums)
	fmt.Println(res)
}
