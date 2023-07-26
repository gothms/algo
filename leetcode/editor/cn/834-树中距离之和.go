//给定一个无向、连通的树。树中有 n 个标记为 0...n-1 的节点以及 n-1 条边 。
//
// 给定整数 n 和数组 edges ， edges[i] = [ai, bi]表示树中的节点 ai 和 bi 之间有一条边。
//
// 返回长度为 n 的数组 answer ，其中 answer[i] 是树中第 i 个节点与所有其他节点之间的距离之和。
//
//
//
// 示例 1:
//
//
//
//
//输入: n = 6, edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
//输出: [8,12,6,10,10,10]
//解释: 树如图所示。
//我们可以计算出 dist(0,1) + dist(0,2) + dist(0,3) + dist(0,4) + dist(0,5)
//也就是 1 + 1 + 2 + 2 + 2 = 8。 因此，answer[0] = 8，以此类推。
//
//
// 示例 2:
//
//
//输入: n = 1, edges = []
//输出: [0]
//
//
// 示例 3:
//
//
//输入: n = 2, edges = [[1,0]]
//输出: [1,1]
//
//
//
//
// 提示:
//
//
// 1 <= n <= 3 * 10⁴
// edges.length == n - 1
// edges[i].length == 2
// 0 <= ai, bi < n
// ai != bi
// 给定的输入保证为有效的树
//
//
// Related Topics 树 深度优先搜索 图 动态规划 👍 393 👎 0

package main

import "fmt"

func main() {
	n := 6
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}
	n = 2
	edges = [][]int{{1, 0}}
	n = 3
	edges = [][]int{{2, 0}, {1, 0}} // [2,3,3]
	n = 4
	edges = [][]int{{1, 2}, {3, 2}, {3, 0}} // [6,6,4,4]
	tree := sumOfDistancesInTree(n, edges)
	fmt.Println(tree)
}

