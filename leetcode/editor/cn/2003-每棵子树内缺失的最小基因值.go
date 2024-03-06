//有一棵根节点为 0 的 家族树 ，总共包含 n 个节点，节点编号为 0 到 n - 1 。给你一个下标从 0 开始的整数数组 parents ，其中
//parents[i] 是节点 i 的父节点。由于节点 0 是 根 ，所以 parents[0] == -1 。
//
// 总共有 10⁵ 个基因值，每个基因值都用 闭区间 [1, 10⁵] 中的一个整数表示。给你一个下标从 0 开始的整数数组 nums ，其中 nums[i]
// 是节点 i 的基因值，且基因值 互不相同 。
//
// 请你返回一个数组 ans ，长度为 n ，其中 ans[i] 是以节点 i 为根的子树内 缺失 的 最小 基因值。
//
// 节点 x 为根的 子树 包含节点 x 和它所有的 后代 节点。
//
//
//
// 示例 1：
//
//
//
// 输入：parents = [-1,0,0,2], nums = [1,2,3,4]
//输出：[5,1,1,1]
//解释：每个子树答案计算结果如下：
//- 0：子树包含节点 [0,1,2,3] ，基因值分别为 [1,2,3,4] 。5 是缺失的最小基因值。
//- 1：子树只包含节点 1 ，基因值为 2 。1 是缺失的最小基因值。
//- 2：子树包含节点 [2,3] ，基因值分别为 [3,4] 。1 是缺失的最小基因值。
//- 3：子树只包含节点 3 ，基因值为 4 。1是缺失的最小基因值。
//
//
// 示例 2：
//
//
//
// 输入：parents = [-1,0,1,0,3,3], nums = [5,4,6,2,1,3]
//输出：[7,1,1,4,2,1]
//解释：每个子树答案计算结果如下：
//- 0：子树内包含节点 [0,1,2,3,4,5] ，基因值分别为 [5,4,6,2,1,3] 。7 是缺失的最小基因值。
//- 1：子树内包含节点 [1,2] ，基因值分别为 [4,6] 。 1 是缺失的最小基因值。
//- 2：子树内只包含节点 2 ，基因值为 6 。1 是缺失的最小基因值。
//- 3：子树内包含节点 [3,4,5] ，基因值分别为 [2,1,3] 。4 是缺失的最小基因值。
//- 4：子树内只包含节点 4 ，基因值为 1 。2 是缺失的最小基因值。
//- 5：子树内只包含节点 5 ，基因值为 3 。1 是缺失的最小基因值。
//
//
// 示例 3：
//
// 输入：parents = [-1,2,3,0,2,4,1], nums = [2,3,4,5,6,7,8]
//输出：[1,1,1,1,1,1,1]
//解释：所有子树都缺失基因值 1 。
//
//
//
//
// 提示：
//
//
// n == parents.length == nums.length
// 2 <= n <= 10⁵
// 对于 i != 0 ，满足 0 <= parents[i] <= n - 1
// parents[0] == -1
// parents 表示一棵合法的树。
// 1 <= nums[i] <= 10⁵
// nums[i] 互不相同。
//
//
// Related Topics 树 深度优先搜索 并查集 动态规划 👍 154 👎 0

package main

import "fmt"

func main() {
	parents := []int{-1, 0, 1, 0, 3, 3}
	nums := []int{3, 4, 6, 2, 5, 1}
	subtree := smallestMissingValueSubtree(parents, nums)
	fmt.Println(subtree)
}

// leetcode submit region begin(Prohibit modification and deletion)
func smallestMissingValueSubtree(parents []int, nums []int) []int {
	f, cur, n := -1, 1, len(parents)
	ret, adj := make([]int, n), make([][]int, n)
	for i := 1; i < n; i++ {
		v := parents[i]
		adj[v] = append(adj[v], i)
	}
	for i := 0; i < n; i++ {
		if nums[i] == 1 { // 案例 3
			f = i
		}
		ret[i] = 1 // 初始化都是 1
	}
	minVals, visited := make(map[int]bool), make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if visited[i] {
			return
		}
		visited[i] = true       // 节点已访问
		minVals[nums[i]] = true // 值已出现
		for _, j := range adj[i] {
			dfs(j)
		}
	}
	for f != -1 {
		dfs(f)
		for minVals[cur] { // 寻找最小值
			cur++
		}
		ret[f], f = cur, parents[f] // 下一个节点
	}
	return ret

	// 个人：有误
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//n := len(parents)
	//adj, uni := make([][]int, n), make([]int, n+1)
	//ret := make([]int, n)
	//for i := 1; i < n; i++ {
	//	v := parents[i]
	//	adj[v] = append(adj[v], i)
	//	if nums[i] == 1 {
	//		ret[i] = 2
	//	} else {
	//		ret[i] = 1
	//	}
	//}
	//if nums[0] == 1 {
	//	ret[0] = 2
	//} else {
	//	ret[0] = 1
	//}
	//var dfs func(int)
	//dfs = func(i int) {
	//	v := nums[i]
	//	uni[v] = ret[i]
	//	for _, j := range adj[i] {
	//		dfs(j)
	//		uni[v] = maxVal(uni[v], uni[nums[j]])
	//		if uni[v] == v {
	//			uni[v]++
	//		}
	//		if uni[v] == nums[j] {
	//			uni[v]++
	//		}
	//	}
	//	ret[i] = uni[v]
	//}
	//fmt.Println(adj)
	//fmt.Println(uni)
	//dfs(0)
	//return ret

	// lc
	//n := len(parents)
	//ans := make([]int, n)
	//for i := range ans {
	//	ans[i] = 1
	//}
	//node := -1
	//for i, v := range nums {
	//	if v == 1 {
	//		node = i
	//		break
	//	}
	//}
	////node := slices.Index(nums, 1)
	//if node < 0 { // 不存在基因值为 1 的点
	//	return ans
	//}
	//
	//// 建树
	//g := make([][]int, n)
	//for i := 1; i < n; i++ {
	//	p := parents[i]
	//	g[p] = append(g[p], i)
	//}
	//
	//vis := make(map[int]bool, n)
	//var dfs func(int)
	//dfs = func(x int) {
	//	vis[nums[x]] = true // 标记基因值
	//	for _, son := range g[x] {
	//		if !vis[nums[son]] { // 避免重复访问节点
	//			dfs(son)
	//		}
	//	}
	//}
	//
	//for mex := 2; node >= 0; node = parents[node] {
	//	dfs(node)
	//	for vis[mex] { // node 子树包含这个基因值
	//		mex++
	//	}
	//	ans[node] = mex // 缺失的最小基因值
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
