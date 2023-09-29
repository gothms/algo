//给你一个 m x n 的二进制矩阵 mat ，请你返回有多少个 子矩形 的元素全部都是 1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：mat = [[1,0,1],[1,1,0],[1,1,0]]
//输出：13
//解释：
//有 6 个 1x1 的矩形。
//有 2 个 1x2 的矩形。
//有 3 个 2x1 的矩形。
//有 1 个 2x2 的矩形。
//有 1 个 3x1 的矩形。
//矩形数目总共 = 6 + 2 + 3 + 1 + 1 = 13 。
//
//
// 示例 2：
//
//
//
//
//输入：mat = [[0,1,1,0],[0,1,1,1],[1,1,1,0]]
//输出：24
//解释：
//有 8 个 1x1 的子矩形。
//有 5 个 1x2 的子矩形。
//有 2 个 1x3 的子矩形。
//有 4 个 2x1 的子矩形。
//有 2 个 2x2 的子矩形。
//有 2 个 3x1 的子矩形。
//有 1 个 3x2 的子矩形。
//矩形数目总共 = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24 。
//
//
//
//
//
// 提示：
//
//
// 1 <= m, n <= 150
// mat[i][j] 仅包含 0 或 1
//
//
// Related Topics 栈 数组 动态规划 矩阵 单调栈 👍 173 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func numSubmat(mat [][]int) int {
	// 单调栈
	m, n := len(mat), len(mat[0])
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n+1)
	}
	for i, row := range mat {
		for j, v := range row {
			if v == 1 {
				memo[i][j+1] += memo[i][j] + 1 // 记录每行的宽度
			}
		}
	}
	ms, col, stack := 0, make([]int, m), make([]int, 1, m+1)
	stack[0] = -1
	for j := 1; j <= n; j++ {
		stack = stack[:1] // -1 为哨兵
		for i := 0; i < m; i++ {
			last := len(stack) - 1
			for last > 0 && memo[stack[last]][j] >= memo[i][j] { // 单调递增栈
				last--
			}
			col[i] = memo[i][j] * (i - stack[last]) // 宽 * 高：(stack[last], i] 区间内的索引，宽度都 >=memo[i][j]
			if last > 0 {
				col[i] += col[stack[last]] // <memo[i][j] 的区间的总数为 col[stack[last]]
			}
			ms += col[i]
			stack = stack[:last+1] // 维持单调栈
			stack = append(stack, i)
		}
	}
	return ms

	// 迭代
}

//leetcode submit region end(Prohibit modification and deletion)
