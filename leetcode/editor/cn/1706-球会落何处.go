package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 1, -1, -1},
		{1, 1, 1, -1, -1},
		{-1, -1, -1, 1, 1},
		{1, 1, 1, 1, -1},
		{-1, -1, -1, -1, -1}}
	grid = [][]int{
		{1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1},
		{1, 1, 1, 1, 1, 1},
		{-1, -1, -1, -1, -1, -1}}
	ball := findBall(grid)
	fmt.Println(ball)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)
out:
	for b := 0; b < n; b++ {
		j := b
		for i := 0; i < m; i++ {
			v := grid[i][j] // 上一个位置
			j += v
			if j == n || j < 0 || v != grid[i][j] {
				ans[b] = -1
				continue out // 不成功
			}
		}
		ans[b] = j // 第几列出
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
