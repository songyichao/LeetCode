package test

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T)  {
	nums := []int{2, 7, 11, 15}
	fmt.Println(nums)
	res := twoSum(nums)
	fmt.Println(res)
}