package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n < 0{
		x = 1/x
		n = -n
	}
	var result float64 = 1.0
	for n > 0{
		if n & 1 == 1{
			result *= x
		}
		x *=x
		n >>=1
	}
	return result
	//超时版本
	//if x == 1.0 {
	//	return 1.0
	//}
	//res := x
	//f := 1
	//if n < 0 {
	//	n = -n
	//	f = -1
	//}
	//for i := n; i < n; i++ {
	//	if f == 1 {
	//		res *= res
	//	} else {
	//		res /= res
	//	}
	//	if res < -2147483648 {
	//		res = -2147483648
	//		break
	//	}
	//	if res > 2147483647 {
	//		res = 2147483647
	//		break
	//	}
	//}
	//return res
}

func main() {
	x := 2.10000
	n := 3
	fmt.Println(myPow(x, n))
}
