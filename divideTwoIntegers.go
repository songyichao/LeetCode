package main

import "fmt"

func divide(dividend int, divisor int) int {
	res := 0
	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}
	for dividend >= divisor {
		var tmp, i = divisor, 1
		for dividend-tmp >= 0 {
			dividend -= tmp
			res += i
			i <<= 1
			tmp <<= 1
		}
	}
	if sign < 0 {
		res = -res
	}
	if res > 2147483647 {
		res = 2147483647
	}
	if res < -2147483648 {
		res = -2147483648
	}
	return res
	//加法 shi
	//res := 0
	//sum := 0
	//sign := 1
	//if dividend < 0 {
	//	sign = -sign
	//	dividend = -dividend
	//}
	//if divisor < 0 {
	//	sign = -sign
	//	divisor = -divisor
	//}
	//for i := 0; sum <= dividend; i++ {
	//	sum += divisor
	//	res = i
	//}
	//if sign < 0 {
	//	res = -res
	//}
	//if res > 2147483647 {
	//	res = 2147483647
	//}
	//if res < -2147483648 {
	//	res = -2147483648
	//}
	//return res
}

func main() {
	dividend := 36
	divisor := 4
	num := divide(dividend, divisor)
	fmt.Println(num)
}
