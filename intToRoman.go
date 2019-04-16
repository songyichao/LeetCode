package main

import "fmt"

func intToRoman(num int) string {
	romanMap := []struct {
		k int
		v string
	}{
		{1000, "M"},
		{500, "D"},
		{100, "C"},
		{50, "L"},
		{10, "X"},
		{5, "V"},
		{1, "I"},
	}
	roman := ""
	if num >= 1 && num <= 3999 {
		for k, v := range romanMap {
			if num == 0 {
				break
			}
			a := num / v.k
			if a > 0 {
				for i := 0; i < a; i++ {
					roman += v.v
				}
			}
			num -= a * v.k
			if k < len(romanMap)-1 {
				if num%v.k >= romanMap[k].k-romanMap[k+1].k && (romanMap[k].k == 5 || romanMap[k].k == 50 || romanMap[k].k == 500) {
					roman += romanMap[k+1].v + v.v
					num -= romanMap[k].k - romanMap[k+1].k
					continue
				} else if k < len(romanMap)-2 && num%v.k >= romanMap[k].k-romanMap[k+2].k {
					roman += romanMap[k+2].v + v.v
					num -= romanMap[k].k - romanMap[k+2].k
					continue
				}
			}
		}

	}
	return roman
}
func main() {
	num := 1994
	roman := intToRoman(num)
	fmt.Println(roman)
}
