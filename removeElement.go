package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	i := 0
	for l := 0; l < len(nums); l++ {
		if nums[l] != val {
			nums[i] = nums[l]
			i++
		}
	}
	return i
}

func main() {
	nums := []int{
		3, 2, 2, 3,
	}
	val := 2
	res := removeElement(nums, val)
	fmt.Println(res)
}
