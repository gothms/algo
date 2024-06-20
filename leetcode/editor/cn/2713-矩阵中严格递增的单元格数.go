package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	mat := [][]int{
		{3, 1, 6},
		{-9, 5, 7},
	}
	cells := maxIncreasingCells(mat)
	fmt.Println(cells)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxIncreasingCells(mat [][]int) int {
	// dp：最长递增子序列
	type key struct {
		v, c int
		pc   int // 示例：假设同一行的子序列为 2,3,3 时，第二个 3 应该由 2 移动到达，而 pc 记录的就是 2 的最长递增子序列
	}
	m, n := len(mat), len(mat[0])
	row, col := make([]key, m), make([]key, n)
	vs := make([][3]int, 0, m*n)
	for i, r := range mat {
		for j, v := range r {
			vs = append(vs, [3]int{i, j, v})
		}
	}
	sort.Slice(vs, func(i, j int) bool {
		return vs[i][2] < vs[j][2]
	})
	for _, s := range vs {
		x, y, v := s[0], s[1], s[2]
		xc, yc := row[x].c, col[y].c
		if v == row[x].v { // 同一行有值相等的单元格
			xc = row[x].pc
		}
		if v == col[y].v { // 同一列有值相等的单元格
			yc = col[y].pc
		}
		c := max(xc, yc) + 1 // 当前最长子序列
		if c > row[x].c {    // 更新
			if v == row[x].v {
				row[x].c = c // 值有重复
			} else {
				row[x].v, row[x].c, row[x].pc = v, c, row[x].c // 值没重复
			}
		}
		if c > col[y].c {
			if v == col[y].v {
				col[y].c = c
			} else {
				col[y].v, col[y].c, col[y].pc = v, c, col[y].c
			}
		}
	}
	return slices.MaxFunc(row, func(a, b key) int { // 返回最长递增子序列长度
		return a.c - b.c
	}).c

	// lc
	//type pair struct{ x, y int }
	//memo := make(map[int][]pair)
	//for i, row := range mat {
	//	for j, v := range row {
	//		memo[v] = append(memo[v], pair{i, j})
	//	}
	//}
	//keys := make([]int, 0, len(memo))
	//for key := range memo {
	//	keys = append(keys, key)
	//}
	//sort.Ints(keys)
	//row, col := make([]int, len(mat)), make([]int, len(mat[0]))
	//for _, key := range keys {
	//	ps := memo[key]
	//	cur := make([]int, 0, len(ps))
	//	for _, p := range ps {
	//		cur = append(cur, max(row[p.x], col[p.y])+1)
	//	}
	//	for i, p := range ps {
	//		row[p.x] = max(row[p.x], cur[i])
	//		col[p.y] = max(col[p.y], cur[i])
	//	}
	//}
	//return slices.Max(row)
}

//leetcode submit region end(Prohibit modification and deletion)
