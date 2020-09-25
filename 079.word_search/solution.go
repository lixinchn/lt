package leetcode

func exist(board [][]byte, word string) bool {
	flag := make(map[int]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				key := i*1000 + j
				flag[key] = true
				if do(board, word, 1, i, j, flag) {
					return true
				}
				delete(flag, key)
			}
		}
	}
	return false
}

func do(board [][]byte, word string, index, x, y int, flag map[int]bool) bool {
	if index == len(word) {
		return true
	}
	if x+1 < len(board) && board[x+1][y] == word[index] {
		key := (x+1)*1000 + y
		if _, ok := flag[key]; !ok {
			flag[key] = true
			if do(board, word, index+1, x+1, y, flag) {
				return true
			}
			delete(flag, key)
		}
	}
	if y+1 < len(board[0]) && board[x][y+1] == word[index] {
		key := x*1000 + (y + 1)
		if _, ok := flag[key]; !ok {
			flag[key] = true
			if do(board, word, index+1, x, y+1, flag) {
				return true
			}
			delete(flag, key)
		}
	}
	if x-1 > -1 && board[x-1][y] == word[index] {
		key := (x-1)*1000 + y
		if _, ok := flag[key]; !ok {
			flag[key] = true
			if do(board, word, index+1, x-1, y, flag) {
				return true
			}
			delete(flag, key)
		}
	}
	if y-1 > -1 && board[x][y-1] == word[index] {
		key := x*1000 + y - 1
		if _, ok := flag[key]; !ok {
			flag[key] = true
			if do(board, word, index+1, x, y-1, flag) {
				return true
			}
			delete(flag, key)
		}
	}
	return false
}
