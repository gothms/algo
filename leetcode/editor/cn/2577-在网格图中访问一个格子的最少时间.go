//给你一个 m x n 的矩阵 grid ，每个元素都为 非负 整数，其中 grid[row][col] 表示可以访问格子 (row, col) 的 最早 时
//间。也就是说当你访问格子 (row, col) 时，最少已经经过的时间为 grid[row][col] 。
//
// 你从 最左上角 出发，出发时刻为 0 ，你必须一直移动到上下左右相邻四个格子中的 任意 一个格子（即不能停留在格子上）。每次移动都需要花费 1 单位时间。
//
//
// 请你返回 最早 到达右下角格子的时间，如果你无法到达右下角的格子，请你返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：grid = [[0,1,3,2],[5,1,2,5],[4,3,8,6]]
//输出：7
//解释：一条可行的路径为：
//- 时刻 t = 0 ，我们在格子 (0,0) 。
//- 时刻 t = 1 ，我们移动到格子 (0,1) ，可以移动的原因是 grid[0][1] <= 1 。
//- 时刻 t = 2 ，我们移动到格子 (1,1) ，可以移动的原因是 grid[1][1] <= 2 。
//- 时刻 t = 3 ，我们移动到格子 (1,2) ，可以移动的原因是 grid[1][2] <= 3 。
//- 时刻 t = 4 ，我们移动到格子 (1,1) ，可以移动的原因是 grid[1][1] <= 4 。
//- 时刻 t = 5 ，我们移动到格子 (1,2) ，可以移动的原因是 grid[1][2] <= 5 。
//- 时刻 t = 6 ，我们移动到格子 (1,3) ，可以移动的原因是 grid[1][3] <= 6 。
//- 时刻 t = 7 ，我们移动到格子 (2,3) ，可以移动的原因是 grid[2][3] <= 7 。
//最终到达时刻为 7 。这是最早可以到达的时间。
//
//
// 示例 2：
//
//
//
//
//输入：grid = [[0,2,4],[3,2,1],[1,0,4]]
//输出：-1
//解释：没法从左上角按题目规定走到右下角。
//
//
//
//
// 提示：
//
//
// m == grid.length
// n == grid[i].length
// 2 <= m, n <= 1000
// 4 <= m * n <= 10⁵
// 0 <= grid[i][j] <= 10⁵
// grid[0][0] == 0
//
//
// Related Topics 广度优先搜索 图 数组 矩阵 最短路 堆（优先队列） 👍 33 👎 0

package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {

}

/*
重点
	1.到 x,y 的步数，必然和 x+y 的奇偶性相同
		nd += (nd - x - y) & 1
	2.更新到 x,y 的步数
		memo
	3.2 个以上的可活动空间
		if grid[0][1] > 1 && grid[1][0] > 1
	4.跳跃：在 2 个活动空间前提下，去掉 grid[x][y] <= q[2]+1 判断
		而使用 nd := maxVal(q[2]+1, grid[x][y])
	5.出队列再判断 if q[0] == m && q[1] == n
	6.避免重复访问(模仿了：图)
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minimumTime(grid [][]int) int {
	// dijkstra
	if grid[0][1] > 1 && grid[1][0] > 1 { // 没有路，否则 2 个以上的可活动空间
		return -1
	}
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	m, n := len(grid), len(grid[0])
	h, memo := &mtHp{{}}, make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = math.MaxInt32
		}
	}
	memo[0][0] = 0
	m--
	n--
	d := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	for h.Len() > 0 {
		q := heap.Pop(h).([3]int)
		if q[2] > memo[q[0]][q[1]] { // 剪枝
			continue
		}
		if q[0] == m && q[1] == n { // 出队列再判断
			return q[2]
		}
		x, y := 0, 0
		for i := 0; i < 4; i++ {
			x, y = q[0]+d[i][0], q[1]+d[i][1]
			if x >= 0 && x <= m && y >= 0 && y <= n {
				nd := maxVal(q[2]+1, grid[x][y]) // 跳跃：在 2 个活动空间前提下，去掉 grid[x][y] <= q[2]+1 判断
				nd += (nd - x - y) & 1
				if nd < memo[x][y] { // 避免了回溯
					memo[x][y] = nd // 更新到 x,y 的步数
					heap.Push(h, [3]int{x, y, nd})
				}
			}
		}
		fmt.Println(h)
	}
	return -1

	// runtime: out of memory: cannot allocate 113246208-byte block (440139776 in use)
	//cnt, m, n := 0, len(grid)-1, len(grid[0])-1
	//queue, d := [][2]int{{}}, [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	//for l := 1; l > 0; l = len(queue) {
	//	cnt++
	//	for _, q := range queue {
	//		x, y := 0, 0
	//		for i := 0; i < 4; i++ {
	//			x, y = q[0]+d[i][0], q[1]+d[i][1]
	//			if x >= 0 && x <= m && y >= 0 && y <= n && grid[x][y] <= cnt {
	//				if x == m && y == n {
	//					return cnt
	//				}
	//				queue = append(queue, [2]int{x, y})
	//			}
	//		}
	//	}
	//	queue = queue[l:]
	//}
	//return -1
}

type mtHp [][3]int

func (m mtHp) Len() int           { return len(m) }
func (m mtHp) Less(i, j int) bool { return m[i][2] < m[j][2] }
func (m mtHp) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m *mtHp) Push(x any)        { *m = append(*m, x.([3]int)) }
func (m *mtHp) Pop() any {
	v := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
