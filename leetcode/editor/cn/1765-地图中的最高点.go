package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func highestPeak(isWater [][]int) [][]int {
	// bfs：优化
	// 入队列的 (i,j) 一定是当前最高的
	// 同样，出队列的 (i,j) 一定是当前最低的，则此时 (i,j) 扩散出去的高度就是正确的
	m, n := len(isWater), len(isWater[0])
	const D = 4
	ans, q := make([][]int, m), make([][3]int, 0)
	dx, dy := [D]int{0, -1, 0, 1}, [D]int{-1, 0, 1, 0}
	cnt := m * n
	for i, r := range isWater {
		ans[i] = make([]int, n)
		for j, v := range r {
			if v == 1 {
				ans[i][j] = 0
				isWater[i][j] = -1 // 已访问
				q = append(q, [3]int{i, j, 1})
				cnt--
			}
		}
	}
	for len(q) > 0 {
		x, y, h := q[0][0], q[0][1], q[0][2]
		for i := 0; i < D; i++ {
			if nx, ny := x+dx[i], y+dy[i]; nx >= 0 && nx < m && ny >= 0 && ny < n &&
				isWater[nx][ny] != -1 {
				ans[nx][ny] = h // 设置高度
				if cnt--; cnt == 0 {
					return ans // 已访问完
				}
				isWater[nx][ny] = -1
				q = append(q, [3]int{nx, ny, h + 1})
			}
		}
		q = q[1:]
	}
	return ans

	// bfs：个人，笨拙
	//const D = 4
	//m, n := len(isWater), len(isWater[0])
	//ans, q := make([][]int, m), make([][2]int, 0)
	//dx, dy := [D]int{0, -1, 0, 1}, [D]int{-1, 0, 1, 0}
	//for i := range ans {
	//	ans[i] = make([]int, n)
	//	for j := range ans[i] {
	//		ans[i][j] = -1
	//	}
	//}
	//fill := func(x, y int) {
	//	for i := 0; i < D; i++ {
	//		if nx, ny := x+dx[i], y+dy[i]; nx >= 0 && nx < m && ny >= 0 && ny < n &&
	//			ans[nx][ny] == -1 {
	//			v := math.MaxInt32
	//			for j := 0; j < D; j++ {
	//				if cx, cy := nx+dx[j], ny+dy[j]; cx >= 0 && cx < m && cy >= 0 && cy < n &&
	//					ans[cx][cy] != -1 {
	//					v = min(v, ans[cx][cy])
	//				}
	//			}
	//			ans[nx][ny] = v + 1
	//			q = append(q, [2]int{nx, ny})
	//		}
	//	}
	//}
	//for i, r := range isWater {
	//	for j, v := range r {
	//		if v == 1 {
	//			ans[i][j] = 0
	//			fill(i, j)
	//		}
	//	}
	//}
	//for l := len(q); l > 0; l = len(q) {
	//	for i := 0; i < l; i++ {
	//		fill(q[i][0], q[i][1])
	//	}
	//	q = q[l:]
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
