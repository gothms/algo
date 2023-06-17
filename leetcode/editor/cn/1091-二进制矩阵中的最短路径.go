//给你一个 n x n 的二进制矩阵 grid 中，返回矩阵中最短 畅通路径 的长度。如果不存在这样的路径，返回 -1 。
//
// 二进制矩阵中的 畅通路径 是一条从 左上角 单元格（即，(0, 0)）到 右下角 单元格（即，(n - 1, n - 1)）的路径，该路径同时满足下述要求
//：
//
//
// 路径途经的所有单元格都的值都是 0 。
// 路径中所有相邻的单元格应当在 8 个方向之一 上连通（即，相邻两单元之间彼此不同且共享一条边或者一个角）。
//
//
// 畅通路径的长度 是该路径途经的单元格总数。
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,1],[1,0]]
//输出：2
//
//
// 示例 2：
//
//
//输入：grid = [[0,0,0],[1,1,0],[1,1,0]]
//输出：4
//
//
// 示例 3：
//
//
//输入：grid = [[1,0,0],[1,1,0],[1,1,0]]
//输出：-1
//
//
//
//
// 提示：
//
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 100
// grid[i][j] 为 0 或 1
//
//
// Related Topics 广度优先搜索 数组 矩阵 👍 288 👎 0

package main

func main() {

}

/*
思路：典型的 bfs
	1.求最短路径，该类问题 bfs 优于 dfs
		在 bfs 中，先找到的结果，所用的寻找次数，肯定低于后找到的结果
	2.搜索问题中，很重要的一点是考虑要不要记忆化搜索
		很显然这里需要做记忆化搜索，以防出现轮回
	3.其他的就是一些细节问题：
		比如可以先判断起点，终点
		在哪里计数次数，在哪里判断终点等
*/
//leetcode submit region begin(Prohibit modification and deletion)
func shortestPathBinaryMatrix(grid [][]int) int {
	cnt, n := 1, len(grid)-1
	if grid[0][0] == 1 || grid[n][n] == 1 {
		return -1
	}
	ok := func(x, y int) bool {
		return x >= 0 && x <= n && y >= 0 && y <= n && grid[x][y] == 0
	}
	memo := [][2]int{{n, n}}
	grid[n][n] = 1
	dx, dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}, []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for l := len(memo); l > 0; l = len(memo) {
		for i := 0; i < l; i++ {
			x, y := memo[i][0], memo[i][1]
			if x|y == 0 {
				return cnt
			}
			for j := 0; j < 8; j++ {
				if nx, ny := x+dx[j], y+dy[j]; ok(nx, ny) {
					memo, grid[nx][ny] = append(memo, [2]int{nx, ny}), 1
				}
			}
		}
		memo, cnt = memo[l:], cnt+1
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
