package main

import "fmt"

func majorityElement(nums []int) int {
	//摩尔投票法
	res := 0
	cnt := 0
	for _, v := range nums {
		if cnt == 0 {
			res = v
			cnt++
		} else {
			if res == v {
				cnt++
			} else {
				cnt--
			}
		}
	}

	return res
	//hash

	//res := 0
	//l := len(nums) / 2
	//nMap := make(map[int]int)
	//for _, v := range nums {
	//	nMap[v]++
	//	if nMap[v]>l {
	//		res = v
	//		break
	//	}
	//}
	//return res
}

func main() {
	nums := []int{
		3, 2, 3,
	}
	fmt.Println(majorityElement(nums))
}
