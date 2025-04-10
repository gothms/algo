package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func tictactoe(board []string) string {
	var row, col, pie, na int32
	var flag bool
	n := len(board)
	var ov, xv = int32('O' * n), int32('X' * n)
	for i, b := range board {
		row, col = 0, 0
		for j, v := range b {
			row += v
			col += int32(board[j][i])
			if v == ' ' {
				flag = true
			}
		}
		if row == ov || col == ov {
			return "O"
		} else if row == xv || col == xv {
			return "X"
		}
		pie += int32(board[i][n-1-i])
		na += int32(board[i][i])
	}
	if pie == ov || na == ov {
		return "O"
	} else if pie == xv || na == xv {
		return "X"
	}
	if flag {
		return "Pending"
	}
	return "Draw"
}

//leetcode submit region end(Prohibit modification and deletion)
