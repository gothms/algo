//Alice 有一棵 n 个节点的树，节点编号为 0 到 n - 1 。树用一个长度为 n - 1 的二维整数数组 edges 表示，其中 edges[i]
//= [ai, bi] ，表示树中节点 ai 和 bi 之间有一条边。
//
// Alice 想要 Bob 找到这棵树的根。她允许 Bob 对这棵树进行若干次 猜测 。每一次猜测，Bob 做如下事情：
//
//
// 选择两个 不相等 的整数 u 和 v ，且树中必须存在边 [u, v] 。
// Bob 猜测树中 u 是 v 的 父节点 。
//
//
// Bob 的猜测用二维整数数组 guesses 表示，其中 guesses[j] = [uj, vj] 表示 Bob 猜 uj 是 vj 的父节点。
//
// Alice 非常懒，她不想逐个回答 Bob 的猜测，只告诉 Bob 这些猜测里面 至少 有 k 个猜测的结果为 true 。
//
// 给你二维整数数组 edges ，Bob 的所有猜测和整数 k ，请你返回可能成为树根的 节点数目 。如果没有这样的树，则返回 0。
//
//
//
// 示例 1：
//
//
//
//
//输入：edges = [[0,1],[1,2],[1,3],[4,2]], guesses = [[1,3],[0,1],[1,0],[2,4]], k =
// 3
//输出：3
//解释：
//根为节点 0 ，正确的猜测为 [1,3], [0,1], [2,4]
//根为节点 1 ，正确的猜测为 [1,3], [1,0], [2,4]
//根为节点 2 ，正确的猜测为 [1,3], [1,0], [2,4]
//根为节点 3 ，正确的猜测为 [1,0], [2,4]
//根为节点 4 ，正确的猜测为 [1,3], [1,0]
//节点 0 ，1 或 2 为根时，可以得到 3 个正确的猜测。
//
//
// 示例 2：
//
//
//
//
//输入：edges = [[0,1],[1,2],[2,3],[3,4]], guesses = [[1,0],[3,4],[2,1],[3,2]], k =
// 1
//输出：5
//解释：
//根为节点 0 ，正确的猜测为 [3,4]
//根为节点 1 ，正确的猜测为 [1,0], [3,4]
//根为节点 2 ，正确的猜测为 [1,0], [2,1], [3,4]
//根为节点 3 ，正确的猜测为 [1,0], [2,1], [3,2], [3,4]
//根为节点 4 ，正确的猜测为 [1,0], [2,1], [3,2]
//任何节点为根，都至少有 1 个正确的猜测。
//
//
//
//
// 提示：
//
//
// edges.length == n - 1
// 2 <= n <= 10⁵
// 1 <= guesses.length <= 10⁵
// 0 <= ai, bi, uj, vj <= n - 1
// ai != bi
// uj != vj
// edges 表示一棵有效的树。
// guesses[j] 是树中的一条边。
// guesses 是唯一的。
// 0 <= k <= guesses.length
//
//
// Related Topics 树 深度优先搜索 哈希表 动态规划 👍 40 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func rootCount(edges [][]int, guesses [][]int, k int) int {
	// dp
	ret, cnt, n := 0, 0, len(edges)+1
	memo, adj := make(map[[2]int]int, n-1), make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	for _, g := range guesses {
		memo[[2]int{g[0], g[1]}] = 1
	}
	var dfs func(f, t int)
	dfs = func(f, t int) {
		for _, x := range adj[t] {
			if x != f {
				if memo[[2]int{t, x}] == 1 {
					cnt++
				}
				dfs(t, x)
			}
		}
	}
	dfs(-1, 0)
	var dfsReroot func(f, t int, c int)
	dfsReroot = func(f, t int, c int) {
		if c >= k {
			ret++
		}
		for _, x := range adj[t] {
			if x != f {
				dfsReroot(t, x, c-memo[[2]int{t, x}]+memo[[2]int{x, t}])
			}
		}
	}
	dfsReroot(-1, 0, cnt)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
