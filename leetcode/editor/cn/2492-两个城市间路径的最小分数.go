//给你一个正整数 n ，表示总共有 n 个城市，城市从 1 到 n 编号。给你一个二维数组 roads ，其中 roads[i] = [ai, bi,
//distancei] 表示城市 ai 和 bi 之间有一条 双向 道路，道路距离为 distancei 。城市构成的图不一定是连通的。
//
// 两个城市之间一条路径的 分数 定义为这条路径中道路的 最小 距离。
//
// 城市 1 和城市 n 之间的所有路径的 最小 分数。
//
// 注意：
//
//
// 一条路径指的是两个城市之间的道路序列。
// 一条路径可以 多次 包含同一条道路，你也可以沿着路径多次到达城市 1 和城市 n 。
// 测试数据保证城市 1 和城市n 之间 至少 有一条路径。
//
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 4, roads = [[1,2,9],[2,3,6],[2,4,5],[1,4,7]]
//输出：5
//解释：城市 1 到城市 4 的路径中，分数最小的一条为：1 -> 2 -> 4 。这条路径的分数是 min(9,5) = 5 。
//不存在分数更小的路径。
//
//
// 示例 2：
//
//
//
//
//输入：n = 4, roads = [[1,2,2],[1,3,4],[3,4,7]]
//输出：2
//解释：城市 1 到城市 4 分数最小的路径是：1 -> 2 -> 1 -> 3 -> 4 。这条路径的分数是 min(2,2,4,7) = 2 。
//
//
//
//
// 提示：
//
//
// 2 <= n <= 10⁵
// 1 <= roads.length <= 10⁵
// roads[i].length == 3
// 1 <= ai, bi <= n
// ai != bi
// 1 <= distancei <= 10⁴
// 不会有重复的边。
// 城市 1 和城市 n 之间至少有一条路径。
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 20 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	roads := [][]int{{1, 2, 9},
		{2, 3, 6},
		{2, 4, 5},
		{1, 4, 7}}
	n := 4
	roads = [][]int{{4, 5, 7468}, // 1885
		{6, 2, 7173},
		{6, 3, 8365},
		{2, 3, 7674},
		{5, 6, 7852},
		{1, 2, 8547},
		{2, 4, 1885},
		{2, 5, 5192},
		{1, 3, 4065},
		{1, 4, 7357}}
	n = 6
	score := minScore(n, roads)
	fmt.Println(score)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minScore(n int, roads [][]int) int {
	// 优化
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	adj, visited := make([][][2]int, n+1), make([]bool, n+1)
	for _, r := range roads {
		adj[r[0]] = append(adj[r[0]], [2]int{r[1], r[2]})
		adj[r[1]] = append(adj[r[1]], [2]int{r[0], r[2]})
	}
	ms := math.MaxInt32
	var dfs func(int)
	dfs = func(f int) {
		for _, t := range adj[f] {
			ms = minVal(ms, t[1])
			if !visited[t[0]] {
				visited[f] = true
				dfs(t[0])
			}
		}
	}
	dfs(1)
	visited[1] = true
	return ms

	// dfs
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//adj, visited := make([][][2]int, n+1), make(map[[2]int]bool, n+1)
	//for _, r := range roads {
	//	adj[r[0]] = append(adj[r[0]], [2]int{r[1], r[2]})
	//	adj[r[1]] = append(adj[r[1]], [2]int{r[0], r[2]})
	//}
	//ms := math.MaxInt32
	//var dfs func([2]int)
	//dfs = func(f [2]int) {
	//	ms = minVal(ms, f[1])
	//	for _, t := range adj[f[0]] {
	//		if !visited[[2]int{f[0], t[0]}] {
	//			visited[[2]int{f[0], t[0]}] = true
	//			visited[[2]int{t[0], f[0]}] = true
	//			dfs(t)
	//		}
	//	}
	//}
	//dfs([2]int{1, ms})
	//return ms
}

//leetcode submit region end(Prohibit modification and deletion)
