//给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
//
// 说明：每次只能向下或者向右移动一步。
//
//
//
// 示例 1：
//
//
//输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
//输出：7
//解释：因为路径 1→3→1→1→1 的总和最小。
//
//
// 示例 2：
//
//
//输入：grid = [[1,2,3],[4,5,6]]
//输出：12
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 200
// 0 <= grid[i][j] <= 200
//
//
// Related Topics 数组 动态规划 矩阵 👍 1674 👎 0

package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6}}
	i := minPathSum(grid)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minPathSum(grid [][]int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func minPathSum(grid [][]int) int {
//    m, n := len(grid), len(grid[0])
//    dp := make([]int, n+1)
//    for i := range dp {
//        dp[i] = math.MaxInt32 // 封路
//    }
//    dp[n-1] = 0                   // 开路
//    for i := m - 1; i >= 0; i-- { // 自底向上
//        for j := n - 1; j >= 0; j-- {
//            dp[j] = min(dp[j], dp[j+1]) + grid[i][j] // 状态转移方程
//        }
//    }
//    return dp[0]
//}
