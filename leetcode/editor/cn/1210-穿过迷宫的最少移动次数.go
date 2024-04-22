//你还记得那条风靡全球的贪吃蛇吗？
//
// 我们在一个 n*n 的网格上构建了新的迷宫地图，蛇的长度为 2，也就是说它会占去两个单元格。蛇会从左上角（(0, 0) 和 (0, 1)）开始移动。我们用
// 0 表示空单元格，用 1 表示障碍物。蛇需要移动到迷宫的右下角（(n-1, n-2) 和 (n-1, n-1)）。
//
// 每次移动，蛇可以这样走：
//
//
// 如果没有障碍，则向右移动一个单元格。并仍然保持身体的水平／竖直状态。
// 如果没有障碍，则向下移动一个单元格。并仍然保持身体的水平／竖直状态。
// 如果它处于水平状态并且其下面的两个单元都是空的，就顺时针旋转 90 度。蛇从（(r, c)、(r, c+1)）移动到 （(r, c)、(r+1, c)）。
//
// 如果它处于竖直状态并且其右面的两个单元都是空的，就逆时针旋转 90 度。蛇从（(r, c)、(r+1, c)）移动到（(r, c)、(r, c+1)）。
//
//
//
// 返回蛇抵达目的地所需的最少移动次数。
//
// 如果无法到达目的地，请返回 -1。
//
//
//
// 示例 1：
//
//
//
// 输入：grid = [[0,0,0,0,0,1],
//               [1,1,0,0,1,0],
//               [0,0,0,0,1,1],
//               [0,0,1,0,1,0],
//               [0,1,1,0,0,0],
//               [0,1,1,0,0,0]]
//输出：11
//解释：
//一种可能的解决方案是 [右, 右, 顺时针旋转, 右, 下, 下, 下, 下, 逆时针旋转, 右, 下]。
//
//
// 示例 2：
//
// 输入：grid = [[0,0,1,1,1,1],
//               [0,0,0,0,1,1],
//               [1,1,0,0,0,1],
//               [1,1,1,0,0,1],
//               [1,1,1,0,0,1],
//               [1,1,1,0,0,0]]
//输出：9
//
//
//
//
// 提示：
//
//
// 2 <= n <= 100
// 0 <= grid[i][j] <= 1
// 蛇保证从空单元格开始出发。
//
//
// Related Topics 广度优先搜索 数组 矩阵 👍 145 👎 0

package main

