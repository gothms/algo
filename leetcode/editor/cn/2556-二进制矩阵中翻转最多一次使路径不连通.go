//给你一个下标从 0 开始的 m x n 二进制 矩阵 grid 。你可以从一个格子 (row, col) 移动到格子 (row + 1, col) 或者 (
//row, col + 1) ，前提是前往的格子值为 1 。如果从 (0, 0) 到 (m - 1, n - 1) 没有任何路径，我们称该矩阵是 不连通 的。
//
// 你可以翻转 最多一个 格子的值（也可以不翻转）。你 不能翻转 格子 (0, 0) 和 (m - 1, n - 1) 。
//
// 如果可以使矩阵不连通，请你返回 true ，否则返回 false 。
//
// 注意 ，翻转一个格子的值，可以使它的值从 0 变 1 ，或从 1 变 0 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[1,1,1],[1,0,0],[1,1,1]]
//输出：true
//解释：按照上图所示我们翻转蓝色格子里的值，翻转后从 (0, 0) 到 (2, 2) 没有路径。
//
//
// 示例 2：
//
//
//
//
//输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
//输出：false
//解释：无法翻转至多一个格子，使 (0, 0) 到 (2, 2) 没有路径。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 1000
// 1 <= m * n <= 10⁵
// grid[0][0] == grid[m - 1][n - 1] == 1
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 动态规划 矩阵 👍 26 👎 0

package main

import "fmt"

func main() {
	grid := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	path := isPossibleToCutPath(grid)
	fmt.Println(path)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isPossibleToCutPath(grid [][]int) bool {
	//m, n := len(grid), len(grid[0])
	//if grid[0][0]+grid[m-1][n-1] != 2 {
	//	return true
	//}
	//memo := make([][]bool, m)
	//for i := 0; i < m; i++ {
	//	memo[i] = make([]bool, n)
	//}
	//var dfs func(int, int, bool) bool
	//dfs = func(i, j int, reverse bool) bool {
	//	//fmt.Println(i,j)
	//	if i < 0 || j < 0 || grid[i][j] == 0 { // 0 阻塞
	//		return true
	//	}
	//	if  i == 0 && j == 0 { // 已经翻转过 or 到达起点
	//		return false
	//	}
	//	//fmt.Println(i, j)
	//	if !reverse {
	//		if i==0&&
	//	}
	//	return dfs(i-1, j, false) && dfs(i, j-1, false) ||
	//}
	////return dfs(m-2, n-1, false) && dfs(m-1, n-2, false)
	//return dfs(m-1, n-1, false)

	// lc
	m, n := len(grid)-1, len(grid[0])-1
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i < 0 || j < 0 || grid[i][j] == 0 { // 此路不通
			return true
		}
		if i == 0 && j == 0 { // 到达起点
			return false
		}
		grid[i][j] = 0                    // 已走过的路径
		return dfs(i-1, j) && dfs(i, j-1) // 左和上都不通，则不通
	}
	if dfs(m, n) {
		return true
	}
	grid[m][n] = 1   // 改回来
	return dfs(m, n) // 重要：两次 dfs
}

//leetcode submit region end(Prohibit modification and deletion)
