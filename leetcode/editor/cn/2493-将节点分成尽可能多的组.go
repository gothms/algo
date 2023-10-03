//给你一个正整数 n ，表示一个 无向 图中的节点数目，节点编号从 1 到 n 。
//
// 同时给你一个二维整数数组 edges ，其中 edges[i] = [ai, bi] 表示节点 ai 和 bi 之间有一条 双向 边。注意给定的图可能是不
//连通的。
//
// 请你将图划分为 m 个组（编号从 1 开始），满足以下要求：
//
//
// 图中每个节点都只属于一个组。
// 图中每条边连接的两个点 [ai, bi] ，如果 ai 属于编号为 x 的组，bi 属于编号为 y 的组，那么 |y - x| = 1 。
//
//
// 请你返回最多可以将节点分为多少个组（也就是最大的 m ）。如果没办法在给定条件下分组，请你返回 -1 。
//
//
//
// 示例 1：
//
//
//
// 输入：n = 6, edges = [[1,2],[1,4],[1,5],[2,6],[2,3],[4,6]]
//输出：4
//解释：如上图所示，
//- 节点 5 在第一个组。
//- 节点 1 在第二个组。
//- 节点 2 和节点 4 在第三个组。
//- 节点 3 和节点 6 在第四个组。
//所有边都满足题目要求。
//如果我们创建第五个组，将第三个组或者第四个组中任何一个节点放到第五个组，至少有一条边连接的两个节点所属的组编号不符合题目要求。
//
//
// 示例 2：
//
// 输入：n = 3, edges = [[1,2],[2,3],[3,1]]
//输出：-1
//解释：如果我们将节点 1 放入第一个组，节点 2 放入第二个组，节点 3 放入第三个组，前两条边满足题目要求，但第三条边不满足题目要求。
//没有任何符合题目要求的分组方式。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 500
// 1 <= edges.length <= 10⁴
// edges[i].length == 2
// 1 <= ai, bi <= n
// ai != bi
// 两个点之间至多只有一条边。
//
//
// Related Topics 广度优先搜索 并查集 图 👍 20 👎 0

package main

import "fmt"

func main() {
	arr := []int{2, 3}
	arr = append(arr, 4)
	arr = arr[2:]
	fmt.Println(arr)
}

// leetcode submit region begin(Prohibit modification and deletion)
func magnificentSets(n int, edges [][]int) int {
	// 并查集：https://leetcode.cn/problems/divide-nodes-into-the-maximum-number-of-groups/solutions/2005540/zhou-sai-322t4golangbing-cha-ji-lian-ton-5t7l/

	// bfs
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	visited, adj := make([]int, n+1), make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]], adj[e[1]] = append(adj[e[0]], e[1]), append(adj[e[1]], e[0])
	}
	color, nodes := make([]int8, n+1), make([]int, 0, n)
	var bipartiteDfs func(int, int8) bool
	bipartiteDfs = func(f int, c int8) bool {
		nodes = append(nodes, f) // 记录连通分量的顶点
		color[f] = c
		for _, t := range adj[f] {
			switch color[t] {
			case 0:
				if !bipartiteDfs(t, -c) {
					return false
				}
			case c: // 不是二分图
				return false
			}
		}
		return true
	}
	ret, clock := 0, 0 // visited 通过 clock 标记多次遍历时，顶点是否已被访问
	bfs := func(f int) int {
		clock++
		visited[f] = clock      // clock 标记本次遍历
		depth, q := 0, []int{f} // BFS 的 queue
		for l := len(q); l > 0; l = len(q) {
			depth++
			for i := 0; i < l; i++ {
				for _, j := range adj[q[i]] {
					if visited[j] != clock {
						visited[j] = clock
						q = append(q, j)
					}
				}
			}
			q = q[l:]
		}
		return depth
	}
	for i := 1; i <= n; i++ {
		if color[i] == 0 { // 未访问（color 同时可标记顶点是否已访问）
			nodes = nodes[:0]        // 连通分量中的顶点
			if !bipartiteDfs(i, 1) { // 非二分图，无法按 |y - x| = 1 分组
				return -1
			}
			maxDepth := 0 // 连通分量的最大深度
			for _, f := range nodes {
				maxDepth = maxVal(maxDepth, bfs(f))
			}
			ret += maxDepth
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
