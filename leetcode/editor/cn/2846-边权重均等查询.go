//现有一棵由 n 个节点组成的无向树，节点按从 0 到 n - 1 编号。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中
//edges[i] = [ui, vi, wi] 表示树中存在一条位于节点 ui 和节点 vi 之间、权重为 wi 的边。
//
// 另给你一个长度为 m 的二维整数数组 queries ，其中 queries[i] = [ai, bi] 。对于每条查询，请你找出使从 ai 到 bi 路
//径上每条边的权重相等所需的 最小操作次数 。在一次操作中，你可以选择树上的任意一条边，并将其权重更改为任意值。
//
// 注意：
//
//
// 查询之间 相互独立 的，这意味着每条新的查询时，树都会回到 初始状态 。
// 从 ai 到 bi的路径是一个由 不同 节点组成的序列，从节点 ai 开始，到节点 bi 结束，且序列中相邻的两个节点在树中共享一条边。
//
//
// 返回一个长度为 m 的数组 answer ，其中 answer[i] 是第 i 条查询的答案。
//
//
//
// 示例 1：
//
//
//输入：n = 7, edges = [[0,1,1],[1,2,1],[2,3,1],[3,4,2],[4,5,2],[5,6,2]], queries =
// [[0,3],[3,6],[2,6],[0,6]]
//输出：[0,0,1,3]
//解释：第 1 条查询，从节点 0 到节点 3 的路径中的所有边的权重都是 1 。因此，答案为 0 。
//第 2 条查询，从节点 3 到节点 6 的路径中的所有边的权重都是 2 。因此，答案为 0 。
//第 3 条查询，将边 [2,3] 的权重变更为 2 。在这次操作之后，从节点 2 到节点 6 的路径中的所有边的权重都是 2 。因此，答案为 1 。
//第 4 条查询，将边 [0,1]、[1,2]、[2,3] 的权重变更为 2 。在这次操作之后，从节点 0 到节点 6 的路径中的所有边的权重都是 2 。因此
//，答案为 3 。
//对于每条查询 queries[i] ，可以证明 answer[i] 是使从 ai 到 bi 的路径中的所有边的权重相等的最小操作次数。
//
//
// 示例 2：
//
//
//输入：n = 8, edges = [[1,2,6],[1,3,4],[2,4,6],[2,5,3],[3,6,6],[3,0,8],[7,0,2]],
//queries = [[4,6],[0,4],[6,5],[7,4]]
//输出：[1,2,2,3]
//解释：第 1 条查询，将边 [1,3] 的权重变更为 6 。在这次操作之后，从节点 4 到节点 6 的路径中的所有边的权重都是 6 。因此，答案为 1 。
//第 2 条查询，将边 [0,3]、[3,1] 的权重变更为 6 。在这次操作之后，从节点 0 到节点 4 的路径中的所有边的权重都是 6 。因此，答案为 2
// 。
//第 3 条查询，将边 [1,3]、[5,2] 的权重变更为 6 。在这次操作之后，从节点 6 到节点 5 的路径中的所有边的权重都是 6 。因此，答案为 2
// 。
//第 4 条查询，将边 [0,7]、[0,3]、[1,3] 的权重变更为 6 。在这次操作之后，从节点 7 到节点 4 的路径中的所有边的权重都是 6 。因此
//，答案为 3 。
//对于每条查询 queries[i] ，可以证明 answer[i] 是使从 ai 到 bi 的路径中的所有边的权重相等的最小操作次数。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁴
// edges.length == n - 1
// edges[i].length == 3
// 0 <= ui, vi < n
// 1 <= wi <= 26
// 生成的输入满足 edges 表示一棵有效的树
// 1 <= queries.length == m <= 2 * 10⁴
// queries[i].length == 2
// 0 <= ai, bi < n
//
//
// Related Topics 树 图 数组 强连通分量 👍 21 👎 0

package main

import "slices"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	type edge struct{ to, wt int }
	g := make([][]edge, n)
	for _, e := range edges {
		v, w, wt := e[0], e[1], e[2]-1
		g[v] = append(g[v], edge{w, wt})
		g[w] = append(g[w], edge{v, wt})
	}

	const mx = 14 // 2^14 > 10^4
	type pair struct {
		p   int
		cnt [26]int
	}
	pa := make([][mx]pair, n)
	depth := make([]int, n)
	var build func(int, int, int)
	build = func(v, p, d int) {
		pa[v][0].p = p
		depth[v] = d
		for _, e := range g[v] {
			if w := e.to; w != p {
				pa[w][0].cnt[e.wt] = 1
				build(w, v, d+1)
			}
		}
	}
	build(0, -1, 0)

	// 倍增模板
	for i := 0; i+1 < mx; i++ {
		for v := range pa {
			if p := pa[v][i]; p.p != -1 {
				pp := pa[p.p][i]
				pa[v][i+1].p = pp.p
				for j := 0; j < 26; j++ {
					pa[v][i+1].cnt[j] = p.cnt[j] + pp.cnt[j]
				}
			} else {
				pa[v][i+1].p = -1
			}
		}
	}
	f := func(v, w int) int {
		pathLen := depth[v] + depth[w] // 最后减去 depth[lca] * 2
		cnt := [26]int{}
		if depth[v] > depth[w] {
			v, w = w, v
		}
		for i := 0; i < mx; i++ {
			if (depth[w]-depth[v])>>i&1 > 0 {
				p := pa[w][i]
				for j := 0; j < 26; j++ {
					cnt[j] += p.cnt[j]
				}
				w = p.p
			}
		}
		if w != v {
			for i := mx - 1; i >= 0; i-- {
				if pv, pw := pa[v][i], pa[w][i]; pv.p != pw.p {
					for j := 0; j < 26; j++ {
						cnt[j] += pv.cnt[j] + pw.cnt[j]
					}
					v, w = pv.p, pw.p
				}
			}
			for j := 0; j < 26; j++ {
				cnt[j] += pa[v][0].cnt[j] + pa[w][0].cnt[j]
			}
			v = pa[v][0].p
		}
		// 现在 v 是 LCA
		return pathLen - depth[v]*2 - slices.Max(cnt[:])
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = f(q[0], q[1])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
