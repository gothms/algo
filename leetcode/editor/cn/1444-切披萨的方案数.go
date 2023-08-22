//给你一个 rows x cols 大小的矩形披萨和一个整数 k ，矩形包含两种字符： 'A' （表示苹果）和 '.' （表示空白格子）。你需要切披萨 k-1
// 次，得到 k 块披萨并送给别人。
//
// 切披萨的每一刀，先要选择是向垂直还是水平方向切，再在矩形的边界上选一个切的位置，将披萨一分为二。如果垂直地切披萨，那么需要把左边的部分送给一个人，如果水平
//地切，那么需要把上面的部分送给一个人。在切完最后一刀后，需要把剩下来的一块送给最后一个人。
//
// 请你返回确保每一块披萨包含 至少 一个苹果的切披萨方案数。由于答案可能是个很大的数字，请你返回它对 10^9 + 7 取余的结果。
//
//
//
// 示例 1：
//
//
//
// 输入：pizza = ["A..","AAA","..."], k = 3
//输出：3
//解释：上图展示了三种切披萨的方案。注意每一块披萨都至少包含一个苹果。
//
//
// 示例 2：
//
// 输入：pizza = ["A..","AA.","..."], k = 3
//输出：1
//
//
// 示例 3：
//
// 输入：pizza = ["A..","A..","..."], k = 1
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= rows, cols <= 50
// rows == pizza.length
// cols == pizza[i].length
// 1 <= k <= 10
// pizza 只包含字符 'A' 和 '.' 。
//
//
// Related Topics 记忆化搜索 数组 动态规划 矩阵 👍 92 👎 0

package main

import "fmt"

func main() {
	pizza := []string{
		"A..",
		"AAA",
		"...",
	}
	k := 3
	pizza = []string{
		".A..A",
		"A.A..",
		"A.AA.",
		"AAAA.",
		"A.AA.",
	}
	k = 5 // 153
	i := ways(pizza, k)
	fmt.Println(i)
}

