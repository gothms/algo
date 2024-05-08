//给定一个整数 n，即有向图中的节点数，其中节点标记为 0 到 n - 1。图中的每条边为红色或者蓝色，并且可能存在自环或平行边。
//
// 给定两个数组 redEdges 和 blueEdges，其中：
//
//
// redEdges[i] = [ai, bi] 表示图中存在一条从节点 ai 到节点 bi 的红色有向边，
// blueEdges[j] = [uj, vj] 表示图中存在一条从节点 uj 到节点 vj 的蓝色有向边。
//
//
// 返回长度为 n 的数组 answer，其中 answer[X] 是从节点 0 到节点 X 的红色边和蓝色边交替出现的最短路径的长度。如果不存在这样的路径，
//那么 answer[x] = -1。
//
//
//
// 示例 1：
//
//
//输入：n = 3, red_edges = [[0,1],[1,2]], blue_edges = []
//输出：[0,1,-1]
//
//
// 示例 2：
//
//
//输入：n = 3, red_edges = [[0,1]], blue_edges = [[2,1]]
//输出：[0,1,-1]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 0 <= redEdges.length, blueEdges.length <= 400
// redEdges[i].length == blueEdges[j].length == 2
// 0 <= ai, bi, uj, vj < n
//
//
// Related Topics 广度优先搜索 图 👍 303 👎 0

package main

import "fmt"

func main() {
	red := [][]int{{0, 1}}
	blue := [][]int{{1, 2}}
	n := 3
	red = [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}
	blue = [][]int{{1, 2}, {2, 3}, {3, 1}}
	n = 5 // 0,1,2,3,7
	paths := shortestAlternatingPaths(n, red, blue)
	fmt.Println(paths)
}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	// bfs
	ans, edges := make([]int, n), make([][][2]bool, n)
	for i := range edges {
		edges[i] = make([][2]bool, n)
		ans[i] = -1 // 默认值
	}
	ans[0] = 0

	for _, e := range redEdges { // 邻接矩阵
		edges[e[0]][e[1]][0] = true
	}
	for _, e := range blueEdges {
		edges[e[0]][e[1]][1] = true
	}

	vis := make([][2]bool, n) // 标记边已访问
	vis[0][0], vis[0][1] = true, true
	//for i := 1; i < n; i++ {
	q := [][3]int{{0, 0, 0}, {0, 1, 0}}
	for ; len(q) > 0; q = q[1:] {
		j, c, d := q[0][0], q[0][1], q[0][2]+1
		for k, b := range edges[j] {
			if !b[c] || vis[k][c] { // 无边 / 已访问
				continue
			}
			vis[k][c] = true
			if ans[k] == -1 || ans[k] > d { // 因为可以从 红/蓝 出发，所以要比较两个值
				ans[k] = d
			}
			q = append(q, [3]int{k, c ^ 1, d}) // bfs
		}
	}
	//}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
