package main

import "fmt"

func bs(nums []int, res []int, target, start, end int) {
	if start > end {
		return
	}
	mid := (start + end) / 2
	if target == nums[mid] {
		if res[0] == -1 || res[0] > mid {
			res[0] = mid
		}
		if res[1] == -1 || res[1] < mid {
			res[1] = mid
		}
		bs(nums, res, target, mid+1, end)
		bs(nums, res, target, start, mid-1)
	} else if target > nums[mid] {
		bs(nums, res, target, mid+1, end)
	} else {
		bs(nums, res, target, start, mid-1)
	}
}
func searchRange(nums []int, target int) []int {
	arr := []int{
		-1, -1,
	}
	i := 0
	l := len(nums) - 1
	bs(nums, arr, target, i, l)
	return arr
}

func main() {
	nums := []int{
		5, 7, 7, 8, 8, 8, 8,
	}
	target := 8
	fmt.Println(searchRange(nums, target))
}
