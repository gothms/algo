package main

import "fmt"

func main() {
	grid := [][]byte{{'a'}, {'a'}, {'b'}}
	cycle := containsCycle(grid)
	fmt.Println(cycle)
	fmt.Println('a', 'A')
}

// leetcode submit region begin(Prohibit modification and deletion)
func containsCycle(grid [][]byte) bool {
	// bfs
	m, n := len(grid), len(grid[0])
	dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	for i, g := range grid {
		for j := range g {
			if vis[i][j] {
				continue
			}
			vis[i][j] = true
			q := [][2]int{{i, j}}
			for len(q) > 0 {
				neibors, size := 0, len(q)
				x, y := q[0][0], q[0][1]
				for k := 0; k < 4; k++ {
					nx, ny := x+dx[k], y+dy[k]
					if nx >= 0 && nx < m && ny >= 0 && ny < n && grid[nx][ny] == grid[x][y] {
						neibors++
						if !vis[nx][ny] {
							vis[nx][ny] = true
							q = append(q, [2]int{nx, ny})
						}
					}
				}
				if neibors-1 > len(q)-size { // 核心算法
					return true
				}
				q = q[1:]
			}
		}
	}
	return false

	// 并查集
	//m, n := len(grid), len(grid[0])
	//pa, size := make([]int, m*n), make([]int, m*n)
	//for i := range pa {
	//	pa[i], size[i] = i, 1
	//}
	//var find func(x int) int
	//find = func(x int) int {
	//	if pa[x] != x {
	//		pa[x] = find(pa[x])
	//	}
	//	return pa[x]
	//}
	//unite := func(x, y int) bool {
	//	x, y = find(x), find(y)
	//	if x == y {
	//		return false
	//	}
	//	if x < y {
	//		x, y = y, x
	//	}
	//	pa[y] = x
	//	size[x] += size[y]
	//	return true
	//}
	//for i, g := range grid {
	//	for j, v := range g {
	//		if i > 0 && grid[i-1][j] == v {	// 核心算法
	//			if !unite(i*n+j, (i-1)*n+j) {
	//				return true
	//			}
	//		}
	//		if j > 0 && grid[i][j-1] == v {
	//			if !unite(i*n+j, i*n+j-1) {
	//				return true
	//			}
	//		}
	//	}
	//}
	//return false

	// dfs
	//m, n := len(grid), len(grid[0])
	//memo := make([][]int, m)
	//for i := range memo {
	//	memo[i] = make([]int, n)
	//}
	//dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	//var cur byte
	//var dfs func(int, int, int) bool
	//dfs = func(x, y, s int) bool {
	//	s++ // 下一步
	//	for i := 0; i < 4; i++ {
	//		if nx, ny := x+dx[i], y+dy[i]; nx >= 0 && nx < m && ny >= 0 && ny < n &&
	//			grid[nx][ny] == cur {
	//			if memo[nx][ny] > 0 { // ==0 则未访问
	//				if memo[nx][ny]+4 <= s { // 最少 4 个格子
	//					return true
	//				}
	//			} else {
	//				memo[nx][ny] = s
	//				if dfs(nx, ny, s) {
	//					return true
	//				}
	//			}
	//		}
	//	}
	//	return false
	//}
	//for i, g := range grid {
	//	for j, c := range g {
	//		if c == 0 || memo[i][j] > 0 {
	//			continue
	//		}
	//		cur = c // 当前字符
	//		memo[i][j] = 1
	//		if dfs(i, j, 1) {
	//			return true
	//		}
	//	}
	//}
	//return false
}

//leetcode submit region end(Prohibit modification and deletion)
