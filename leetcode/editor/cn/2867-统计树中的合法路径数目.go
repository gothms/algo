//给你一棵 n 个节点的无向树，节点编号为 1 到 n 。给你一个整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[i] =
// [ui, vi] 表示节点 ui 和 vi 在树中有一条边。
//
// 请你返回树中的 合法路径数目 。
//
// 如果在节点 a 到节点 b 之间 恰好有一个 节点的编号是质数，那么我们称路径 (a, b) 是 合法的 。
//
// 注意：
//
//
// 路径 (a, b) 指的是一条从节点 a 开始到节点 b 结束的一个节点序列，序列中的节点 互不相同 ，且相邻节点之间在树上有一条边。
// 路径 (a, b) 和路径 (b, a) 视为 同一条 路径，且只计入答案 一次 。
//
//
//
//
// 示例 1：
//
//
//
//
//输入：n = 5, edges = [[1,2],[1,3],[2,4],[2,5]]
//输出：4
//解释：恰好有一个质数编号的节点路径有：
//- (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
//- (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
//- (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
//- (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
//只有 4 条合法路径。
//
//
// 示例 2：
//
//
//
//
//输入：n = 6, edges = [[1,2],[1,3],[2,4],[3,5],[3,6]]
//输出：6
//解释：恰好有一个质数编号的节点路径有：
//- (1, 2) 因为路径 1 到 2 只包含一个质数 2 。
//- (1, 3) 因为路径 1 到 3 只包含一个质数 3 。
//- (1, 4) 因为路径 1 到 4 只包含一个质数 2 。
//- (1, 6) 因为路径 1 到 6 只包含一个质数 3 。
//- (2, 4) 因为路径 2 到 4 只包含一个质数 2 。
//- (3, 6) 因为路径 3 到 6 只包含一个质数 3 。
//只有 6 条合法路径。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// edges.length == n - 1
// edges[i].length == 2
// 1 <= ui, vi <= n
// 输入保证 edges 形成一棵合法的树。
//
//
// Related Topics 树 深度优先搜索 数学 动态规划 数论 👍 19 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countPaths(n int, edges [][]int) int64 {
	ret, adj := int64(0), make([][]int, n+1)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	//sizes, set := make([]int, n+1), make([]int, 0)
	sizes := make([]int, n+1)
	var set []int
	var dfs func(int, int)
	dfs = func(f, t int) {
		set = append(set, t)
		for _, v := range adj[t] {
			if v != f && cpPrimes[v] {
				dfs(t, v)
			}
		}
	}
	for i := 1; i <= n; i++ {
		if cpPrimes[i] {
			continue
		}
		sum := 0
		for _, t := range adj[i] {
			if !cpPrimes[t] {
				continue
			}
			if sizes[t] == 0 {
				//set = set[:0]
				set = []int{}
				dfs(-1, t)
				for _, v := range set {
					sizes[v] = len(set)
				}
			}
			ret += int64(sum * sizes[t])
			sum += sizes[t]
		}
		ret += int64(sum)
	}
	return ret
}

var cpPrimes []bool

func init() {
	const P = 100_001
	cpPrimes = make([]bool, P)
	cpPrimes[1] = true // 1
	for i := 2; i*i < P; i++ {
		if !cpPrimes[i] {
			for j := i * i; j < P; j += i {
				cpPrimes[j] = true
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
