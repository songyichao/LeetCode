package main

import "fmt"

var l int

func backtrack(str *[]string, ans string, open int, close int) {
	if len(ans) == 2*l {
		*str = append(*str, ans)
		return
	}
	if open < l {
		backtrack(str, ans+"(", open+1, close)
	}
	if open > close {
		backtrack(str, ans+")", open, close+1)
	}
}
func generateParenthesis(n int) []string {
	//回溯
	//只有在我们知道序列仍然保持有效时才添加 '(' or ')'，而不是像方法一那样每次添加。我们可以通过跟踪到目前为止放置的左括号和右括号的数目来做到这一点，
	//
	//如果我们还剩一个位置，我们可以开始放一个左括号。 如果它不超过左括号的数量，我们可以放一个右括号。
	var str []string
	l = n
	backtrack(&str, "", 0, 0)
	return str
	//找到可以放置一整个括号的位置插入括号
	//defaultStr := "()"
	//var str []string
	//if n == 0 {
	//	str = append(str, "")
	//}
	//for i := 0; i < n; i++ {
	//	if len(str) < 1 {
	//		str = append(str, defaultStr)
	//	} else {
	//		var strSub []string
	//		strMap := make(map[string]string)
	//		for _, v := range str {
	//			m := len(v)
	//			for l := 0; l < m; l++ {
	//				a := string(v[l])
	//				b := ""
	//				if l < m-1 && ((a == "(" && string(v[l+1]) == ")") || (a == "(" && string(v[l+1]) == "(")) {
	//					b = string(v[:l+1]) + defaultStr + string(v[l+1:m])
	//					strSub = append(strSub, )
	//				} else if l == m-1 && a == ")" && string(v[l-1]) == "(" {
	//					b = v + defaultStr
	//				}
	//				if _, ok := strMap[b]; !ok && b != "" {
	//					strMap[b] = b
	//					strSub = append(strSub, b)
	//				}
	//			}
	//		}
	//		str = strSub
	//	}
	//}
	//return str
}

func main() {
	n := 4
	str := generateParenthesis(n)
	fmt.Println(str)
}
