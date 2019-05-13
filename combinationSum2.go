package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs2(candidates, target)
}

func dfs2(candidates []int, target int) [][]int {
	l := len(candidates)
	if l == 0 {
		return nil
	}
	var res [][]int
	next := target - candidates[0] // 下一个和
	switch {
	case next < 0:
		return nil
	case next == 0:
		return [][]int{{candidates[0]}}
	case next > 0: // 寻找子和的任意组合
		combines := dfs2(candidates[1:], next)
		for _, combine := range combines {
			res = append(res, append(combine, candidates[0])) // 将子和的任意组合添加到解中
		}
	}
	s := 1
	for i := 1; i < l; i++ {
		if candidates[i] == candidates[0] {
			s = i+1
		} else {
			break
		}
	}
	res = append(res, dfs2(candidates[s:], target)...) // 在后续数组中尝试和为 target 的组合
	return res
}

func main() {
	candidates := []int{
		10,1,2,7,6,1,5,
	}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
