package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	total := len1 + len2
	var i, l int = 0, 0
	var arr []int
	n := len1
	if len2 > len1 {
		n = len2
	}
	for i < n || l < n {
		if i < len1 && (l >= len2 || nums1[i] < nums2[l]) {
			arr = append(arr, nums1[i])
			i++
		} else {
			if l < len2 && (i >= len1 || nums1[i] >= nums2[l]) {
				arr = append(arr, nums2[l])
				l++
			}
		}
		if (i >= len1) && (l >= len2) {
			break
		}
	}
	var a float64 = 0
	mid := total / 2
	if total%2 == 0 {
		a = float64(arr[int(mid)-1]+arr[int(mid)]) / float64(2)
	} else {
		a = float64(arr[int(mid)])
	}
	return a
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	num := findMedianSortedArrays(nums1, nums2)
	fmt.Println(num)
}
