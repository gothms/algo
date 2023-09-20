//骑士在一张 n x n 的棋盘上巡视。在有效的巡视方案中，骑士会从棋盘的 左上角 出发，并且访问棋盘上的每个格子 恰好一次 。
//
// 给你一个 n x n 的整数矩阵 grid ，由范围 [0, n * n - 1] 内的不同整数组成，其中 grid[row][col] 表示单元格 (
//row, col) 是骑士访问的第 grid[row][col] 个单元格。骑士的行动是从下标 0 开始的。
//
// 如果 grid 表示了骑士的有效巡视方案，返回 true；否则返回 false。
//
// 注意，骑士行动时可以垂直移动两个格子且水平移动一个格子，或水平移动两个格子且垂直移动一个格子。下图展示了骑士从某个格子出发可能的八种行动路线。
//
//
//
// 示例 1：
// 输入：grid = [[0,11,16,5,20],[17,4,19,10,15],[12,1,8,21,6],[3,18,23,14,9],[24,13
//,2,7,22]]
//输出：true
//解释：grid 如上图所示，可以证明这是一个有效的巡视方案。
//
//
// 示例 2：
// 输入：grid = [[0,3,6],[5,8,1],[2,7,4]]
//输出：false
//解释：grid 如上图所示，考虑到骑士第 7 次行动后的位置，第 8 次行动是无效的。
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 3 <= n <= 7
// 0 <= grid[row][col] < n * n
// grid 中的所有整数 互不相同
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 模拟 👍 8 👎 0

package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 3, 6},
		{5, 8, 1},
		{2, 7, 4}}
	grid = [][]int{
		{0, 13, 1, 7, 20},
		{3, 8, 19, 12, 15},
		{18, 2, 14, 21, 6},
		{9, 4, 23, 16, 11},
		{24, 17, 10, 5, 22}}
	validGrid := checkValidGrid(grid)
	fmt.Println(validGrid)
}

// leetcode submit region begin(Prohibit modification and deletion)
func checkValidGrid(grid [][]int) bool {
	// lc
	n := len(grid)
	if grid[0][0] != 0 || n == 3 { // fast path
		return false
	}
	absVal := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	step := make([][2]int, n*n)
	for i, r := range grid {
		for j, v := range r {
			step[v] = [2]int{i, j} // 对既定的步骤排序
		}
	}
	for i := 1; i < len(step); i++ {
		if absVal((step[i][0]-step[i-1][0])*(step[i][1]-step[i-1][1])) == 2 {
			continue // 有效的下一步
		}
		return false // 无效的下一步
	}
	return true

	// 迭代
	//	n := len(grid)
	//	if grid[0][0] != 0 || n == 3 { // fast fail
	//		return false
	//	}
	//	d := [8][2]int{
	//		{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}
	//	v, m := 1, n*n // 路线区间
	//out:
	//	for i, j := 0, 0; v < m; v++ {
	//		for _, p := range d {
	//			x, y := i+p[0], j+p[1]
	//			if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == v {
	//				i, j = x, y // 找到下一步
	//				continue out
	//			}
	//		}
	//		return false // 没找到下一步
	//	}
	//	return true

	// dfs
	//n, can := len(grid), len(grid)*len(grid)-1
	//if n == 3 {
	//	return false
	//}
	//d := [8][2]int{
	//	{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}
	//var dfs func(int, int, int) bool
	//dfs = func(i, j, v int) bool {
	//	if v == can {
	//		return true
	//	}
	//	v++
	//	x, y := 0, 0
	//	for _, p := range d {
	//		x, y = i+p[0], j+p[1]
	//		if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == v &&
	//			dfs(x, y, v) {
	//			return true
	//		}
	//	}
	//	return false
	//}
	//return dfs(0, 0, 0)

	// 原代码：[0,0] 可以是任意数
	//if grid[0][0] != 0 {	// 坑！[0,0] 必须为 0
	//	return false
	//}
	//n, can := len(grid), len(grid)*len(grid)
	//d := [8][2]int{
	//	{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}
	//var dfs func(int, int, int, int) bool
	//dfs = func(i, j, diff, v int) bool {
	//	v += diff
	//	if v < 0 || v == can {
	//		return true
	//	}
	//	x, y := 0, 0
	//	for _, p := range d {
	//		x, y = i+p[0], j+p[1]
	//		if x >= 0 && x < n && y >= 0 && y < n && grid[x][y] == v &&
	//			dfs(x, y, diff, v) {
	//			return true
	//		}
	//	}
	//	return false
	//}
	//return dfs(0, 0, 1, grid[0][0]) && dfs(0, 0, -1, grid[0][0])
}

//leetcode submit region end(Prohibit modification and deletion)
