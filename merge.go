package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	var res [][]int
	tmp := make([][]int, len(intervals))
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	i := 0
	for k, v := range intervals {
		if k == 0 {
			tmp[k] = v
		} else {
			if v[0] <= tmp[k-i-1][1] {
				if v[1] >= tmp[k-i-1][1] {
					tmp[k-i-1][1] = v[1]
				}
				i++
			} else {
				tmp[k-i] = v
			}
		}
	}
	for _, v := range tmp {
		if v != nil {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	intervals := [][]int{
		{0, 1,}, {1, 4,}, {4, 5,}, {15, 18,},
	}
	fmt.Println(merge(intervals))
}
