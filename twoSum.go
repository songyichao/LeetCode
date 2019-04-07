package app

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		l := target - nums[i]
		v, e := m[l]
		if e {
			return []int{i, v}
		}
		m[nums[i]] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(nums)
	res := twoSum(nums, 9)
	fmt.Println(res)
}
