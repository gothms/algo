//给你一个下标从 0 开始的 m x n 整数矩阵 grid 。你一开始的位置在 左上角 格子 (0, 0) 。
//
// 当你在格子 (i, j) 的时候，你可以移动到以下格子之一：
//
//
// 满足 j < k <= grid[i][j] + j 的格子 (i, k) （向右移动），或者
// 满足 i < k <= grid[i][j] + i 的格子 (k, j) （向下移动）。
//
//
// 请你返回到达 右下角 格子 (m - 1, n - 1) 需要经过的最少移动格子数，如果无法到达右下角格子，请你返回 -1 。
//
//
//
// 示例 1：
//
//
//
// 输入：grid = [[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]]
//输出：4
//解释：上图展示了到达右下角格子经过的 4 个格子。
//
//
// 示例 2：
//
//
//
// 输入：grid = [[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]]
//输出：3
//解释：上图展示了到达右下角格子经过的 3 个格子。
//
//
// 示例 3：
//
//
//
// 输入：grid = [[2,1,0],[1,0,0]]
//输出：-1
//解释：无法到达右下角格子。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10⁵
// 1 <= m * n <= 10⁵
// 0 <= grid[i][j] < m * n
// grid[m - 1][n - 1] == 0
//
//
// Related Topics 栈 广度优先搜索 并查集 数组 动态规划 矩阵 单调栈 堆（优先队列） 👍 24 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{5, 3, 2, 0}
	ret := sort.Search(len(arr), func(i int) bool {
		return arr[i] <= 1
	})
	fmt.Println(ret)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumVisitedCells(grid [][]int) int {
	// 单调栈
	ret, m, n := 0, len(grid), len(grid[0])
	colSts := make([][][2]int, n)
	rowSt := make([][2]int, 0)
	for i := m - 1; i >= 0; i-- { // 倒序遍历
		rowSt = rowSt[:0]
		for j := n - 1; j >= 0; j-- {
			colSt := colSts[j]
			if i < m-1 || j < n-1 { // ret 默认值
				ret = math.MaxInt32
			}
			if v := grid[i][j]; v > 0 {
				k := sort.Search(len(rowSt), func(x int) bool {
					return rowSt[x][0] <= j+v // 二分查找：[j,j+v] 区间内，单调栈已保证“最少次数”
				})
				if k < len(rowSt) {
					ret = rowSt[k][1]
				}
				k = sort.Search(len(colSt), func(x int) bool {
					return colSt[x][0] <= i+v
				})
				if k < len(colSt) {
					ret = min(ret, colSt[k][1])
				}
			}
			if ret < math.MaxInt32 { // 单调递增栈
				ret++
				for len(rowSt) > 0 && ret <= rowSt[len(rowSt)-1][1] {
					rowSt = rowSt[:len(rowSt)-1] // row[j][0]：索引是倒序，row[j][1]：已入栈的索引更大，且值更大，则被淘汰
				}
				rowSt = append(rowSt, [2]int{j, ret})
				for len(colSt) > 0 && ret <= colSt[len(colSt)-1][1] {
					colSt = colSt[:len(colSt)-1]
				}
				colSts[j] = append(colSt, [2]int{i, ret})
			}
		}
	}
	if ret < math.MaxInt32 {
		return ret
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
