//给你一个整数 n ，表示一张 无向图 中有 n 个节点，编号为 0 到 n - 1 。同时给你一个二维整数数组 edges ，其中 edges[i] = [
//ai, bi] 表示节点 ai 和 bi 之间有一条 无向 边。
//
// 请你返回 无法互相到达 的不同 点对数目 。
//
//
//
// 示例 1：
//
//
//
// 输入：n = 3, edges = [[0,1],[0,2],[1,2]]
//输出：0
//解释：所有点都能互相到达，意味着没有点对无法互相到达，所以我们返回 0 。
//
//
// 示例 2：
//
//
//
// 输入：n = 7, edges = [[0,2],[0,5],[2,4],[1,6],[5,4]]
//输出：14
//解释：总共有 14 个点对互相无法到达：
//[[0,1],[0,3],[0,6],[1,2],[1,3],[1,4],[1,5],[2,3],[2,6],[3,4],[3,5],[3,6],[4,6]
//,[5,6]]
//所以我们返回 14 。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// 0 <= edges.length <= 2 * 10⁵
// edges[i].length == 2
// 0 <= ai, bi < n
// ai != bi
// 不会有重复边。
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 91 👎 0

package main

import "fmt"

func main() {
	n := 7
	edges := [][]int{{0, 2},
		{0, 5},
		{2, 4},
		{1, 6},
		{5, 4}}
	i := countPairs(n, edges)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func countPairs(n int, edges [][]int) int64 {
	// dfs + 前缀和
	adj, visited := make([][]int, n), make([]bool, n) // 邻接表
	for _, e := range edges {
		x, y := e[0], e[1]
		adj[x] = append(adj[x], y)
		adj[y] = append(adj[y], x)
	}
	var cnt, sum int64
	var dfs func(int) int
	dfs = func(i int) int {
		visited[i] = true
		c := 1
		for _, j := range adj[i] {
			if !visited[j] {
				c += dfs(j)
			}
		}
		return c
	}
	for i := 0; i < n; i++ {
		if !visited[i] {
			v := int64(dfs(i))
			cnt += sum * v
			sum += v
		}
	}
	return cnt

	// 并查集
	//adj := make([][]int, n) // 邻接表
	//for _, e := range edges {
	//	x, y := e[0], e[1]
	//	adj[x] = append(adj[x], y)
	//	adj[y] = append(adj[y], x)
	//}
	//uni, counter := make([]int, n), make([]int, n)
	//for i := 0; i < n; i++ {
	//	uni[i] = i
	//	counter[i] = 1 // 集合的元素个数
	//}
	//find := func(p int) int {
	//	for p != uni[p] {
	//		p, uni[p] = uni[p], uni[uni[p]]
	//	}
	//	return uni[p]
	//}
	//union := func(p, q int) bool {
	//	rootP, rootQ := find(p), find(q)
	//	if rootP != rootQ {
	//		uni[rootQ] = rootP
	//		counter[p] += counter[q]
	//		counter[q] = 0
	//		return true
	//	}
	//	return false
	//}
	//var dfs func(int, int)
	//dfs = func(f, t int) {
	//	for _, i := range adj[t] {
	//		if counter[i] != 0 { // 0 标记已遍历
	//			union(f, i)
	//			dfs(f, i)
	//		}
	//	}
	//}
	//for i := 0; i < n; i++ {
	//	if counter[i] == 1 {
	//		dfs(i, i)
	//	}
	//}
	//var cnt, sum int64
	//for i := range counter {
	//	v := int64(counter[i])
	//	if v > 0 {
	//		cnt += v * sum
	//		sum += v // 前缀和
	//	}
	//}
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
