package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var newNums [][]int
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			k := j + 1
			l := n - 1
			for k < l {
				if nums[i]+nums[j]+nums[k]+nums[l] == target {
					newNums = append(newNums, []int{nums[i], nums[j], nums[k], nums[l],})
					k++
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for j < n-2 && nums[j] == nums[j+1] {
						j++
					}
					for i < n-3 && nums[i] == nums[i+1] {
						i++
					}
				} else if nums[i]+nums[j]+nums[k]+nums[l] > target {
					l--
				} else {
					k++
				}
			}
		}
	}
	return newNums
}
func main() {
	nums := []int{
		-3, -2, -1, 0, 0, 1, 2, 3,
	}
	target := 0
	newNums := fourSum(nums, target)
	fmt.Println(newNums)
}
