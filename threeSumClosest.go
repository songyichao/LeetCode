package main

import (
	"fmt"
	"sort"
)

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -1 * num
	}
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := 0
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if i == 0 && j == 1 && k == n-1 {
				res = sum
			}
			if sum-target == 0 {
				return target
			} else if sum-target > 0 {
				if abs(res-target) >= abs(sum-target) {
					res = sum
				}
				k--
			} else {
				if abs(res-target) >= abs(sum-target) {
					res = sum
				}
				j++
			}
		}
	}
	return res
}

func main() {
	nums := []int{
		1, 2, 5, 10, 11,
	}
	target := 12
	a := threeSumClosest(nums, target)
	fmt.Println(a)
}
