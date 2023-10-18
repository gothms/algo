package main

import "math/bits"

/*
N 皇后
	51
	52
	面试题 08.12
*/

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
		//v := c | pie | na
		//for i := 0; i < n; i++ { // 遍历所有位置
		//	if j := 1 << i; j&v == 0 { // 尝试选择一个位置（是否可选）
		//		path[row] = i                          // 记录位置
		//		dfs(row+1, c|j, (pie|j)<<1, (na|j)>>1) // 撇左移，捺右移
		//	}
		//}
		// 旧写法
		oks := (1<<n - 1) & ^(c | pie | na) // 可选择的位置
		for oks != 0 {
			last := oks & -oks                              // 选择最小的位置
			path[row] = bits.TrailingZeros(uint(last))      // 记录位置
			dfs(row+1, c|last, (pie|last)<<1, (na|last)>>1) // 撇左移，捺右移
			oks &^= last                                    // 标记已选
		}
	}
	dfs(0, 0, 0, 0)
	return queens
}

func totalNQueens(n int) int {
	cnt, bitN := 0, 1<<n-1
	var dfs func(int, int, int, int)
	dfs = func(row, c, pie, na int) {
		if row == n {
			cnt++
			return
		}
		oks := bitN & ^(c | pie | na)
		for last := 0; oks != 0; oks &^= last {
			last = oks & -oks
			dfs(row+1, c|last, (pie|last)<<1, (na|last)>>1)
		}
	}
	dfs(0, 0, 0, 0)
	return cnt
}
