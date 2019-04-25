package main

import (
	"fmt"
	"sort"
)

func nextPermutation(nums []int) {
	//newNums := nums
	//sort.Ints(newNums)
	//fmt.Println(nums)
	//fmt.Println(newNums)
	//max := 0
	n := len(nums)
	l := 0
	m := 0
	for i := n - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			if i == n-1 {
				a := nums[i]
				nums[i] = nums[i-1]
				nums[i-1] = a
			} else {
				for m = i; m < n-1; m++ {
					if nums[m] >= nums[m+1] && nums[m] > nums[i-1] && (nums[m+1] <= nums[i-1] || m+1 == n-1) {
						if nums[m+1] <= nums[i-1] {
							a := nums[i-1]
							nums[i-1] = nums[m]
							nums[m] = a
							break
						}else if  m+1 == n-1 {
							a := nums[i-1]
							nums[i-1] = nums[m+1]
							nums[m+1] = a
							break
						}
					}
				}
				newNums := nums[i:]
				sort.Ints(newNums)
				nums = append(nums[:i], newNums...)
			}
			l++
			break
		}
	}
	if l == 0 {
		sort.Ints(nums)
	}
}

func main() {
	nums := []int{
		2,3,1,3,3,
	}
	nextPermutation(nums)
	fmt.Println(nums)
}
