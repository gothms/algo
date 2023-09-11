//你准备参加一场远足活动。给你一个二维 rows x columns 的地图 heights ，其中 heights[row][col] 表示格子 (row,
// col) 的高度。一开始你在最左上角的格子 (0, 0) ，且你希望去最右下角的格子 (rows-1, columns-1) （注意下标从 0 开始编号）。你
//每次可以往 上，下，左，右 四个方向之一移动，你想要找到耗费 体力 最小的一条路径。
//
// 一条路径耗费的 体力值 是路径上相邻格子之间 高度差绝对值 的 最大值 决定的。
//
// 请你返回从左上角走到右下角的最小 体力消耗值 。
//
//
//
// 示例 1：
//
//
//
//
//输入：heights = [[1,2,2],[3,8,2],[5,3,5]]
//输出：2
//解释：路径 [1,3,5,3,5] 连续格子的差值绝对值最大为 2 。
//这条路径比路径 [1,2,2,2,5] 更优，因为另一条路径差值最大值为 3 。
//
//
// 示例 2：
//
//
//
//
//输入：heights = [[1,2,3],[3,8,4],[5,3,5]]
//输出：1
//解释：路径 [1,2,3,4,5] 的相邻格子差值绝对值最大为 1 ，比路径 [1,3,5,3,5] 更优。
//
//
// 示例 3：
//
//
//输入：heights = [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
//输出：0
//解释：上图所示路径不需要消耗任何体力。
//
//
//
//
// 提示：
//
//
// rows == heights.length
// columns == heights[i].length
// 1 <= rows, columns <= 100
// 1 <= heights[i][j] <= 10⁶
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 二分查找 矩阵 堆（优先队列） 👍 373 👎 0

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	path := minimumEffortPath(heights)
	fmt.Println(path)
}

// leetcode submit region begin(Prohibit modification and deletion)

func minimumEffortPath(heights [][]int) int {
	// Dijkstra
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	d := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	n, m := len(heights), len(heights[0])
	maxDiff := make([][]int, n) // 记录每个 point 的体力消耗
	for i := 0; i < n; i++ {
		maxDiff[i] = make([]int, m)
		for j := 0; j < m; j++ {
			maxDiff[i][j] = 1e6 // 初始化为最大值，如果被访问过，则新的访问必须更优（起到了记录已访问作用）
		}
	}
	maxDiff[0][0] = 0 // 记录已访问
	n--
	m--
	h := &point{{}}
	for h.Len() > 0 { // [0,0]
		p := heap.Pop(h).([3]int)
		if p[0] == n && p[1] == m { // 终点
			return p[2]
		}
		//if maxDiff[p[0]][p[1]] < p[2] { // maxDiff[p[0]][p[1]] 就是 p[2]，所以判断无效（因为它们是同步更新的）
		//	continue
		//}
		for _, v := range d {
			if i, j := p[0]+v[0], p[1]+v[1]; i >= 0 && i <= n && j >= 0 && j <= m {
				// 需要记录路径上最大高差，所以要通过 maxVal 把 p[2] 传递下去
				// diff == maxDiff[i][j] 是来的路径，所以不会走回头路
				// 如果 maxDiff[i][j] 已被访问过，则新的访问必须更优
				if diff := maxVal(p[2], abs(heights[p[0]][p[1]]-heights[i][j])); diff < maxDiff[i][j] {
					maxDiff[i][j] = diff // 和 heap 中记录的 diff 是同步更新的
					heap.Push(h, [3]int{i, j, diff})
				}
			}
		}
	}
	return -1

	// 并查集
	// 二分查找

	// lc
	//n, m := len(heights), len(heights[0])
	//maxDiff := make([][]int, n)
	//for i := range maxDiff {
	//	maxDiff[i] = make([]int, m)
	//	for j := range maxDiff[i] {
	//		maxDiff[i][j] = math.MaxInt64
	//	}
	//}
	//maxDiff[0][0] = 0
	//h := &hp{{}}
	//for {
	//	p := heap.Pop(h).(point)
	//	if p.x == n-1 && p.y == m-1 {
	//		return p.maxDiff
	//	}
	//	if maxDiff[p.x][p.y] < p.maxDiff {	//
	//		continue
	//	}
	//	for _, d := range dir4 {
	//		if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < n && 0 <= y && y < m {
	//			if diff := max(p.maxDiff, abs(heights[x][y]-heights[p.x][p.y])); diff < maxDiff[x][y] {
	//				maxDiff[x][y] = diff
	//				heap.Push(h, point{x, y, diff})
	//			}
	//		}
	//	}
	//	//fmt.Println(h)
	//	//fmt.Println(maxDiff)
	//}
}

type point [][3]int

func (p point) Len() int {
	return len(p)
}

func (p point) Less(i, j int) bool {
	return p[i][2] < p[j][2]
}

func (p point) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *point) Push(x any) {
	*p = append(*p, x.([3]int))
}

func (p *point) Pop() any {
	v := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
