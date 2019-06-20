package main

import "fmt"

func containsDuplicatecontain(nums []int) bool {
	var res bool
	numsMap := make(map[int]int)
	res = false
	for _, v := range nums {
		numsMap[v]++
		if numsMap[v] >= 2 {
			res = true
			break
		}
	}
	return res
}

func main() {
	nums := []int{
		1, 1, 1, 3, 3, 4, 3, 2, 4, 2,
	}
	fmt.Println(containsDuplicatecontain(nums))
}
