package main

import "fmt"

func search(nums []int, target int) int {
	//二分法查找
	n := -1
	l := len(nums) - 1
	i := 0
	if len(nums) == 0 || (target < nums[0] && target > nums[l]) {
		return -1
	}
	if target == nums[0] {
		return 0
	}
	if target == nums[l] {
		return l
	}
	for i+1 < l {
		mid := (l + i) / 2
		if nums[mid] == target {
			n = mid
			break
		}
		if nums[i] < nums[l] {
			if nums[mid] < target {
				i = mid + 1
			} else {
				l = mid - 1
			}
		}
		//左侧有序
		if nums[i] < nums[mid] {
			if nums[i] < target && target < nums[mid] {
				l = mid
			} else {
				i = mid
			}
			continue
		} else {
			if nums[mid] < target && target < nums[l] {
				i = mid
			} else {
				l = mid
			}
		}
	}
	if i <= len(nums)-1 && nums[i] == target {
		n = i
	}
	return n
	//遍历一次查找
	//n := -1
	//for k, v := range nums {
	//	if v == target {
	//		n = k
	//		break
	//	}
	//}
	//return n
}

func main() {
	nums := []int{
		 1, 3,
	}
	target := 2
	fmt.Println(search(nums, target))
}
