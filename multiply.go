package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	numMap := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}
	numArr := []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	res := ""
	l1 := len(num1)
	l2 := len(num2)
	var resArr []int
	n := 0
	for i := l1 - 1; i >= 0; i-- {
		numS := string(num1[i])
		o := 0
		for m := l2 - 1; m >= 0; m-- {
			num2S := string(num2[m])
			resV := 0
			if n < len(resArr) && n+(l2-1-m) < len(resArr) {
				resV = resArr[n+(l2-1-m)]
			}
			s := numMap[numS]*numMap[num2S] + o + resV
			q := s % 10
			if n+(l2-1-m) < len(resArr) {
				resArr[n+(l2-1-m)] = q
			} else {
				resArr = append(resArr, q)
			}
			o = (s - q) / 10
		}
		if o != 0 {
			resArr = append(resArr, o)
		}
		n++
	}
	for _, v := range resArr {
		res = numArr[v] + res
	}
	return res
}

func main() {
	num1 := "140"
	num2 := "721"
	fmt.Println(multiply(num1, num2))
}
