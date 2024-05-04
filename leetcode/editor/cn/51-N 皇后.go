//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
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
// Related Topics 数组 回溯 👍 1928 👎 0

package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func main() {
	n := 4
	queens := solveNQueens(n)
	fmt.Println(queens)
}

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	ans, path := make([][]string, 0), make([]int, n)
	dots := make([]byte, n)
	for i := range dots {
		dots[i] = '.'
	}
	nq := func() []string {
		snq := make([]string, n)
		for i, v := range path {
			idx := bits.Len(uint(v)) - 1
			dots[idx] = 'Q'
			snq[i] = string(slices.Clone(dots))
			dots[idx] = '.'
		}
		return snq
	}
	var dfs func(int, int, int, int)
	dfs = func(i, c, pie, na int) {
		if i == n {
			ans = append(ans, nq()) // 找到一个解法
			return
		}
		bit := ^(c | pie | na) & (1<<n - 1)
		//bit := (1<<n - 1) &^ (c | pie | na) // na 可能超出 1<<n-1，所以用 &^
		for ; bit > 0; bit = bit & (bit - 1) { // 去掉最后一位
			k := bit & -bit // 取出最后一位
			path[i] = k
			dfs(i+1, c|k, (pie|k)>>1, (na|k)<<1)
		}
	}
	dfs(0, 0, 0, 0)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
