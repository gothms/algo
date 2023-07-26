//你是一位施工队的工长，根据设计师的要求准备为一套设计风格独特的房子进行室内装修。
//
// 房子的客厅大小为 n x m，为保持极简的风格，需要使用尽可能少的 正方形 瓷砖来铺盖地面。
//
// 假设正方形瓷砖的规格不限，边长都是整数。
//
// 请你帮设计师计算一下，最少需要用到多少块方形瓷砖？
//
//
//
// 示例 1：
//
//
//
// 输入：n = 2, m = 3
//输出：3
//解释：3 块地砖就可以铺满卧室。
//     2 块 1x1 地砖
//     1 块 2x2 地砖
//
// 示例 2：
//
//
//
// 输入：n = 5, m = 8
//输出：5
//
//
// 示例 3：
//
//
//
// 输入：n = 11, m = 13
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= n <= 13
// 1 <= m <= 13
//
//
// Related Topics 动态规划 回溯 👍 101 👎 0

package main

import (
	"fmt"
)

func main() {
	n, m := 2, 3 // 3
	n, m = 5, 8  // 5
	//n, m = 11, 13 // 6
	//n, m = 12, 11 // 7
	//n, m = 12, 13 // 7
	rectangle := tilingRectangle(n, m)
	fmt.Println(rectangle)

	//for i := 15; i >= 1; i-- {
	//	t := tilingRectangle(i, i-1)
	//	fmt.Println(i, t)
	//}
}

