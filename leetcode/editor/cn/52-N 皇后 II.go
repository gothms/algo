//n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：2
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 回溯 👍 473 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func totalNQueens(n int) int {
	cnt, bit := 0, 1<<n-1
	var dfs func(int, int, int, int)
	dfs = func(row, c, pie, na int) {
		if row == n {
			cnt++
			return
		}
		oks := bit & ^(c | pie | na)
		for last := 0; oks != 0; oks &^= last {
			last = oks & -oks
			dfs(row+1, c|last, (pie|last)<<1, (na|last)>>1)
		}
	}
	dfs(0, 0, 0, 0)
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
