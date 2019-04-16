package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	//双指针
	var res, sums [][]int
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				sums = append(sums, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}
	res = sums

	//先排序，查找时去重
	//var res, sums [][]int
	//n := len(nums)
	//sort.Ints(nums)
	//for i := 0; i < n-2; i++ {
	//	if i > 0 && nums[i] == nums[i-1] {
	//		continue
	//	}
	//	for m := i + 1; m < n; m++ {
	//		if m > i+1 && nums[m] == nums[m-1] {
	//			continue
	//		}
	//		for l := m + 1; l < n; l++ {
	//			if l > m+1 && nums[l] == nums[l-1] {
	//				continue
	//			}
	//			if nums[i]+nums[m]+nums[l] == 0 {
	//				sum := []int{nums[i], nums[m], nums[l],}
	//				sums = append(sums, sum)
	//			}
	//		}
	//	}
	//}
	//res = sums
	//先查找后去重
	//if len(sums) > 1 {
	//	var mSum []map[int]int
	//	m := make(map[int]int)
	//	if _, ok := m[sums[0][0]]; ok {
	//		m[sums[0][0]]++
	//	} else {
	//		m[sums[0][0]] = 1
	//	}
	//	if _, ok := m[sums[0][1]]; ok {
	//		m[sums[0][1]]++
	//	} else {
	//		m[sums[0][1]] = 1
	//	}
	//	if _, ok := m[sums[0][2]]; ok {
	//		m[sums[0][2]]++
	//	} else {
	//		m[sums[0][2]] = 1
	//	}
	//	res = append(res, sums[0])
	//	mSum = append(mSum, m)
	//	for i := 1; i < len(sums); i++ {
	//		q := 0
	//		o := make(map[int]int)
	//		for l := 0; l < len(mSum); l++ {
	//			p := 0
	//			r := make(map[int]int)
	//			for k, v := range mSum[l] {
	//				r[k] = v
	//			}
	//			for m := 0; m < 3; m++ {
	//				if _, ok := r[sums[i][m]]; ok && r[sums[i][m]] >= 1 {
	//					r[sums[i][m]]--
	//					p++
	//				}
	//			}
	//			if p == 3 {
	//				q = 1
	//				break
	//			}
	//		}
	//		for m := 0; m < 3; m++ {
	//			if _, ok := o[sums[i][m]]; ok {
	//				o[sums[i][m]]++
	//			} else {
	//				o[sums[i][m]] = 1
	//			}
	//		}
	//		if q == 0 {
	//			mSum = append(mSum, o)
	//			res = append(res, sums[i])
	//		}
	//	}
	//} else {
	//	res = sums
	//}
	return res
}

func main() {
	nums := []int{
		0,0,0,
	}
	sum := threeSum(nums)
	fmt.Println(sum)
}