/*
思路：回溯
	lc：https://leetcode.cn/problems/sum-of-distances-in-tree/solution/shu-zhong-ju-chi-zhi-he-by-leetcode-solution/
	更优秀的题解：
		https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/
扩展
	思考题
		如果只需要算所有点对的距离之和，你能想出一个只需要一次 DFS 的算法吗？
			提示：我在前天的每日一题中讲解了一种技巧，请看【图解】没有思路？进来秒懂！
		把题目中的「距离之和」改成「距离的平方和」要怎么做？
			提示 1：换根，把「变化量的平方」这个式子展开。
			提示 2：除了计算子树大小，还需要计算子树中的每个节点的深度之和。
		改成「距离的立方和」要怎么做？
练习：换根 DP
	310. 最小高度树（做法不止一种）
	2581. 统计可能的树根数目
	Codeforces 771C. Bear and Tree Jumps（本题进阶版）

思路：dp + dfs
    1.把树看成图，每个节点是一个顶点，edges[i]是一条边，根据 edges 构建邻接表adj（无向）
		注意，0 并不一定是树的 root 节点
	2.假设只用dfs解决
		遍历 [0,n) 的每个顶点 i，顶点 i 的距离和是：与 i 连接的顶点 j 的距离和 + j 子树的节点数
		伪代码：时间复杂度是 O(n^2)
		ans := make([]int, n)
		func dfs(c int, cnt, nodesCnt []int) {
			for i range adj[c] {
				dfs(i)
				cnt[c] += cnt[i] + nodesCnt[i]	// 记录节点 c 的子节点们的距离总和
				nodesCnt[c] += nodesCnt[i]	// 记录节点 c 为root的子树的节点总数
			}
		}
		for i range n {
			cnt, nodesCnt := make([]int, n), make([]int, n)
			dfs(i, cnt, nodesCnt)
			ans[i] = cnt[i]	// 计算每一个节点的距离和
		}
	3.上面方案的时间复杂度有优化的空间吗？
		3.1.发现上面的dfs中，先是把邻接表中的边的关系（可以理解为父节点 & 子节点）往下一层“递”
		3.2.然后子节点把 它的子树距离和 和 它的子树节点总数 往上一层“归”
		3.3.当我们计算得到 当前root节点i 的距离和后，如果继续把这个距离和“递”给它的子节点 j
			就可以“逆向”计算，把 j 作为root节点，并把 i 当成 j 的子节点
	4.dp
		4.1.我们改造 for i range n {} 的实现为 dp，只发起一次root节点的dfs
		4.2.我们记录 dfs 方案中，计算所得结果为 dp0 = cnt，tree0 = nodesCnt
			新建 dp 和 tree 作为dp方案的计算值的存储
			dp[i]：计算顶点 i 的最终距离和
			tree[i]：计算顶点 i 的最终子树节点总数
		4.3.状态转移方程
			dp[i] += dp[c] - dp0[i] - tree0[i] + tree[c] - tree0[i]
			tree[i] += tree[c] - tree0[i]
			释义：c为“父”节点，i为当前节点
				dp[c] - dp0[i] - tree0[i]：父节点抹除分支 子节点i 的距离和
				tree[c] - tree0[i]：节点i 把 c 当成子节点，c子树的所有节点都 +1 距离，“到达”i
		4.4.重点：如代码所示
			dfs方法，对值的算法是“归”，所以先递后归值
			dfsDP方法，对值的算法是“递”，所以先递值后归
	5.优化：dp 和 tree 都是多余的
		tree[i] 的值都是 n
		dp 的计算是按照邻接表 dfs 的，每次 i range adj[c] 计算后的值都是最终值
*/
//leetcode submit region begin(Prohibit modification and deletion)
func sumOfDistancesInTree(n int, edges [][]int) []int {
	//adj := make([][]int, n)                      // 邻接表
	//dp0, tree0 := make([]int, n), make([]int, n) // 初始化
	//dp, tree := make([]int, n), make([]int, n)   // 最终值
	//for _, e := range edges {
	//	adj[e[0]], adj[e[1]] = append(adj[e[0]], e[1]), append(adj[e[1]], e[0])
	//}
	//for i := 0; i < n; i++ {
	//	tree0[i] = 1 // 每个子树最小节点数为 1
	//}
	//var dfs func(int, int, bool)
	//dfs = func(p, c int, l bool) {
	//	for _, i := range adj[c] {
	//		if i == p {
	//			continue
	//		}
	//		if l {
	//			dp[i] += dp[c] - dp0[i] - tree0[i] + tree[c] - tree0[i] // 逆向计算
	//			tree[i] += tree[c] - tree0[i]
	//			dfs(c, i, l)
	//		} else {
	//			dfs(c, i, l)
	//			dp0[c] += dp0[i] + tree0[i]
	//			tree0[c] += tree0[i]
	//		}
	//	}
	//}
	//dfs(-1, 0, false) // 初始化
	//copy(dp, dp0)
	//copy(tree, tree0)
	//dfs(-1, 0, true) // 最终值
	//return dp
	// dp：最终，dp 和 tree 都是多余的
	dp0, tree0, adj := make([]int, n), make([]int, n), make([][]int, n)
	for _, e := range edges { // 初始化邻接表
		adj[e[0]], adj[e[1]] = append(adj[e[0]], e[1]), append(adj[e[1]], e[0])
	}
	var dfs func(int, int)
	dfs = func(p, c int) { // “归”值，所以计算的是 c
		tree0[c] = 1 // 每个节点的子树，至少有一个节点，就是它本身
		for _, i := range adj[c] {
			if i != p {
				dfs(c, i)
				dp0[c] += dp0[i] + tree0[i]
				tree0[c] += tree0[i]
			}
		}
	}
	dfs(-1, 0) //  // 初始化 dp0 和 tree0（以 0 为 root，计算各子树的距离和）
	//dp, tree := make([]int, n), make([]int, n)	// dp 和 tree 都是多余的
	//dp := make([]int, n)
	//copy(dp, dp0)
	//copy(tree, tree0)
	var dfsDP func(int, int)
	dfsDP = func(p, c int) { // “递”值，所以计算的是 i
		for _, i := range adj[c] {
			if i != p { // 以 i 为 root，根据“父节点”逆向计算 i 的 距离和
				//dp[i] += dp[c] - dp0[i] + n - tree0[i]<<1
				dp0[i] += dp0[c] - dp0[i] + n - tree0[i]<<1
				//tree[i] += tree[c] - tree0[i] // i 当前子树节点总数
				//tree[i] = n
				dfsDP(c, i)
			}
		}
	}
	dfsDP(-1, 0) // 计算每一个节点的距离和
	return dp0

	// lc
	//graph := make([][]int, n)
	//for _, e := range edges {
	//	u, v := e[0], e[1]
	//	graph[u] = append(graph[u], v)
	//	graph[v] = append(graph[v], u)
	//}
	//
	//sz := make([]int, n)
	//dp := make([]int, n)
	//var dfs func(u, f int)
	//dfs = func(u, f int) {
	//	sz[u] = 1
	//	for _, v := range graph[u] {
	//		if v == f {
	//			continue
	//		}
	//		dfs(v, u)
	//		dp[u] += dp[v] + sz[v]
	//		sz[u] += sz[v]
	//	}
	//}
	//dfs(0, -1)
	//
	//ans := make([]int, n)
	//var dfs2 func(u, f int)
	//dfs2 = func(u, f int) {
	//	ans[u] = dp[u]
	//	for _, v := range graph[u] {
	//		if v == f {
	//			continue
	//		}
	//		pu, pv := dp[u], dp[v]
	//		su, sv := sz[u], sz[v]
	//
	//		dp[u] -= dp[v] + sz[v]
	//		sz[u] -= sz[v]
	//		dp[v] += dp[u] + sz[u]
	//		sz[v] += sz[u]
	//
	//		dfs2(v, u)
	//
	//		dp[u], dp[v] = pu, pv
	//		sz[u], sz[v] = su, sv
	//	}
	//}
	//dfs2(0, -1)
	//return ans

	// dp：个人写法，题意是任一节点可以多父节点，多子节点
	//dp, adj, rAdj, rVisited, m := make([]int, n), make([][]int, n), make([][]int, n), make([]bool, n), n-1
	//if n == 1 {
	//	return dp
	//}
	//for i := 0; i < m; i++ {
	//	//dp[edges[i][0]]++
	//	p, c := edges[i][0], edges[i][1]
	//	adj[p] = append(adj[p], c)
	//	rAdj[c] = append(rAdj[c], p)
	//}
	//child := make([]int, n)
	//root := edges[0][0]
	////rAdj[0] = -1
	//rVisited[root] = true
	//rQueue := make([]int, 0)
	//for i := 0; i < n; i++ {
	//	if len(adj[i]) == 0 {
	//		rQueue = append(rQueue, i)
	//		rVisited[i] = true
	//	}
	//}
	//fmt.Println(rQueue)
	//for l := len(rQueue); l > 0; l = len(rQueue) {
	//	for i := 0; i < l; i++ {
	//		for _, p := range rAdj[rQueue[i]] {
	//			if !rVisited[p] {
	//				rQueue = append(rQueue, p)
	//				rVisited[p] = true
	//			}
	//			dp[p] += dp[rQueue[i]]<<1 + 1
	//			child[p] += child[rQueue[i]] + 1
	//		}
	//		//p := rAdj[rQueue[i]]
	//		//fmt.Println(p)
	//		//if !rVisited[p] {
	//		//	rQueue = append(rQueue, p)
	//		//	rVisited[p] = true
	//		//}
	//		//dp[p] += dp[rQueue[i]]<<1 + 1
	//		//child[p] += child[rQueue[i]] + 1
	//		//fmt.Println(rQueue[i], dp[rQueue[i]], dp[p])
	//	}
	//	rQueue = rQueue[l:]
	//	//fmt.Println(rQueue)
	//}
	////fmt.Println(rVisited)
	//fmt.Println("dp", dp)
	//fmt.Println(adj)
	//fmt.Println(rAdj)
	////fmt.Println(len(adj[1]))
	//fmt.Println("child", child)
	//queue, visited := make([]int, 0), make([]bool, n)
	//for i := 0; i < n; i++ {
	//	if len(rAdj[i]) == 0 {
	//		queue = append(queue, i)
	//		visited[i] = true
	//	}
	//}
	//fmt.Println(queue)
	//for l := len(queue); l > 0; l = len(queue) {
	//	for i := 0; i < l; i++ {
	//		for _, c := range adj[queue[i]] {
	//			//dp[j] += n - 2 - child[j] + 1 + dp[i] - 1 - dp[j]<<1
	//			dp[c] += n - 2 - child[c] + 1 + dp[queue[i]] - 1 - dp[c]<<1
	//			fmt.Println(c, queue[i], dp[c])
	//			if !visited[c] {
	//				queue = append(queue, c)
	//				visited[c] = true
	//			}
	//		}
	//	}
	//	//1 + 3 + 8-3<<1-1 + 1
	//	queue = queue[l:]
	//	//fmt.Println("queue", queue)
	//}
	//return dp
}

//leetcode submit region end(Prohibit modification and deletion)
