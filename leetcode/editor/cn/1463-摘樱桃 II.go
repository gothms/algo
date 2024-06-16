package main

import "fmt"

func main() {
	grid := [][]int{{3, 1, 1},
		{2, 5, 1},
		{1, 5, 5},
		{2, 1, 1}}
	grid = [][]int{{0, 8, 7, 10, 9, 10, 0, 9, 6},
		{8, 7, 10, 8, 7, 4, 9, 6, 10},
		{8, 1, 1, 5, 1, 5, 5, 1, 2},
		{9, 4, 10, 8, 8, 1, 9, 5, 0},
		{4, 3, 6, 10, 9, 2, 4, 8, 10},
		{7, 3, 2, 8, 3, 3, 5, 9, 8},
		{1, 2, 6, 5, 6, 2, 0, 10, 0}} // 96
	pickup := cherryPickup(grid)
	fmt.Println(pickup)
}

// leetcode submit region begin(Prohibit modification and deletion)
func cherryPickup(grid [][]int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func cherryPickup_(grid [][]int) int {
	// dp：个人
	// 考虑每行的任意两个樱桃的组合
	m, n := len(grid), len(grid[0])
	dp, temp := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i], temp[i] = make([]int, n), make([]int, n)
	}
	dp[0][n-1] = grid[0][0] + grid[0][n-1] // 第 0 行
	for i := 1; i < m; i++ {               // 第 i 行
		for l := 0; l <= min(i, n-2); l++ { // 左边机器人
			for r := max(l+1, n-1-i); r < n; r++ { // 右边机器人
				for nl := max(0, l-1); nl <= min(r, l+1); nl++ { // 当 l,r 相邻时，可以有 nl == r
					for nr := max(nl+1, r-1); nr <= min(n-1, r+1); nr++ { // nr 的左边界是 nl+1
						temp[l][r] = max(temp[l][r], dp[nl][nr]) // 重点是几个 for 的边界考虑
					}
				}
				temp[l][r] += grid[i][l] + grid[i][r] // 本行采集的草莓
			}
		}
		dp, temp = temp, dp // 滚动
	}
	var ans int
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}
