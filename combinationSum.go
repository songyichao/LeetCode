package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(candidates []int, target int) [][]int {
	l := len(candidates)
	if l == 0 {
		return nil
	}
	var res [][]int
	next := target - candidates[0]
	switch {
	case next < 0:
		return nil
	case next == 0:
		return [][]int{{candidates[0]}}
	case next > 0:
		combines := dfs(candidates, next)
		for _, combine := range combines {
			res = append(res, append(combine, candidates[0]))
		}
	}
	res = append(res, dfs(candidates[1:], target)...)
	return res
}

func main() {
	candidates := []int{
		2, 3, 6, 7,
	}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
