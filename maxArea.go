package main

import (
	"fmt"
)

func maxArea(height []int) int {
	//双指针
	i, j := 0, len(height)-1
	ans := 0
	for ; i < j; {
		h := height[i]
		if height[j] < h {
			h = height[j]
		}
		if h*(j-i) > ans {
			ans = h * (j - i)
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return ans

	//maxArea := 0
	//n := len(height)
	//for i := 0; i < n; i++ {
	//	for l := 0; l < n; l++ {
	//		a := height[i]
	//		if l != i {
	//			if height[l] < a {
	//				a = height[l]
	//			}
	//			b := a * int(math.Abs(float64(i-l)))
	//			if maxArea < b {
	//				maxArea = b
	//			}
	//		}
	//	}
	//}
	//return maxArea
}
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	maxArea := maxArea(height)
	fmt.Println(maxArea)
}
