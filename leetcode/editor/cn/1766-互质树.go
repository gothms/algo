//给你一个 n 个节点的树（也就是一个无环连通无向图），节点编号从 0 到 n - 1 ，且恰好有 n - 1 条边，每个节点有一个值。树的 根节点 为 0
//号点。
//
// 给你一个整数数组 nums 和一个二维数组 edges 来表示这棵树。nums[i] 表示第 i 个点的值，edges[j] = [uj, vj] 表示节
//点 uj 和节点 vj 在树中有一条边。
//
// 当 gcd(x, y) == 1 ，我们称两个数 x 和 y 是 互质的 ，其中 gcd(x, y) 是 x 和 y 的 最大公约数 。
//
// 从节点 i 到 根 最短路径上的点都是节点 i 的祖先节点。一个节点 不是 它自己的祖先节点。
//
// 请你返回一个大小为 n 的数组 ans ，其中 ans[i]是离节点 i 最近的祖先节点且满足 nums[i] 和 nums[ans[i]] 是 互质的
//，如果不存在这样的祖先节点，ans[i] 为 -1 。
//
//
//
// 示例 1：
//
//
//
//
//输入：nums = [2,3,3,2], edges = [[0,1],[1,2],[1,3]]
//输出：[-1,0,0,1]
//解释：上图中，每个节点的值在括号中表示。
//- 节点 0 没有互质祖先。
//- 节点 1 只有一个祖先节点 0 。它们的值是互质的（gcd(2,3) == 1）。
//- 节点 2 有两个祖先节点，分别是节点 1 和节点 0 。节点 1 的值与它的值不是互质的（gcd(3,3) == 3）但节点 0 的值是互质的(gcd(
//2,3) == 1)，所以节点 0 是最近的符合要求的祖先节点。
//- 节点 3 有两个祖先节点，分别是节点 1 和节点 0 。它与节点 1 互质（gcd(3,2) == 1），所以节点 1 是离它最近的符合要求的祖先节点。
//
//
//
// 示例 2：
//
//
//
//
//输入：nums = [5,6,10,2,3,6,15], edges = [[0,1],[0,2],[1,3],[1,4],[2,5],[2,6]]
//输出：[-1,0,-1,0,0,0,-1]
//
//
//
//
// 提示：
//
//
// nums.length == n
// 1 <= nums[i] <= 50
// 1 <= n <= 10⁵
// edges.length == n - 1
// edges[j].length == 2
// 0 <= uj, vj < n
// uj != vj
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 数组 数学 数论 👍 33 👎 0

package main

import "fmt"

func main() {
	nums := []int{9, 16, 30, 23, 33, 35, 9, 47, 39, 46, 16, 38, 5, 49, 21, 44, 17, 1, 6, 37, 49, 15, 23, 46, 38, 9, 27, 3, 24, 1, 14, 17, 12, 23, 43, 38, 12, 4, 8, 17, 11, 18, 26, 22, 49, 14, 9}
	edges := [][]int{ // [-1,21,17,43,10,42,7,13,29,44,17,31,39,10,10,29,32,0,40,23,12,39,12,40,25,35,15,38,40,40,17,24,5,1,19,14,17,21,25,24,14,17,40,25,37,17,10]
		{17, 0},
		{30, 17},
		{41, 30},
		{10, 30},
		{13, 10},
		{7, 13},
		{6, 7},
		{45, 10},
		{2, 10},
		{14, 2},
		{40, 14},
		{28, 40},
		{29, 40},
		{8, 29},
		{15, 29},
		{26, 15},
		{23, 40},
		{19, 23},
		{34, 19},
		{18, 23},
		{42, 18},
		{5, 42},
		{32, 5},
		{16, 32},
		{35, 14},
		{25, 35},
		{43, 25},
		{3, 43},
		{36, 25},
		{38, 36},
		{27, 38},
		{24, 36},
		{31, 24},
		{11, 31},
		{39, 24},
		{12, 39},
		{20, 12},
		{22, 12},
		{21, 39},
		{1, 21},
		{33, 1},
		{37, 1},
		{44, 37},
		{9, 44},
		{46, 2},
		{4, 46}}
	coprimes := getCoprimes(nums, edges)
	fmt.Println(coprimes)
}

