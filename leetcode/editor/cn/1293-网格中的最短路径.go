package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 0},
		{1, 1, 0},
		{0, 0, 0},
		{0, 1, 1},
		{0, 0, 0}}
	k := 1 // 6
	path := shortestPath(grid, k)
	fmt.Println(path)
}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestPath(grid [][]int, k int) int {

}

//leetcode submit region end(Prohibit modification and deletion)
//func shortestPath(grid [][]int, k int) int {
//	// bfs：优化
//	m, n := len(grid), len(grid[0])
//	if m|n == 1 { // fast path
//		return 0
//	}
//	k = min(k, m+n-3)                // 虽然 1 <= k <= m*n，但最多可能消除 m+n-3 个障碍物，即可到达终点
//	vis := make([][]map[int]bool, m) // 充分利用 bfs 的性质，则用 vis 来记录格子和消除的障碍物的访问情况
//	for i := range vis {
//		vis[i] = make([]map[int]bool, n)
//		for j := range vis[i] {
//			vis[i][j] = make(map[int]bool)
//		}
//	}
//	vis[0][0][0] = true
//	dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
//	q := [][4]int{{}} // 放入 [0 0 0 0]
//	for ; len(q) > 0; q = q[1:] {
//		cur := q[0]
//		x, y, l, v := cur[0], cur[1], cur[2], cur[3]
//		for i := 0; i < 4; i++ {
//			nx, ny := x+dx[i], y+dy[i]
//			if nx < 0 || nx == m || ny < 0 || ny == n {
//				continue
//			}
//			switch grid[nx][ny] {
//			case 0:
//				if nx == m-1 && ny == n-1 { // (m-1,n-1) = 0
//					return v + 1
//				}
//				if !vis[nx][ny][l] {
//					vis[nx][ny][l] = true
//					q = append(q, [4]int{nx, ny, l, v + 1}) // bfs：最先到达的，一定是最短路径
//				}
//			case 1:
//				if l < k && !vis[nx][ny][l+1] { // 必须 l < k
//					vis[nx][ny][l+1] = true
//					q = append(q, [4]int{nx, ny, l + 1, v + 1})
//				}
//			}
//		}
//	}
//	return -1
//
//	// bfs：个人
//	// 执行耗时:87 ms,击败了6.90% 的Go用户
//	// 内存消耗:8.8 MB,击败了6.90% 的Go用户
//	//const inf = 1600
//	//m, n := len(grid), len(grid[0])
//	//memo := make([][]map[int]int, m)
//	//for i := range memo {
//	//	memo[i] = make([]map[int]int, n)
//	//	for j := range memo[i] {
//	//		memo[i][j] = make(map[int]int)
//	//	}
//	//}
//	//memo[0][0][0] = 0
//	//dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
//	//q := [][4]int{{}}
//	//for ; len(q) > 0; q = q[1:] {
//	//	cur := q[0]
//	//	x, y, l, v := cur[0], cur[1], cur[2], cur[3]
//	//	for i := 0; i < 4; i++ {
//	//		nx, ny := x+dx[i], y+dy[i]
//	//		if nx < 0 || nx == m || ny < 0 || ny == n {
//	//			continue
//	//		}
//	//		if nx == m-1 && ny == n-1 {
//	//			return v + 1
//	//		}
//	//		switch grid[nx][ny] {
//	//		case 0:
//	//			if nv, ok := memo[nx][ny][l]; !ok || nv > v+1 {
//	//				memo[nx][ny][l] = v + 1
//	//				q = append(q, [4]int{nx, ny, l, v + 1})
//	//			}
//	//		case 1:
//	//			if l == k {
//	//				break
//	//			}
//	//			if nv, ok := memo[nx][ny][l+1]; !ok || nv > v+1 {
//	//				memo[nx][ny][l+1] = v + 1
//	//				q = append(q, [4]int{nx, ny, l + 1, v + 1})
//	//			}
//	//		}
//	//	}
//	//}
//	//ans := inf
//	//for _, v := range memo[m-1][n-1] {
//	//	ans = min(ans, v)
//	//}
//	//if ans == inf {
//	//	return -1
//	//}
//	//return -1
//}
