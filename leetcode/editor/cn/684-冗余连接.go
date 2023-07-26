//树可以看成是一个连通且 无环 的 无向 图。
//
// 给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信
//息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。
//
// 请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的那个。
//
//
//
// 示例 1：
//
//
//
//
//输入: edges = [[1,2], [1,3], [2,3]]
//输出: [2,3]
//
//
// 示例 2：
//
//
//
//
//输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
//输出: [1,4]
//
//
//
//
// 提示:
//
//
// n == edges.length
// 3 <= n <= 1000
// edges[i].length == 2
// 1 <= ai < bi <= edges.length
// ai != bi
// edges 中无重复元素
// 给定的图是连通的
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 576 👎 0

package main

import "fmt"

func main() {
	// [1,3]
	edges := [][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}}
	// [2,8]
	edges = [][]int{{2, 7}, {7, 8}, {3, 6}, {2, 5}, {6, 8}, {4, 8}, {2, 8}, {1, 8}, {7, 10}, {3, 9}}
	connection := findRedundantConnection(edges)
	fmt.Println(connection)
}

/*
思路：dfs
	1.根据题意，树有 n 个节点，n-1 条边
		则任意添加一条边时，会形成一个环，并且仅一个
	2.尝试 dfs 找出找个环中所有的节点
		那么任意去掉这个环中的一条边，就会再次恢复为树
	3.关键点：找出环中的节点
		当dfs时遇到已经遍历过的节点，说明找到了环，此时在递归中返回这个节点
		当dfs的递归中“归”到这个节点时，说明环已经结束了
		这个归的过程遍历到的节点，就是环中所有的节点
思路：并查集
	1.任意一条边的两个节点，不属于同一个连通分量时，说明两个节点之前属于两个树
	2.如果属于同一连通分量，则说明两个节点之前已经是在同在同一个树中
		当前这条边就是多余的
*/
// leetcode submit region begin(Prohibit modification and deletion)
func findRedundantConnection(edges [][]int) []int {
	// dfs
	n := len(edges) + 1
	circle, adj, visited := make([]bool, n), make([][]int, n), make([]bool, n)
	for _, e := range edges { // 邻接表
		adj[e[0]], adj[e[1]] = append(adj[e[0]], e[1]), append(adj[e[1]], e[0])
	}
	var dfs func(int, int) (int, bool)
	dfs = func(p, c int) (r int, f bool) {
		visited[c] = true // 记录已访问的节点
		for _, i := range adj[c] {
			if i == p { // p是“父”节点
				continue
			}
			if visited[i] { // 找到环
				r, f, circle[c] = i, true, true
				return
			}
			r, f = dfs(c, i)
			if f {
				circle[c] = true // 记录环节点
				break
			}
		}
		if c == r { // 环终止
			return 0, false
		}
		return r, f // 也可以只返回r，通过 r>0 判断是否在环中
	}
	dfs(0, 1)
	for i := n - 2; i >= 0; i-- {
		if circle[edges[i][0]] && circle[edges[i][1]] { // 两个节点都在环中
			return edges[i]
		}
	}
	return nil

	// 并查集
	//n := len(edges)
	//uni := make([]int, n+1)
	//var find func(int) int
	//find = func(p int) int {
	//	if p == uni[p] {
	//		return p
	//	}
	//	for uni[p] != uni[uni[p]] { // 这种写法没有回溯功能
	//		uni[p] = uni[uni[p]]
	//	}
	//	//if p != uni[p] {
	//	//	uni[p] = find(uni[p])
	//	//}
	//	return uni[p]
	//}
	//union := func(p, q int) bool {
	//	rootP, rootQ := find(p), find(q)
	//	if rootP == rootQ { // 环出现了
	//		return false
	//	}
	//	uni[rootQ] = rootP // 合并两棵树
	//	return true
	//}
	//for i := 1; i <= n; i++ {
	//	uni[i] = i
	//}
	//for i := 0; i < n; i++ {
	//	if !union(edges[i][0], edges[i][1]) {
	//		return edges[i]
	//	}
	//	//fmt.Println(uni)
	//}
	//return nil
}

/*
两种写法的过程变化是一样

迭代写法
	[0 1 2 3 4 5 6 2 8 9 10]
	[0 1 2 3 4 5 6 2 2 9 10]
	[0 1 2 3 4 5 3 2 2 9 10]
	[0 1 2 3 4 2 3 2 2 9 10]
	[0 1 3 3 4 2 3 2 2 9 10]
	[0 1 3 4 4 2 3 2 3 9 10]
递归写法
	[0 1 2 3 4 5 6 2 8 9 10]
	[0 1 2 3 4 5 6 2 2 9 10]
	[0 1 2 3 4 5 3 2 2 9 10]
	[0 1 2 3 4 2 3 2 2 9 10]
	[0 1 3 3 4 2 3 2 2 9 10]
	[0 1 3 4 4 2 3 2 3 9 10]
*/
//leetcode submit region end(Prohibit modification and deletion)