/*
思路：记忆化搜索
	1.回溯：
		使用回溯思想，尝试在任意位置横切/竖切，我们会遍历所有的组合可能
		从这些组合中，排除不合法的，就是可切的方案
		排除的思路是，一刀下去后，上/下都有苹果，或者左右都有苹果
	2.记忆化
		在回溯思想中，常使用记忆化搜索，来记录已经搜索过的情况
		用三维数组 memo 记录已搜索过的结果，第三维是切 ki 刀的结果，ki = [0,k]
	3.其他细节
		使用 cache 数组记录矩阵中苹果的数量，如 cache[i][j] 记录了 [0,0] 和 [i,j] 矩阵
		其他细节请看代码
思路：dp
	1.记忆化搜索方法中，使用自顶向下的编程方式，dp中自底向上编程会简化矩阵中苹果数量的判断
		其他思路和记忆化搜索时类似的
	2.状态转移方程
		dp[i][j][ki] = sum(dp[x][j][ki-1]) + sum(dp[i][y][ki-1])，其中 x=[i+1,n)，y=[j+1,n)
		累加的前提是：cache[x][j] < cache[i][j] 或 cache[i][y] < cache[i][j]
		即：一刀下去后，上/下都有苹果，或者左右都有苹果
*/
// leetcode submit region begin(Prohibit modification and deletion)
func ways(pizza []string, k int) int {
	// dp：自底向上
	const mod = 1_000_000_007
	n, m := len(pizza), len(pizza[0])
	dp, cache := make([][][]int, n), make([][]int, n+1)
	cache[n] = make([]int, m+1)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, m)
		cache[i] = make([]int, m+1)
		for j := 0; j < m; j++ {
			dp[i][j] = make([]int, k+1)
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			cache[i][j] = cache[i+1][j] + cache[i][j+1] - cache[i+1][j+1]
			if pizza[i][j] == 'A' {
				cache[i][j]++ // 预处理矩阵和
			}
			if cache[i][j] > 0 {
				dp[i][j][0] = 1 // 初始化最后一刀
			}
		}
	}
	for i := n - 1; i >= 0; i-- { // 自底向上
		for j := m - 1; j >= 0; j-- {
			for p := 1; p < k; p++ {
				ni, nj := i+1, j+1
				for ni < n && cache[ni][j] == cache[i][j] { // 淘汰苹果为空的行
					ni++
				}
				for nj < m && cache[i][nj] == cache[i][j] { // 淘汰苹果为空的列
					nj++
				}
				for ; ni < n; ni++ { // 切掉每一行
					dp[i][j][p] = (dp[i][j][p] + dp[ni][j][p-1]) % mod
				}
				for ; nj < m; nj++ { // 切掉每一列
					dp[i][j][p] = (dp[i][j][p] + dp[i][nj][p-1]) % mod
				}
			}
		}
	}
	return dp[0][0][k-1]

	// 记忆化搜索：自顶向下
	//const mod = 1000000007
	//n, m := len(pizza), len(pizza[0])
	//memo, cache := make([][][]int, n+1), make([][]int, n+1)
	//for i := n; i >= 0; i-- {
	//	memo[i] = make([][]int, m+1) // 记忆化搜索的容器
	//	cache[i] = make([]int, m+1)  // 记录苹果矩阵和
	//	for j := m; j >= 0; j-- {
	//		memo[i][j] = make([]int, k+1)
	//	}
	//}
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= m; j++ { // 预处理矩阵和
	//		cache[i][j] = cache[i-1][j] + cache[i][j-1] - cache[i-1][j-1]
	//		if pizza[i-1][j-1] == 'A' {
	//			cache[i][j]++
	//		}
	//		if cache[i][j] > 0 {
	//			memo[i][j][1] = 1 // 初始化最后一刀
	//		}
	//	}
	//}
	//ok := func(x, y, r, c int) bool { // 判断矩阵 [r,c] 和 [x,y] 的苹果差值
	//	return cache[r][c]+cache[x-1][y-1] > cache[r][y-1]+cache[x-1][c]
	//}
	//var dfs func(int, int, int) int
	//dfs = func(x int, y int, k int) int {
	//	if k == 1 { // 最后一刀
	//		//if cache[n][m]+cache[x-1][y-1] > cache[n][y-1]+cache[x-1][m] {
	//		if ok(x, y, n, m) {
	//			return 1 // 最后一刀有苹果，方案可行
	//		}
	//		return 0
	//	}
	//	var v *int
	//	if v = &memo[x][y][k]; *v > 0 {
	//		return *v // 已搜索过，直接返回值
	//	}
	//	for x < n && !ok(x, y, x, m) { // 淘汰苹果为空的行
	//		x++
	//	}
	//	for y < m && !ok(x, y, n, y) { // 淘汰苹果为空的列
	//		y++
	//	}
	//	cnt, i, j := 0, x, y
	//	for ; i < n; i++ { // 切掉每一行
	//		cnt = (cnt + dfs(i+1, y, k-1)) % mod
	//	}
	//	for ; j < m; j++ { // 切掉每一列
	//		cnt = (cnt + dfs(x, j+1, k-1)) % mod
	//	}
	//	*v = cnt // 记忆化
	//	return cnt
	//}
	//return dfs(1, 1, k) // 从1,1开始搜索（因为cache的初始化方式）

	// 优化前
	//var dfs func(int, int, int) int
	//dfs = func(x int, y int, k int) int {
	//	if k == 1 {
	//		//if cache[n][m]+cache[x-1][y-1] > cache[n][y-1]+cache[x-1][m] {
	//		if ok(x, y, n, m) {
	//			return 1
	//		}
	//		return 0
	//	}
	//	var v *int
	//	if v = &memo[x][y][k]; *v > 0 {
	//		return *v
	//	}
	//	cnt, i, j := 0, x, y
	//	//for i < n && cache[i][m]+cache[i-1][j-1] == cache[i][j-1]+cache[i-1][m] {
	//	for i < n && !ok(i, j, i, m) {
	//		i++
	//	}
	//	//for j < m && cache[n][j]+cache[i-1][j-1] == cache[i-1][j]+cache[n][j-1] {
	//	for j < m && !ok(i, j, n, j) {
	//		j++
	//	}
	//	for ; i < n; i++ {
	//		cnt = (cnt + dfs(i+1, j, k-1)) % mod
	//	}
	//	for ; j < m; j++ {
	//		cnt = (cnt + dfs(x, j+1, k-1)) % mod
	//	}
	//	*v = cnt
	//	return cnt
	//}
	//return dfs(1, 1, k)
}

//leetcode submit region end(Prohibit modification and deletion)
