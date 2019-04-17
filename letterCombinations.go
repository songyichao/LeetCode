package main

import "fmt"

func letterCombinations(digits string) []string {
	num := map[string][]string{
		"2": {"a", "b", "c",},
		"3": {"d", "e", "f",},
		"4": {"g", "h", "i",},
		"5": {"j", "k", "l",},
		"6": {"m", "n", "o",},
		"7": {"p", "q", "r", "s",},
		"8": {"t", "u", "v",},
		"9": {"w", "x", "y", "z",},
	}
	var strArr []string
	for _, v := range digits {
		n := num[string(v)]
		var aStrArr []string
		for _, val := range n {
			if len(strArr) > 0 {
				for _, value := range strArr {
					aStrArr = append(aStrArr, value+val)
				}
			} else {
				aStrArr = append(aStrArr, val)
			}
		}
		strArr = aStrArr
	}

	return strArr
}
func main() {
	digits := "237"
	strArr := letterCombinations(digits)
	fmt.Println(strArr)
}
