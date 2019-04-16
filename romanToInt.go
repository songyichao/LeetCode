package main

import "fmt"

func romanToInt(s string) int {
	num := 0
	romanMap := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		if i > 0 {
			if (string(s[i]) == "V" && string(s[i-1]) == "I") ||
				(string(s[i]) == "X" && string(s[i-1]) == "I") ||
				(string(s[i]) == "L" && string(s[i-1]) == "X") ||
				(string(s[i]) == "C" && string(s[i-1]) == "X") ||
				(string(s[i]) == "D" && string(s[i-1]) == "C") ||
				(string(s[i]) == "M" && string(s[i-1]) == "C") {
				num += romanMap[string(s[i])] - romanMap[string(s[i-1])]
				i--
				continue
			}
		}
		num += romanMap[string(s[i])]
	}
	return num
}

func main() {
	roman := "LVIII"
	i := romanToInt(roman)
	fmt.Println(i)
}
