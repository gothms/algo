//给你一个 m * n 的矩阵 seats 表示教室中的座位分布。如果座位是坏的（不可用），就用 '#' 表示；否则，用 '.' 表示。
//
// 学生可以看到左侧、右侧、左上、右上这四个方向上紧邻他的学生的答卷，但是看不到直接坐在他前面或者后面的学生的答卷。请你计算并返回该考场可以容纳的同时参加考试
//且无法作弊的 最大 学生人数。
//
// 学生必须坐在状况良好的座位上。
//
//
//
// 示例 1：
//
//
//
//
//输入：seats = [["#",".","#","#",".","#"],
//              [".","#","#","#","#","."],
//              ["#",".","#","#",".","#"]]
//输出：4
//解释：教师可以让 4 个学生坐在可用的座位上，这样他们就无法在考试中作弊。
//
//
// 示例 2：
//
//
//输入：seats = [[".","#"],
//              ["#","#"],
//              ["#","."],
//              ["#","#"],
//              [".","#"]]
//输出：3
//解释：让所有学生坐在可用的座位上。
//
//
// 示例 3：
//
//
//输入：seats = [["#",".",".",".","#"],
//              [".","#",".","#","."],
//              [".",".","#",".","."],
//              [".","#",".","#","."],
//              ["#",".",".",".","#"]]
//输出：10
//解释：让学生坐在第 1、3 和 5 列的可用座位上。
//
//
//
//
// 提示：
//
//
// seats 只包含字符 '.' 和'#'
// m == seats.length
// n == seats[i].length
// 1 <= m <= 8
// 1 <= n <= 8
//
//
// Related Topics 位运算 数组 动态规划 状态压缩 矩阵 👍 169 👎 0

package main

import (
	"math/bits"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxStudents(seats [][]byte) int {
	// 状压 dp
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	ms, m, n := 0, len(seats), len(seats[0])
	dp, z := make(map[int]int), 1<<n
	checkRow := func(r, s int) bool { // 当前行的合法性
		for i := 0; i < n; i++ {
			if (s>>i)&1 == 1 && (seats[r][i] == '#' || i > 0 && (s>>(i-1))&1 == 1) {
				return false
			}
		}
		return true
	}
	checkPre := func(d, u int) bool { // d：当前行，u：上一行，检查上一行的合法性
		for i := 0; i < n; i++ {
			if (d>>i)&1 == 1 && (i > 0 && (u>>(i-1))&1 == 1 || i < n-1 && (u>>(i+1))&1 == 1) {
				return false
			}
		}
		return true
	}
	var dfs func(int, int) int
	dfs = func(r, s int) int {
		key := r<<n + s
		if _, ok := dp[key]; !ok {
			if !checkRow(r, s) {
				dp[key] = -1
				return -1
			}
			c, p := bits.OnesCount(uint(s)), 0 // 当前行的座位数
			if r > 0 {
				for i := 0; i < z; i++ { // 尝试上一行的座位
					if checkPre(s, i) {
						p = maxVal(p, dfs(r-1, i))
					}
				}
			}
			dp[key] = c + p // 记忆化
		}
		return dp[key]
	}
	for i := 0; i < z; i++ {
		ms = maxVal(ms, dfs(m-1, i)) // 最后一行
	}
	return ms
}

//leetcode submit region end(Prohibit modification and deletion)
