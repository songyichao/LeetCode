package main

import "fmt"

func findSubstring(s string, words []string) []int {
	ret := make([]int, 0)
	strLen := len(s)

	if strLen == 0 {
		return ret
	}

	wordsCount := len(words)
	if wordsCount == 0 {
		return ret
	}

	singleWordLen := len(words[0])

	record := make(map[string]int, 0)
	remain := make(map[string]int, 0)
	for _, word := range words {
		if len(word) != singleWordLen {
			return ret
		}
		if _, ok := record[word]; !ok {
			record[word] = 0
			remain[word] = 0
		}
		record[word]++
		remain[word]++
	}

	tmpWordsCount := wordsCount
	for i := 0; i < singleWordLen; i++ {
		start, cur := i, i
		reset := func() {
			if tmpWordsCount == wordsCount {
				// 没有使用
				return
			}

			for str, val := range record {
				remain[str] = val
			}
			tmpWordsCount = wordsCount
		}

		moveStart := func() {
			// start 指向下一个单词前，需要修改统计记录
			// 增加 start 指向的 word 可出现次数一次，
			remain[s[start:start+singleWordLen]]++
			// 连续符合条件的单词数减少一个
			tmpWordsCount++
			// start 后移一个 word 的长度
			start += singleWordLen
		}

		reset()

		// 字符串必须大于等于 所有单词的长度
		for strLen-start >= singleWordLen*wordsCount {
			tmpWord := s[cur : cur+singleWordLen]
			wordLeftTimes, ok := remain[tmpWord]

			switch {
			case !ok:
				// 不包含 下一组
				start, cur = cur+singleWordLen, cur+singleWordLen
				reset()
			case wordLeftTimes == 0:
				// word 用完了 移动start
				moveStart()
				// cur 不变
			default:
				remain[tmpWord]--
				tmpWordsCount--
				cur = cur + singleWordLen
				if tmpWordsCount == 0 {
					// 完全匹配
					ret = append(ret, start)
					moveStart()
				}
			}

		}

	}

	return ret

	//自己写的
	//var arr []int
	//if len(words) == 0 {
	//	return arr
	//}
	//strMap := make(map[string]int)
	//n := 0
	//m := len(words)
	//for _, v := range words {
	//	n = len(v)
	//	strMap[v]++
	//}
	//oneMap := make(map[string]int)
	//var i, l, o = 0, 0, -1
	//for i <= len(s) {
	//	//是否找到第一个在words里的字符串
	//	if i <= len(s)-n && strMap[s[i:i+n]] > oneMap[s[i:i+n]] {
	//		oneMap[s[i:i+n]]++
	//		if o == -1 {
	//			o = i
	//		}
	//		i += n
	//		l += n
	//		if i == len(s) && l == m*n {
	//			arr = append(arr, o)
	//			return arr
	//		}
	//	} else {
	//		oneMap = make(map[string]int)
	//		if l == m*n {
	//			arr = append(arr, o)
	//			i = o + 1
	//			l = 0
	//			o = -1
	//		} else if l == 0 {
	//			i++
	//		} else {
	//			i = o + 1
	//			l = 0
	//			o = -1
	//		}
	//	}
	//}
	//return arr
}

func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string{
		"word","good","best",
	}
	arr := findSubstring(s, words)
	fmt.Println(arr)
}
