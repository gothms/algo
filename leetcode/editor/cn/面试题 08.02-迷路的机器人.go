package main

import (
	"fmt"
	"slices"
)

func main() {
	grid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	obstacles := pathWithObstacles(grid)
	fmt.Println(obstacles)
}

// leetcode submit region begin(Prohibit modification and deletion)
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	var ans [][]int
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return ans
	}
	path, vis := [][]int{{0, 0}}, make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if x == m-1 && y == n-1 {
			ans = slices.Clone(path)
			return true
		}
		if y+1 < n && obstacleGrid[x][y+1] == 0 && !vis[x][y+1] {
			vis[x][y+1] = true
			path = append(path, []int{x, y + 1})
			if dfs(x, y+1) {
				return true
			}
			path = path[:len(path)-1]
		}
		if x+1 < m && obstacleGrid[x+1][y] == 0 && !vis[x+1][y] {
			vis[x+1][y] = true
			path = append(path, []int{x + 1, y})
			if dfs(x+1, y) {
				return true
			}
			path = path[:len(path)-1]
		}
		return false
	}
	dfs(0, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
