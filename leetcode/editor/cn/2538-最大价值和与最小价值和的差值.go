//给你一个 n 个节点的无向无根图，节点编号为 0 到 n - 1 。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中
//edges[i] = [ai, bi] 表示树中节点 ai 和 bi 之间有一条边。
//
// 每个节点都有一个价值。给你一个整数数组 price ，其中 price[i] 是第 i 个节点的价值。
//
// 一条路径的 价值和 是这条路径上所有节点的价值之和。
//
// 你可以选择树中任意一个节点作为根节点 root 。选择 root 为根的 开销 是以 root 为起点的所有路径中，价值和 最大的一条路径与最小的一条路径
//的差值。
//
// 请你返回所有节点作为根节点的选择中，最大 的 开销 为多少。
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 6, edges = [[0,1],[1,2],[1,3],[3,4],[3,5]], price = [9,8,7,6,10,5]
//输出：24
//解释：上图展示了以节点 2 为根的树。左图（红色的节点）是最大价值和路径，右图（蓝色的节点）是最小价值和路径。
//- 第一条路径节点为 [2,1,3,4]：价值为 [7,8,6,10] ，价值和为 31 。
//- 第二条路径节点为 [2] ，价值为 [7] 。
//最大路径和与最小路径和的差值为 24 。24 是所有方案中的最大开销。
//
//
// 示例 2：
//
//
//
//
//输入：n = 3, edges = [[0,1],[1,2]], price = [1,1,1]
//输出：2
//解释：上图展示了以节点 0 为根的树。左图（红色的节点）是最大价值和路径，右图（蓝色的节点）是最小价值和路径。
//- 第一条路径包含节点 [0,1,2]：价值为 [1,1,1] ，价值和为 3 。
//- 第二条路径节点为 [0] ，价值为 [1] 。
//最大路径和与最小路径和的差值为 2 。2 是所有方案中的最大开销。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// edges.length == n - 1
// 0 <= ai, bi <= n - 1
// edges 表示一棵符合题面要求的树。
// price.length == n
// 1 <= price[i] <= 10⁵
//
//
// Related Topics 树 深度优先搜索 数组 动态规划 👍 41 👎 0

package main

import "fmt"

func main() {
	n := 6
	edges := [][]int{{0, 1},
		{1, 2},
		{1, 3},
		{3, 4},
		{3, 5}}
	price := []int{9, 8, 7, 6, 10, 5}
	//n = 9
	//edges = [][]int{{1, 7},
	//	{5, 2},
	//	{2, 3},
	//	{6, 0},
	//	{0, 4},
	//	{4, 7},
	//	{7, 3},
	//	{3, 8}}
	//price = []int{6, 13, 8, 10, 4, 5, 8, 3, 12}
	output := maxOutput(n, edges, price)
	fmt.Println(output)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxOutput(n int, edges [][]int, price []int) int64 {
	// 查表
	//adj := make([][]int, n)
	//for _, e := range edges {
	//	x, y := e[0], e[1]
	//	adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	//}
	//var dfs func(int, int)int
	//dfs = func(f, t int)int {
	//	for _, i := range adj[t] {
	//		if i!=f{
	//			dfs(t,i)
	//		}
	//	}
	//	return
	//}
	//path := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	path[i] = make([]int, n)
	//	path[i][i] = price[i]
	//}
	////for i, v := range adj {
	////	for _, j := range v {
	////
	////	}
	////}
	//dfs(-1, 0)
	//return 0

	// 树形 dp
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	adj := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	}
	mo := 0
	var dfs func(int, int) (int, int)
	dfs = func(f, t int) (int, int) {
		maxP1, maxP2 := price[t], 0 // 带 / 不带叶子节点的最大路径和
		for _, i := range adj[t] {
			if i != f {
				p1, p2 := dfs(t, i)
				mo = maxVal(mo, maxVal(maxP1+p2, maxP2+p1)) // 带叶子 + 不带叶子 & 不带叶子 + 带叶子
				maxP1 = maxVal(maxP1, p1+price[t])          // 叶子节点的出度为 0，入度 i == f，所以这里肯定不是叶子节点 +price[t]
				maxP2 = maxVal(maxP2, p2+price[t])          // 除非 0 是叶子节点，但是并没有使用 dfs(-1, 0) 的返回值
			}
		}
		return maxP1, maxP2 // 叶子节点将会直接返回初始值：price[t], 0
	}
	dfs(-1, 0) // 返回：带两个叶子节点的值，不带叶子节点 + price[0] 的值
	return int64(mo)

	// dfs：超时
	//maxVal := func(a, b int64) int64 {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//maxV := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//adj := make([][]int, n)
	//for _, e := range edges {
	//	x, y := e[0], e[1]
	//	adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	//}
	//mo := int64(0)
	//var dfs func(int, int, int, int)
	//dfs = func(f, t, sum, s int) {
	//	if f != -1 && len(adj[t]) == 1 {
	//		mo = maxVal(mo, int64(sum+maxV(price[s], price[t])))
	//		return
	//	}
	//	if f != -1 {
	//		sum += price[t]
	//	}
	//	for _, i := range adj[t] {
	//		if i != f {
	//			dfs(t, i, sum, s)
	//		}
	//	}
	//}
	//queue := make([]int, 0)
	//for i, v := range adj {
	//	if len(v) == 1 {
	//		queue = append(queue, i)
	//	}
	//}
	//visited := make([]bool, n)
	//for _, v := range queue {
	//	if !visited[v] {
	//		dfs(-1, v, 0, v)
	//		visited[v] = true
	//	}
	//}
	//return mo
}

//leetcode submit region end(Prohibit modification and deletion)