import "fmt"

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	moves := minimumMoves(grid)
	fmt.Println(moves)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumMoves(grid [][]int) int {
	// lc
	type tuple struct{ x, y, s int }
	var dirs = []tuple{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	n := len(grid)
	vis := make([][][2]bool, n)
	for i := range vis {
		vis[i] = make([][2]bool, n)
	}
	vis[0][0][0] = true // 初始位置
	q := []tuple{{}}
	for step, l := 1, 1; l > 0; step, l = step+1, len(q) {
		for _, t := range q[:l] {
			for _, d := range dirs {
				x, y, s := t.x+d.x, t.y+d.y, t.s^d.s
				x2, y2 := x+s, y+(s^1) // 蛇头
				if x2 < n && y2 < n && !vis[x][y][s] &&
					grid[x][y] == 0 && grid[x2][y2] == 0 && (d.s == 0 || grid[x+1][y+1] == 0) {
					if x == n-1 && y == n-2 { // 此时蛇头一定在 (n-1,n-1)
						return step
					}
					vis[x][y][s] = true
					q = append(q, tuple{x, y, s})
				}
			}
		}
		q = q[l:]
	}
	return -1

	// bfs：超时原因
	//在 bfs 算法中，先出队列的肯定 优于 后出队列的
	//所以不用记录 最优值，或判断出队列的值是否比已记录的值 更优
	//const e, s = 0, 1 // 向右，向下
	//n := len(grid)
	//memo := make([][][2]bool, n) // 标记已访问
	//for i := range memo {
	//	memo[i] = make([][2]bool, n) // 两个方向都要标记
	//}
	//q := [][4]int{{0, 0, e, 0}} // 不可能返回，所以不用初始 [2]bool{ture, true}
	//redirect := func(x, y, m, d int) {
	//	if !memo[x][y][m^1] {
	//		q = append(q, [4]int{x, y, m ^ 1, d})
	//		memo[x][y][m^1] = true
	//	}
	//}
	//move := func(x, y, m, d int) {
	//	if !memo[x][y][m] {
	//		q = append(q, [4]int{x, y, m, d})
	//		memo[x][y][m] = true
	//	}
	//}
	//for ; len(q) > 0; q = q[1:] {
	//	x, y, m, d := q[0][0], q[0][1], q[0][2], q[0][3]
	//	//	if x == n-1 && y == n-2 && m == e {
	//	if x == n-1 && y == n-2 { // 蛇头一定在 [n-1][n-1]
	//		return d
	//	}
	//	d++ // 更新距离
	//	switch m {
	//	case e: // 横向
	//		if y+2 < n && grid[x][y+2] == 0 { // 沿着朝向前进
	//			move(x, y+1, m, d)
	//		}
	//		if x+1 < n && grid[x+1][y+1] == 0 && grid[x+1][y] == 0 { // 转向：grid[x+1][y+1] == 1 时，不能转向
	//			redirect(x, y, m, d) // 转向
	//			move(x+1, y, m, d)   // 垂直于朝向前进
	//		}
	//	case s: // 竖向
	//		if x+2 < n && grid[x+2][y] == 0 {
	//			move(x+1, y, m, d)
	//		}
	//		if y+1 < n && grid[x+1][y+1] == 0 && grid[x][y+1] == 0 {
	//			redirect(x, y, m, d)
	//			move(x, y+1, m, d)
	//		}
	//	}
	//}
	//return -1

	// bfs
	//const e, s = 0, 1 // 向右，向下
	//n := len(grid)
	//memo := make([][][2]bool, n)
	//for i := range memo {
	//	memo[i] = make([][2]bool, n)
	//}
	//for q := [][4]int{{0, 0, e, 0}}; len(q) > 0; q = q[1:] {
	//	x, y, m, d := q[0][0], q[0][1], q[0][2], q[0][3]
	//	if x == n-1 && y == n-2 && m == e {
	//		return d
	//	}
	//	d++
	//	switch m {
	//	case e: // 横向
	//		if y+2 < n && grid[x][y+2] == 0 && !memo[x][y+1][m] {
	//			q = append(q, [4]int{x, y + 1, m, d}) // 右移一格
	//			memo[x][y+1][m] = true
	//		}
	//		if x+1 < n && grid[x+1][y+1] == 0 && grid[x+1][y] == 0 {
	//			if !memo[x][y][m^1] {
	//				q = append(q, [4]int{x, y, m ^ 1, d}) // 转向：grid[x+1][y+1] == 1 时，不能转向
	//				memo[x][y][m^1] = true
	//			}
	//			if !memo[x+1][y][m] {
	//				q = append(q, [4]int{x + 1, y, m, d}) // 下移
	//				memo[x+1][y][m] = true
	//			}
	//		}
	//	case s: // 竖向
	//		if x+2 < n && grid[x+2][y] == 0 && !memo[x+1][y][m] {
	//			q = append(q, [4]int{x + 1, y, m, d}) // 下移一格
	//			memo[x+1][y][m] = true
	//		}
	//		if y+1 < n && grid[x+1][y+1] == 0 && grid[x][y+1] == 0 {
	//			if !memo[x][y][m^1] {
	//				q = append(q, [4]int{x, y, m ^ 1, d}) // 转向
	//				memo[x][y][m^1] = true
	//			}
	//			if !memo[x][y+1][m] {
	//				q = append(q, [4]int{x, y + 1, m, d}) // 右移
	//				memo[x][y+1][m] = true
	//			}
	//		}
	//	}
	//}
	//return -1
}

//leetcode submit region end(Prohibit modification and deletion)