// leetcode submit region begin(Prohibit modification and deletion)
//func init() {
//	gcd := func(a, b int) int {
//		for a != 0 {
//			b, a = a, b%a
//		}
//		return b
//	}
//	g1766 = make([][]int, n1766)   // 互质列表
//	g1766[1] = append(g1766[1], 1) // 1 和 1
//	for i := 1; i < n1766; i++ {   // 包含 1
//		for j := i + 1; j < n1766; j++ {
//			if gcd(i, j) == 1 {
//				g1766[i] = append(g1766[i], j)
//				g1766[j] = append(g1766[j], i)
//			}
//		}
//	}
//	//for _, g := range g1766 {
//	//	fmt.Println(g)
//	//}
//}
//
//const n1766 = 51 // 值的区间为 [1,50]
//var g1766 [][]int
//
//func getCoprimes(nums []int, edges [][]int) []int {
//  // 已AC
//	n := len(nums)
//	var (
//		memo  = make([][]int, n1766) // 值为 i 的节点列表（节点编号）
//		adj   = make([][]int, n)     // 邻接表
//		depth = make([]int, n)       // 记录节点的深度
//		ret   = make([]int, n)
//	)
//	for i := 0; i < n; i++ {
//		ret[i], depth[i] = -1, -1
//	}
//	for _, e := range edges {
//		adj[e[0]] = append(adj[e[0]], e[1])
//		adj[e[1]] = append(adj[e[1]], e[0])
//	}
//	var dfs func(int, int)
//	dfs = func(i, d int) { // 节点编号，深度
//		depth[i] = d
//		v := nums[i]
//		for _, g := range g1766[v] { // v 和 g 互质，i 和 last 为编号
//			if len(memo[g]) > 0 {
//				last := memo[g][len(memo[g])-1]                  // 后入栈的值，肯定是最近的祖先
//				if ret[i] == -1 || depth[last] > depth[ret[i]] { // 更新不同的值：更近的祖先
//					ret[i] = last
//				}
//			}
//		}
//		memo[v] = append(memo[v], i)
//		for _, j := range adj[i] { // dfs
//			if depth[j] == -1 { // 未访问
//				dfs(j, d+1)
//			}
//		}
//		memo[v] = memo[v][:len(memo[v])-1] // 回溯
//	}
//	dfs(0, 0)
//	return ret

//n := len(nums)
//gcds := make([][]int, 51)
//tmp := make([][]int, 51)
//ans := make([]int, n)
//dep := make([]int, n)
//g := make([][]int, n)
//// 初始化
//for i := 0; i < 51; i++ {
//	gcds[i] = []int{}
//	tmp[i] = []int{}
//}
//for i := 0; i < n; i++ {
//	g[i] = []int{}
//	ans[i], dep[i] = -1, -1
//}
//
//var dfs func(x, depth int)
//dfs = func(x, depth int) {
//	dep[x] = depth
//	for _, val := range gcds[nums[x]] {
//		if len(tmp[val]) == 0 {
//			continue
//		}
//		las := tmp[val][len(tmp[val])-1]
//		if ans[x] == -1 || dep[las] > dep[ans[x]] {
//			//if ans[x] == -1 {
//			ans[x] = las
//			//break
//		}
//	}
//	tmp[nums[x]] = append(tmp[nums[x]], x)
//	for _, val := range g[x] {
//		if dep[val] == -1 { // 被访问过的点dep不为-1
//			dfs(val, depth+1)
//		}
//	}
//	tmp[nums[x]] = tmp[nums[x]][:len(tmp[nums[x]])-1]
//}
//gcd := func(a, b int) int {
//	for b != 0 {
//		a, b = b, a%b
//	}
//	return a
//}
//for i := 1; i <= 50; i++ {
//	for j := 1; j <= 50; j++ {
//		if gcd(i, j) == 1 {
//			gcds[i] = append(gcds[i], j)
//		}
//	}
//}
//for _, edge := range edges {
//	x := edge[0]
//	y := edge[1]
//	g[x] = append(g[x], y)
//	g[y] = append(g[y], x)
//}
//
//dfs(0, 1)
//return ans
//}

func init() {
	gcd := func(a, b int) bool {
		for a != 0 {
			b, a = a, b%a
		}
		return b == 1
	}
	g1766 = make([][]bool, n1766) // 互质列表
	for i := 1; i < n1766; i++ {
		g1766[i] = make([]bool, n1766)
	}
	g1766[1][1] = true
	for i := 1; i < n1766; i++ { // 包含 1
		for j := i + 1; j < n1766; j++ {
			g1766[i][j] = gcd(i, j)
			g1766[j][i] = g1766[i][j]
		}
	}
	//for _, g := range g1766 {
	//	fmt.Println(g)
	//}
}

const n1766 = 51 // 值的区间为 [1,50]
var g1766 [][]bool

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	var (
		set  = make(map[int]struct{}, n1766)
		memo = make([][][2]int, n1766) // 值为 i 的节点列表（节点编号）
		adj  = make([][]int, n)        // 邻接表
		vis  = make([]bool, n)
		ret  = make([]int, n)
	)
	for _, v := range nums { // 树中所有值的集合
		set[v] = struct{}{}
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	var dfs func(int, int)
	dfs = func(i, d int) { // 节点编号，深度
		if vis[i] {
			return
		}
		vis[i] = true
		r := [2]int{-1, -1} // 必须记录当前的深度
		x := nums[i]
		for y := range set { // 遍历所有值
			if len(memo[y]) > 0 && g1766[x][y] {
				last := memo[y][len(memo[y])-1]
				if last[1] > r[1] { // 更新深度
					r = last
				}
			}
		}
		ret[i] = r[0]
		memo[x] = append(memo[x], [2]int{i, d})
		for _, j := range adj[i] { // dfs
			dfs(j, d+1)
		}
		memo[x] = memo[x][:len(memo[x])-1] // 回溯
	}
	dfs(0, 0)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
