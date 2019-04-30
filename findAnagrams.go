package main

import "fmt"

func findAnagrams(s string, p string) []int {
	//滑动窗口  滑滑滑
	//每次前移的长度一致 会保证sMap中的长度最多为len(p)最少为0
	var res []int
	l := len(p)
	m := make(map[string]int)
	for _, v := range p {
		m[string(v)]++
	}
	sMap := make(map[string]int)
	for k := 0; k < len(s); k++ {
		if k >= l {
			sMap[string(s[k-l])]--
		}
		sMap[string(s[k])]++
		f := 1
		for i, v := range m {
			if sMap[i] != v {
				f = 0
			}
		}
		if f == 1 {
			res = append(res, k-l+1)
		}
	}
	return res
	//又是一个超时的版本
	//var res []int
	//l := len(p)
	//m := make(map[string]int)
	//for _, v := range p {
	//	if _, ok := m[string(v)]; !ok {
	//		m[string(v)] = 1
	//
	//	} else {
	//		m[string(v)]++
	//	}
	//}
	//for k := 0; k < len(s)-l+1; k++ {
	//	o := make(map[string]int)
	//	if _, ok := m[string(s[k])]; ok {
	//		newS := string(s[k])
	//		o[string(s[k])] = 1
	//		for i := 1; i < l; i++ {
	//			if _, ok := m[string(s[k+i])]; ok && m[string(s[k+i])] != o[string(s[k+i])] {
	//				newS += string(s[k+i])
	//				o[string(s[k+i])]++
	//			}
	//		}
	//		if len(newS) == l {
	//			res = append(res, k)
	//		}
	//	}
	//}
	//return res
}
func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
