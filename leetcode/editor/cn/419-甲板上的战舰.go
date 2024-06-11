package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countBattleships(board [][]byte) int {
	ans := 0
	for i, b := range board {
		for j, v := range b {
			if v == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
				ans++
			}
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
