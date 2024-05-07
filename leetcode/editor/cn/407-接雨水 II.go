//给你一个 m x n 的矩阵，其中的值均为非负整数，代表二维高度图每个单元的高度，请计算图中形状最多能接多少体积的雨水。
//
//
//
// 示例 1:
//
//
//
//
//输入: heightMap = [[1,4,3,1,3,2],[3,2,1,3,2,4],[2,3,3,2,3,1]]
//输出: 4
//解释: 下雨后，雨水将会被上图蓝色的方块中。总的接雨水量为1+2+1=4。
//
//
// 示例 2:
//
//
//
//
//输入: heightMap = [[3,3,3,3,3],[3,2,2,2,3],[3,2,1,2,3],[3,2,2,2,3],[3,3,3,3,3]]
//输出: 10
//
//
//
//
// 提示:
//
//
// m == heightMap.length
// n == heightMap[i].length
// 1 <= m, n <= 200
// 0 <= heightMap[i][j] <= 2 * 10⁴
//
//
//
//
// Related Topics 广度优先搜索 数组 矩阵 堆（优先队列） 👍 693 👎 0

package main

import (
	"container/heap"
	"fmt"
)

func main() {
	hm := [][]int{
		{12, 13, 1, 12},
		{13, 4, 13, 12},
		{13, 8, 10, 12},
		{12, 13, 12, 12},
		{13, 13, 13, 13}} // 14
	water := trapRainWater(hm)
	fmt.Println(water)
}

// leetcode submit region begin(Prohibit modification and deletion)
func trapRainWater(heightMap [][]int) int {
	// heap
	ans, m, n := 0, len(heightMap), len(heightMap[0])
	sh := h(make([][3]int, 0, max(0, (m-2)<<1+(n-2)<<1+1)))
	hp := &sh
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	m--
	n--
	for i := 1; i < m; i++ {
		heap.Push(hp, [3]int{i, 0, heightMap[i][0]})
		heap.Push(hp, [3]int{i, n, heightMap[i][n]})
	}
	for j := 1; j < n; j++ {
		heap.Push(hp, [3]int{0, j, heightMap[0][j]})
		heap.Push(hp, [3]int{m, j, heightMap[m][j]})
	}
	for hp.Len() > 0 {
		cur := heap.Pop(hp).([3]int)
		x, y, height := cur[0], cur[1], cur[2]
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx > 0 && nx < m && ny > 0 && ny < n && !vis[nx][ny] {
				ans += max(0, height-heightMap[nx][ny])
				heap.Push(hp, [3]int{nx, ny, max(height, heightMap[nx][ny])})
				vis[nx][ny] = true
			}
		}
	}
	return ans

	// bfs
	//ans, m, n := 0, len(heightMap), len(heightMap[0])
	//q, memo := make([][2]int, 0, max(0, (m-2)<<1+(n-2)<<1+1)), make([][]int, m)
	//dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
	//for i := range memo {
	//	memo[i] = make([]int, n)
	//	for j := range memo[i] {
	//		memo[i][j] = math.MaxInt32
	//	}
	//}
	//for i := 1; i < m-1; i++ {
	//	memo[i][0], memo[i][n-1] = heightMap[i][0], heightMap[i][n-1]
	//	q = append(q, [2]int{i, 0}, [2]int{i, n - 1})
	//}
	//for j := 1; j < n-1; j++ {
	//	memo[0][j], memo[m-1][j] = heightMap[0][j], heightMap[m-1][j]
	//	q = append(q, [2]int{0, j}, [2]int{m - 1, j})
	//}
	//for ; len(q) > 0; q = q[1:] {
	//	x, y := q[0][0], q[0][1]
	//	for i := 0; i < 4; i++ {
	//		nx, ny := x+dx[i], y+dy[i]
	//		if nx > 0 && nx < m-1 && ny > 0 && ny < n-1 && memo[nx][ny] > memo[x][y] {
	//			memo[nx][ny] = max(memo[x][y], heightMap[nx][ny])
	//			q = append(q, [2]int{nx, ny})
	//		}
	//	}
	//}
	//for i := 1; i < m-1; i++ {
	//	for j := 1; j < n-1; j++ {
	//		ans += memo[i][j] - heightMap[i][j]
	//	}
	//}
	//return ans
}

type h [][3]int

func (h h) Len() int {
	return len(h)
}

func (h h) Less(i, j int) bool {
	return h[i][2] < h[j][2]
}

func (h h) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *h) Push(x any) {
	*h = append(*h, x.([3]int))
}

func (h *h) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)

