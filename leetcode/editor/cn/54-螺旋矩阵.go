package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func spiralOrder(matrix [][]int) []int {
	const inf = -102
	d := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	m, n := len(matrix), len(matrix[0])
	ans := make([]int, m*n)
	for i, k, x, y := 0, 0, 0, -1; i < m*n; {
		if nx, ny := x+d[k][0], y+d[k][1]; nx < 0 || nx == m || ny < 0 || ny == n || matrix[nx][ny] == inf {
			k = (k + 1) % 4
		} else {
			ans[i], x, y = matrix[nx][ny], nx, ny
			i++
			matrix[x][y] = inf
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
