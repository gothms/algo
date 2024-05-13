package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func orangesRotting(grid [][]int) int {
	// bfs
	ans, cnt, m, n := 0, 0, len(grid), len(grid[0])
	q := make([][2]int, 0)
	dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	for i, g := range grid {
		for j, v := range g {
			switch v {
			case 1:
				cnt++
			case 2:
				q = append(q, [2]int{i, j})
			}
		}
	}
	if cnt == 0 {
		return ans
	}
	for l := len(q); l > 0; l = len(q) { // 使用 temp 可以更少扩容？
		for i := 0; i < l; i++ {
			x, y := q[i][0], q[i][1]
			for k := 0; k < 4; k++ {
				if nx, ny := x+dx[k], y+dy[k]; nx >= 0 && nx < m && ny >= 0 && ny < n &&
					grid[nx][ny] == 1 {
					cnt--
					grid[nx][ny] = 2
					q = append(q, [2]int{nx, ny})
				}
			}
		}
		ans++
		if cnt == 0 {
			return ans
		}
		q = q[l:]
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
