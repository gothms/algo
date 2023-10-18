//设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整
//个棋盘的那两条对角线。
//
// 注意：本题相对原题做了扩展
//
// 示例:
//
//  输入：4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释: 4 皇后问题存在如下两个不同的解法。
//[
// [".Q..",  // 解法 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // 解法 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]
//
//
// Related Topics 数组 回溯 👍 190 👎 0

package main

import (
	"fmt"
)

func main() {
	n := 4
	queens := solveNQueens(n)
	fmt.Println(queens)
}

// leetcode submit region begin(Prohibit modification and deletion)
func solveNQueens(n int) [][]string {
	queens, path, buf := make([][]string, 0), make([]int, n), make([]byte, n)
	for i := 0; i < n; i++ { // 复用 buf
		buf[i] = '.'
	}
	queen := func() {
		queue := make([]string, n)
		for i, v := range path {
			buf[v] = 'Q'
			queue[i] = string(buf)
			buf[v] = '.'
		}
		queens = append(queens, queue)
	}
	var dfs func(int, int, int, int)
	dfs = func(row int, c, pie, na int) {
		if row == n { // 一个答案
			queen()
			return
		}
		// 新写法
		v := c | pie | na
		for i := 0; i < n; i++ { // 遍历所有位置
			if j := 1 << i; j&v == 0 { // 尝试选择一个位置（是否可选）
				path[row] = i                          // 记录位置
				dfs(row+1, c|j, (pie|j)<<1, (na|j)>>1) // 撇左移，捺右移
			}
		}
		// 旧写法
		//oks := (1<<n - 1) & ^(c | pie | na) // 可选择的位置
		//for oks != 0 {
		//	last := oks & -oks                              // 选择最小的位置
		//	path[row] = bits.TrailingZeros(uint(last))      // 记录位置
		//	dfs(row+1, c|last, (pie|last)<<1, (na|last)>>1) // 撇左移，捺右移
		//	oks &^= last                                    // 标记已选
		//}
	}
	dfs(0, 0, 0, 0)
	return queens
}

//leetcode submit region end(Prohibit modification and deletion)
