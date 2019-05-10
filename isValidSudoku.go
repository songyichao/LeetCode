package main

import "fmt"

func isValidSudoku(board [][]string) bool {
	res := true
	//行map
	lMap := make(map[string]int)
	//列map
	rMap := make(map[string]int)
	//块map
	bMap := make([]map[string]int, 9)
	for i := 0; i < 9; i++ {
		for l := 0; l < 9; l++ {
			sIL:=string(board[i][l])
			sLI:=string(board[l][i])
			//处理行
			lMap[sIL]++
			//处理行
			rMap[sLI]++
			//处理块
			bIndex := (i/3)*3 + l/3
			if bMap[bIndex] == nil {
				bMap[bIndex] = make(map[string]int)
			}
			bMap[bIndex][sIL]++
			if (sIL != "." && (lMap[sIL] > 1 || bMap[bIndex][sIL] > 1)) || (sLI != "." && rMap[sLI] > 1){
				res = false
				break
			}
		}
		lMap = make(map[string]int)
		rMap = make(map[string]int)
		if !res {
			break
		}
	}
	return res
}

func main() {
	board := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", "5", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	fmt.Println(isValidSudoku(board))
}
