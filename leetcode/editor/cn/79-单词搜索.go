package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(i, j, idx int) bool {
		if idx == len(word) {
			return true
		}
		if i < 0 || j < 0 || i == m || j == n || board[i][j] != word[idx] {
			return false
		}
		idx++
		c := board[i][j]
		board[i][j] = ' '
		if dfs(i, j-1, idx) || dfs(i-1, j, idx) || dfs(i, j+1, idx) || dfs(i+1, j, idx) {
			return true
		}
		board[i][j] = c // 回溯
		return false
	}
	for i := range board {
		for j := range board[i] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
