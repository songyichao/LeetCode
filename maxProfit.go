package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	if len(prices) <= 0 {
		return res
	}
	min := prices[0]
	for _, v := range prices {
		if v < min {
			min = v
		}
		if v-min > res {
			res = v - min
		}
	}
	return res
}

func main() {
	prices := []int{
		7, 6, 4, 3, 1,
	}
	fmt.Println(maxProfit(prices))
}
