package main

import "fmt"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minCost(grid [][]int) int {
	// lc：0-1 贪心广度优先搜索 TODO
	// 用 list 包的链表实现双端队列，如果与当前方向相符，则插入队列前，如果相反，则插入队列后

	// bfs：个人
	// 右 左 下 上
	dx, dy := [4]int{0, 0, 1, -1}, [4]int{1, -1, 0, 0}
	//ans, m, n := 0, len(grid), len(grid[0])
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	vis[0][0] = true
	for q, cost := [][2]int{{}}, 0; ; cost++ { // 每一轮 bfs 就是一次 cost 修改
		for _, cur := range q { // 1.尝试顺着方向走
			x, y := cur[0], cur[1]
			for { // 一直顺着走
				k := grid[x][y] - 1 // 1-4 改为 0-3
				if x, y = x+dx[k], y+dy[k]; x < 0 || x == m || y < 0 || y == n || vis[x][y] {
					break // 该路径已走不通
				}
				q = append(q, [2]int{x, y})
				vis[x][y] = true
			}
		}
		l := len(q)
		for _, cur := range q[:l] { // 2.尝试修改方向走
			x, y := cur[0], cur[1]
			if x == m-1 && y == n-1 { // 终止条件
				return cost
			}
			for k := 0; k < 4; k++ { // 修改方向
				i, j := x+dx[k], y+dy[k]
				if i >= 0 && i < m && j >= 0 && j < n && !vis[i][j] {
					q = append(q, [2]int{i, j})
					vis[i][j] = true
				}
			}
		}
		q = q[l:]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
