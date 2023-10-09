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

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func trapRainWater(heightMap [][]int) int {
	// bfs
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	ret, m, n := 0, len(heightMap), len(heightMap[0])
	dequeue, memo := make([][2]int, 0), make([][]int, m)
	d := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			//max = maxVal(max, heightMap[i][j])	// 最高水位
			//memo[i][j] = math.MaxInt32 // 最高水位
			if i == 0 || i == m-1 || j == 0 || j == n-1 {
				memo[i][j] = heightMap[i][j]            // 边界的水位
				dequeue = append(dequeue, [2]int{i, j}) // bfs 初始化
			} else {
				memo[i][j] = math.MaxInt32 // 最高水位
			}
		}
	}
	for ; len(dequeue) > 0; dequeue = dequeue[1:] { // bfs 逻辑
		x, y := dequeue[0][0], dequeue[0][1]
		for i := 0; i < 4; i++ {
			nx, ny := x+d[i][0], y+d[i][1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n-1 &&
				//memo[nx][ny] > memo[x][y] && memo[nx][ny] > heightMap[nx][ny] {	// 漏水
				memo[nx][ny] > memo[x][y] { // 漏水
				memo[nx][ny] = maxVal(memo[x][y], heightMap[nx][ny]) // 最高水位 / 恢复初始柱高
				dequeue = append(dequeue, [2]int{nx, ny})
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			ret += memo[i][j] - heightMap[i][j] // 统计水量
		}
	}
	return ret

	// 堆
	//ret, m, n := 0, len(heightMap)-1, len(heightMap[0])-1
	//if m < 2 || n < 2 {
	//	return 0
	//}
	//h, visited := &trwHp{}, make([][]bool, m+1)
	//d := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	//for i := 0; i <= m; i++ {
	//	visited[i] = make([]bool, n+1)
	//	heap.Push(h, [3]int{i, 0, heightMap[i][0]}) // 行的两侧
	//	heap.Push(h, [3]int{i, n, heightMap[i][n]})
	//	visited[i][0], visited[i][n] = true, true // 已访问
	//	if i == 0 || i == m {
	//		for j := 1; j < n; j++ {
	//			heap.Push(h, [3]int{i, j, heightMap[i][j]}) // 行的中间
	//			visited[i][j] = true
	//		}
	//	}
	//}
	//for h.Len() > 0 {
	//	top := heap.Pop(h).([3]int) // 当前最小高
	//	x, y := top[0], top[1]
	//	for i := 0; i < 4; i++ { // 四连通
	//		nx, ny := x+d[i][0], y+d[i][1]
	//		if nx >= 0 && nx <= m && ny >= 0 && ny <= n && !visited[nx][ny] {
	//			visited[nx][ny] = true
	//			nh := heightMap[nx][ny]
	//			if top[2]-nh > 0 { // 可积攒雨水
	//				ret += top[2] - nh // 雨水量
	//				nh = top[2]        // 更新高
	//			}
	//			heap.Push(h, [3]int{nx, ny, nh}) // 入堆
	//		}
	//	}
	//}
	//return ret
}

type trwHp [][3]int

func (t trwHp) Len() int           { return len(t) }
func (t trwHp) Less(i, j int) bool { return t[i][2] < t[j][2] }
func (t trwHp) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t *trwHp) Push(x any)        { *t = append(*t, x.([3]int)) }
func (t *trwHp) Pop() any {
	v := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return v
}

//leetcode submit region end(Prohibit modification and deletion)