//func trapRainWater(heightMap [][]int) int {
//	// bfs 变种（劣）：参考 https://leetcode.cn/problems/trapping-rain-water-ii/solutions/372754/da-bao-li-san-ge-da-forjie-jun-golang72msji-bai-li/
//	//ret, n, m := 0, len(heightMap), len(heightMap[0])
//	//queue, memo := make([][2]int, 0), make([][]int, n)
//	//for i := 0; i < n; i++ { // 4个角可以不要
//	//	memo[i] = make([]int, m)
//	//	memo[i][0], memo[i][m-1] = heightMap[i][0], heightMap[i][m-1]
//	//}
//	//n--
//	//m--
//	//for j := 1; j < m; j++ {
//	//	memo[0][j], memo[n][j] = heightMap[0][j], heightMap[n][j]
//	//}
//	//for i := 1; i < n; i++ {
//	//	for j := 1; j < m; j++ {
//	//		memo[i][j] = math.MaxInt32
//	//		queue = append(queue, [2]int{i, j}) // 只考虑被围起来的坐标
//	//	}
//	//}
//	//dx, dy := [4]int{0, -1, 0, 1}, [4]int{-1, 0, 1, 0}
//	//for ; len(queue) > 0; queue = queue[1:] {
//	//	x, y := queue[0][0], queue[0][1]
//	//	old := memo[x][y]
//	//	for i := 0; i < 4; i++ {
//	//		if nx, ny := x+dx[i], y+dy[i]; memo[x][y] > memo[nx][ny] { // 是否需要降低水位
//	//			memo[x][y] = memo[nx][ny]
//	//		}
//	//	}
//	//	memo[x][y] = max(memo[x][y], heightMap[x][y])
//	//	if memo[x][y] < old { // 水位降低
//	//		queue = append(queue, [2]int{x, y}) // 需要重新评估
//	//		for i := 0; i < 4; i++ {
//	//			// 降水位后，是否影响周边。如影响，则要重新评估
//	//			if nx, ny := x+dx[i], y+dy[i]; memo[x][y] < memo[nx][ny] && memo[nx][ny] > heightMap[nx][ny] {
//	//				queue = append(queue, [2]int{nx, ny})
//	//			}
//	//		}
//	//	}
//	//}
//	//for i := 1; i < n; i++ {
//	//	for j := 1; j < m; j++ {
//	//		ret += memo[i][j] - heightMap[i][j]
//	//	}
//	//}
//	//return ret
//
//	// bfs
//	var (
//		ret int
//		n   = len(heightMap)
//		m   = len(heightMap[0])
//	)
//	var (
//		dx    = [4]int{0, -1, 0, 1}
//		dy    = [4]int{-1, 0, 1, 0}
//		memo  = make([][]int, n)
//		queue = make([][2]int, 0)
//	)
//	for i := 0; i < n; i++ {
//		memo[i] = make([]int, m)
//		memo[i][0], memo[i][m-1] = heightMap[i][0], heightMap[i][m-1]
//		queue = append(queue, [2]int{i, 0}, [2]int{i, m - 1})
//	}
//	for j := m - 2; j > 0; j-- {
//		memo[0][j], memo[n-1][j] = heightMap[0][j], heightMap[n-1][j]
//		queue = append(queue, [2]int{0, j}, [2]int{n - 1, j})
//	}
//	for i := n - 2; i > 0; i-- {
//		for j := m - 2; j > 0; j-- {
//			memo[i][j] = math.MaxInt32
//		}
//	}
//	for ; len(queue) > 0; queue = queue[1:] {
//		x, y := queue[0][0], queue[0][1]
//		for i := 0; i < 4; i++ {
//			//if nx, ny := x+dx[i], y+dy[i]; 0 <= nx && nx < n && 0 <= ny && ny < m &&
//			if nx, ny := x+dx[i], y+dy[i]; 0 < nx && nx < n-1 && 0 < ny && ny < m-1 &&
//				memo[nx][ny] > memo[x][y] { // 主要逻辑：漏水
//				memo[nx][ny] = max(heightMap[nx][ny], memo[x][y]) // 降低蓄水量
//				queue = append(queue, [2]int{nx, ny})
//			}
//		}
//	}
//	for i := n - 2; i > 0; i-- {
//		for j := m - 2; j > 0; j-- {
//			ret += memo[i][j] - heightMap[i][j]
//		}
//	}
//	return ret
//
//	// 堆
//	//ret, m, n := 0, len(heightMap)-1, len(heightMap[0])-1
//	//if m < 2 || n < 2 {
//	//	return 0
//	//}
//	//h, visited := &trwHp{}, make([][]bool, m+1)
//	//d := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
//	//for i := 0; i <= m; i++ {
//	//	visited[i] = make([]bool, n+1)
//	//	heap.Push(h, [3]int{i, 0, heightMap[i][0]}) // 行的两侧
//	//	heap.Push(h, [3]int{i, n, heightMap[i][n]})
//	//	visited[i][0], visited[i][n] = true, true // 已访问
//	//	if i == 0 || i == m {
//	//		for j := 1; j < n; j++ {
//	//			heap.Push(h, [3]int{i, j, heightMap[i][j]}) // 行的中间
//	//			visited[i][j] = true
//	//		}
//	//	}
//	//}
//	//for h.Len() > 0 {
//	//	top := heap.Pop(h).([3]int) // 当前最小高
//	//	x, y := top[0], top[1]
//	//	for i := 0; i < 4; i++ { // 四连通
//	//		nx, ny := x+d[i][0], y+d[i][1]
//	//		if nx >= 0 && nx <= m && ny >= 0 && ny <= n && !visited[nx][ny] {
//	//			visited[nx][ny] = true
//	//			nh := heightMap[nx][ny]
//	//			if top[2]-nh > 0 { // 可积攒雨水
//	//				ret += top[2] - nh // 雨水量
//	//				nh = top[2]        // 更新高
//	//			}
//	//			heap.Push(h, [3]int{nx, ny, nh}) // 入堆
//	//		}
//	//	}
//	//}
//	//return ret
//}
//
//type hp407 [][3]int
//
//func (h hp407) Len() int           { return len(h) }
//func (h hp407) Less(i, j int) bool { return h[i][2] < h[j][2] }
//func (h hp407) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//func (h *hp407) Push(x any)        { *h = append(*h, x.([3]int)) }
//func (h *hp407) Pop() any {
//	v := (*h)[len(*h)-1]
//	*h = (*h)[:len(*h)-1]
//	return v
//}