/*
思路：回溯 + 位运算
	1.示例 1 和 2，可以使用贪心，但是示例 3 不能使用贪心，假若使用贪心算法：
		1.1.n==m：f(n,m) = 1，即填充 1 个正方形
		1.2.n!=m：f(n,m) = 1+f(n,m-n)，假设 m>n 的，又由于 f(1,m) = m
			可见 f(n,m) 的下降速度，总是 大于等于 max(n,m) 的下降速度
		1.3.综上有 f(n,m)<=max(n,m)
	2.由给出的3个示例也可以看出，穷举的任一解都可能是最优解
		固使用穷举法，在 n*m 的矩阵中寻找可能的解法，求得需要的 最少 正方形数
	3.因为 n m 最大值 为 13，则可以记录 两个数组 r[] c[]——参考N皇后问题
		3.1.通过 r[i] c[j] 的 二进制位 标记 n*m 的矩阵 rect[i][j] 是否已被正方形覆盖
		3.2.下一个能放下的正方形的区间在 [j,m)
			a)设放下边长为 k 的正方形，总数 + 1，并且将 正方形k 覆盖的范围 二进制 标记 1
			b)放下 正方形k，寻找下一个正方形
			c)上一个分支寻找完后，通过 回溯 将 正方形k 覆盖的范围 二进制 置为 0
	4.优化：
		4.1.n==m 结果为 1
		4.2.i 行 [j,m) 区间，被覆盖的 j 都应被跳过
		4.3.贪心：i 行 [j,m) 区间中，优先放 k=m-j 的正方形，然后再放略小的 k-1 正方形
		4.4.i 行 [j,m) 区间都没被覆盖，则 (i,n) 行 [j,m) 区间也都没被覆盖
			则只用一个数组 r[i]，舍弃 c[j]
*/
//leetcode submit region begin(Prohibit modification and deletion)
func tilingRectangle(n int, m int) int {
	// 优化：终版
	if n == m {
		return 1
	}
	minVal := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	minTR, r := n, make([]int, n)
	if n < m {
		minTR = m // 初始化 最少 正方形数
	}
	var dfs func(int, int, int)
	dfs = func(i int, j int, cnt int) {
		if i == n { // 找到 更少 正方形数
			minTR = cnt
			return
		}
		for b := 1 << j; j < m && b&r[i] > 0; b <<= 1 { // 已被覆盖，跳过
			j++
		}
		if j == m { // 行尾
			dfs(i+1, 0, cnt)
			return
		}
		k := minVal(n-i, m-j)
		v := (1<<k - 1) << j
		if cnt++; cnt >= minTR || v&r[i] > 0 || k == 0 {
			return // 剪枝 & 贪心：检测 最大正方形k 能否放下
		}
		for idx := 0; idx < k; idx++ { // 标记 最大k 的 二进制位 为 1
			r[idx+i] ^= v
		}
		for b := 1 << (k - 1 + j); k > 0; b >>= 1 {
			dfs(i, j+k, cnt) // 放下 k
			k--
			for idx := 0; idx < k; idx++ {
				r[idx+i] ^= b // 回溯：j=k-1+j 的 二进制位 置为 0
			}
			r[k+i] ^= v // 回溯（最后一行）：i=k-1+i 行中，k覆盖的范围全置为 0
			v ^= b
		} // 由大到小，放下 k
	}
	dfs(0, 0, 0)
	return minTR

	// 优化
	//if n == m {
	//	return 1
	//}
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//minTR := n
	//if n < m {
	//	minTR = m // 初始化 最少 正方形数
	//}
	//r := make([]int, n)
	//tiling := func(i, j, k int) { // 正方形k 覆盖范围的 二进制 置为 1/0
	//	for idx, v := 0, (1<<k-1)<<j; idx < k; idx++ {
	//		r[idx+i] ^= v
	//	}
	//}
	//var dfs func(int, int, int)
	//dfs = func(i int, j int, cnt int) {
	//	if cnt >= minTR { // 剪枝
	//		return
	//	}
	//	if i == n { // 找到 更少 正方形数
	//		minTR = cnt
	//		return
	//	}
	//	for v := 1 << j; j < m && v&r[i] > 0; v <<= 1 { // 已被覆盖，跳过
	//		j++
	//	}
	//	if j == m { // 行尾
	//		dfs(i+1, 0, cnt)
	//		return
	//	}
	//	for k := minVal(n-i, m-j); k > 0 && (1<<k-1)<<j&r[i] == 0; k-- {
	//		if i == 4 && j == 7 {
	//			fmt.Println(i, j, k, r)
	//		}
	//		tiling(i, j, k)    // 标记 二进制位 为 1
	//		dfs(i, j+k, cnt+1) // 放下 k
	//		tiling(i, j, k)    // 回溯 二进制位 为 0
	//	} // 由大到小（贪心），检测 正方形k 能否放下
	//}
	//dfs(0, 0, 0)
	//return minTR

	// 位运算
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//minTR := n
	//if n < m {
	//	minTR = m // 初始化 最少 正方形数
	//}
	//r, c := make([]int, n), make([]int, m)
	//ok := func(i, j, k int) bool { // 检测边长为 k 的正方形，是否能放下
	//	return (1<<k-1)<<i&c[j] == 0 && (1<<k-1)<<j&r[i] == 0
	//}
	//tiling := func(i, j, k int) { // 正方形k 覆盖范围的 二进制 置为 1/0
	//	for idx := 0; idx < k; idx++ {
	//		c[idx+j] ^= (1<<k - 1) << i
	//		r[idx+i] ^= (1<<k - 1) << j
	//	}
	//}
	//var dfs func(int, int, int)
	//dfs = func(i int, j int, cnt int) {
	//	if cnt >= minTR { // 剪枝
	//		return
	//	}
	//	if i == n { // 找到 更少 正方形数
	//		minTR = cnt
	//		return
	//	}
	//	for j < m && 1<<i&c[j] > 0 { // 1<<i&c[j] 和 1<<j&r[i] 正"负"相同
	//		j++
	//	}
	//	if j == m { // 行尾
	//		dfs(i+1, 0, cnt)
	//		return
	//	}
	//	for k := minVal(n-i, m-j); k > 0 && ok(i, j, k); k-- {
	//		tiling(i, j, k)    // 标记 二进制位 为 1
	//		dfs(i, j+k, cnt+1) // 可以放下
	//		tiling(i, j, k)    // 回溯 二进制位 为 0
	//	} // 由大到小，检测 正方形k 能否放下
	//}
	//dfs(0, 0, 0)
	//return minTR

	// dp：错误
	//var recursion func(int, int) int
	//recursion = func(i, j int) int {
	//	if i == j {
	//		return 1
	//	}
	//	if i == 1 {
	//		return j
	//	}
	//	if j == 1 {
	//		return i
	//	}
	//	if i > j {
	//		return recursion(i-j, j) + 1
	//	} else {
	//		return recursion(i, j-i) + 1
	//	}
	//}
	//return recursion(n, m)

	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//if n > m {
	//	n, m = m, n
	//}
	//dp := make([][]int, n)
	//for i := 0; i < n; i++ {
	//	dp[i] = make([]int, m)
	//}
	//for i := 0; i < n; i++ {
	//	dp[i][0], dp[i][i] = i+1, 1
	//}
	//for i := 0; i < m; i++ {
	//	dp[0][i] = i + 1
	//}
	//for i := 1; i < n; i++ {
	//	for j := 1; j < m; j++ {
	//
	//	}
	//}
}

//leetcode submit region end(Prohibit modification and deletion)
