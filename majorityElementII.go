package main

import "fmt"

func majorityElementII(nums []int) []int {
	//摩尔投票法 改造版
	var res []int
	var res1, res2, cnt1, cnt2 = 0, 0, 0, 0
	for _, v := range nums {
		if v == res1 {
			cnt1++
		} else if v == res2 {
			cnt2++
		} else if cnt1 == 0 {
			res1 = v
			cnt1 = 1
		} else if cnt2 == 0 {
			res2 = v
			cnt2 = 1
		} else {
			cnt2--
			cnt1--
		}
	}
	cnt1 = 0
	cnt2 = 0
	for _, v := range nums {
		if v == res1 {
			cnt1++
		} else if v == res2 {
			cnt2++
		}
	}
	l := len(nums) / 3
	if cnt1 > l {
		res = append(res, res1)
	}
	if res1 != res2 && cnt2 > l {
		res = append(res, res2)
	}

	return res
	//取巧的hash
	//var res []int
	//l := len(nums) / 3
	//nMap := make(map[int]int)
	//rMap := make(map[int]int)
	//for _, v := range nums {
	//	nMap[v]++
	//	if nMap[v] > l {
	//		if rMap[v] == 0 {
	//			res = append(res, v)
	//			rMap[v] = 1
	//		}
	//	}
	//}
	//return res
}

func main() {
	nums := []int{
		22, 22,
	}
	fmt.Println(majorityElementII(nums))
}
