//给你一个 n 个节点的无向无根树，节点编号从 0 到 n - 1 。给你整数 n 和一个长度为 n - 1 的二维整数数组 edges ，其中 edges[
//i] = [ai, bi] 表示树中节点 ai 和 bi 之间有一条边。再给你一个长度为 n 的数组 coins ，其中 coins[i] 可能为 0 也可能为
// 1 ，1 表示节点 i 处有一个金币。
//
// 一开始，你需要选择树中任意一个节点出发。你可以执行下述操作任意次：
//
//
// 收集距离当前节点距离为 2 以内的所有金币，或者
// 移动到树中一个相邻节点。
//
//
// 你需要收集树中所有的金币，并且回到出发节点，请你返回最少经过的边数。
//
// 如果你多次经过一条边，每一次经过都会给答案加一。
//
//
//
// 示例 1：
//
//
//
// 输入：coins = [1,0,0,0,0,1], edges = [[0,1],[1,2],[2,3],[3,4],[4,5]]
//输出：2
//解释：从节点 2 出发，收集节点 0 处的金币，移动到节点 3 ，收集节点 5 处的金币，然后移动回节点 2 。
//
//
// 示例 2：
//
//
//
// 输入：coins = [0,0,0,1,1,0,0,1], edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[5,6],[5
//,7]]
//输出：2
//解释：从节点 0 出发，收集节点 4 和 3 处的金币，移动到节点 2 处，收集节点 7 处的金币，移动回节点 0 。
//
//
//
//
// 提示：
//
//
// n == coins.length
// 1 <= n <= 3 * 10⁴
// 0 <= coins[i] <= 1
// edges.length == n - 1
// edges[i].length == 2
// 0 <= ai, bi < n
// ai != bi
// edges 表示一棵合法的树。
//
//
// Related Topics 树 图 拓扑排序 数组 👍 34 👎 0

package main

import "fmt"

func main() {
	coins := []int{1, 0, 0, 0, 0, 1}
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}
	coins = []int{0, 0, 0, 1, 1, 0, 0, 1}
	edges = [][]int{{0, 1},
		{0, 2},
		{1, 3},
		{1, 4},
		{2, 5},
		{5, 6},
		{5, 7}}
	theCoins := collectTheCoins(coins, edges)
	fmt.Println(theCoins)
}

// leetcode submit region begin(Prohibit modification and deletion)
func collectTheCoins(coins []int, edges [][]int) int {
	// topological sort
	cnt, n := len(edges), len(coins) // cnt = 总边数
	kahn, adj := make([]int, n), make([][]int, n)
	for _, e := range edges { // 初始化邻接表和 kahn
		x, y := e[0], e[1]
		adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
		kahn[x]++
		kahn[y]++
	}
	queue := make([]int, 0, n)
	for i := range kahn {
		if kahn[i] == 1 && coins[i] == 0 { // 叶子节点，且没有硬币
			queue = append(queue, i)
		}
	}
	for ; len(queue) > 0; queue = queue[1:] {
		idx := queue[0]
		cnt-- // 删除“叶子”节点的边
		for _, i := range adj[idx] {
			kahn[i]--
			if kahn[i] == 1 && coins[i] == 0 { // 新的叶子节点，且没有硬币
				queue = append(queue, i)
			}
		}
	}
	for i := range kahn {
		if kahn[i] == 1 && coins[i] == 1 { // 叶子节点，且没有硬币
			queue = append(queue, i)
		}
	}
	cnt -= len(queue) // 删除有硬币的边
	for _, idx := range queue {
		for _, i := range adj[idx] {
			kahn[i]--
			if kahn[i] == 1 { // 删除与硬币节点相连的边
				cnt-- // =1 表示不用经过这个节点
			}
		}
	}
	if cnt <= 0 { // 没有硬币 / 无需挪动位置
		return 0
	}
	return cnt << 1 // 回到出发节点

	// topological sort
	//cnt, n := len(edges), len(coins)
	//kahn, adj := make([]int, n), make([][]int, n)
	//for _, e := range edges {
	//	x, y := e[0], e[1]
	//	adj[x], adj[y] = append(adj[x], y), append(adj[y], x)
	//	kahn[x]++
	//	kahn[y]++
	//}
	//queue := make([]int, 0, n)
	//for i, c := range kahn {
	//	if c == 1 && coins[i] == 0 { // 多余节点
	//		queue = append(queue, i)
	//	}
	//}
	//for ; len(queue) > 0; queue = queue[1:] {
	//	i := queue[0]
	//	cnt--
	//	for _, j := range adj[i] {
	//		kahn[j]--
	//		if kahn[j] == 1 && coins[j] == 0 {
	//			queue = append(queue, j)
	//		}
	//	}
	//}
	//for i, c := range kahn {
	//	if c == 1 && coins[i] == 1 { // 金币节点
	//		queue = append(queue, i)
	//	}
	//}
	//for _, i := range queue {
	//	for _, j := range adj[i] {
	//		kahn[j]--
	//		if kahn[j] == 1 {
	//			cnt--
	//		}
	//	}
	//}
	//if cnt -= len(queue); cnt < 0 {
	//	return 0
	//}
	//return cnt << 1

	// lc
	//n := len(coins)
	//g := make([][]int, n)
	//deg := make([]int, n)
	//for _, e := range edges {
	//	x, y := e[0], e[1]
	//	g[x] = append(g[x], y)
	//	g[y] = append(g[y], x) // 建图
	//	deg[x]++
	//	deg[y]++ // 统计每个节点的度数（邻居个数）
	//}
	//fmt.Println(deg)
	//
	//leftEdges := n - 1 // 剩余边数
	//// 拓扑排序，去掉没有金币的子树
	//q := []int{}
	//for i, d := range deg {
	//	if d == 1 && coins[i] == 0 { // 没有金币的叶子
	//		q = append(q, i)
	//	}
	//}
	//fmt.Println(q)
	//for len(q) > 0 {
	//	x := q[len(q)-1]
	//	q = q[:len(q)-1]
	//	leftEdges-- // 删除节点 x 到其父节点的边
	//	for _, y := range g[x] {
	//		deg[y]--
	//		if deg[y] == 1 && coins[y] == 0 { // 没有金币的叶子
	//			q = append(q, y)
	//		}
	//	}
	//}
	//
	//// 再次拓扑排序
	//for i, d := range deg {
	//	if d == 1 && coins[i] > 0 { // 有金币的叶子（判断 coins[i] 是避免把没有金币的叶子也算进来）
	//		q = append(q, i)
	//	}
	//}
	//fmt.Println(q)
	//leftEdges -= len(q)   // 删除所有叶子（到其父节点的边）
	//for _, x := range q { // 遍历所有叶子
	//	for _, y := range g[x] {
	//		deg[y]--
	//		if deg[y] == 1 { // y 现在是叶子了
	//			leftEdges-- // 删除 y（到其父节点的边）
	//		}
	//	}
	//}
	//if leftEdges > 0 {
	//	return leftEdges << 1
	//}
	//return 0
}

//leetcode submit region end(Prohibit modification and deletion)
