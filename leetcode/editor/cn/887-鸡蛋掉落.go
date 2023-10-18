//给你 k 枚相同的鸡蛋，并可以使用一栋从第 1 层到第 n 层共有 n 层楼的建筑。
//
// 已知存在楼层 f ，满足 0 <= f <= n ，任何从 高于 f 的楼层落下的鸡蛋都会碎，从 f 楼层或比它低的楼层落下的鸡蛋都不会破。
//
// 每次操作，你可以取一枚没有碎的鸡蛋并把它从任一楼层 x 扔下（满足 1 <= x <= n）。如果鸡蛋碎了，你就不能再次使用它。如果某枚鸡蛋扔下后没有摔碎
//，则可以在之后的操作中 重复使用 这枚鸡蛋。
//
// 请你计算并返回要确定 f 确切的值 的 最小操作次数 是多少？
//
// 示例 1：
//
//
//输入：k = 1, n = 2
//输出：2
//解释：
//鸡蛋从 1 楼掉落。如果它碎了，肯定能得出 f = 0 。
//否则，鸡蛋从 2 楼掉落。如果它碎了，肯定能得出 f = 1 。
//如果它没碎，那么肯定能得出 f = 2 。
//因此，在最坏的情况下我们需要移动 2 次以确定 f 是多少。
//
//
// 示例 2：
//
//
//输入：k = 2, n = 6
//输出：3
//
//
// 示例 3：
//
//
//输入：k = 3, n = 14
//输出：4
//
//
//
//
// 提示：
//
//
// 1 <= k <= 100
// 1 <= n <= 10⁴
//
//
// Related Topics 数学 二分查找 动态规划 👍 958 👎 0

package main

import (
	"fmt"
)

func main() {
	k, n := 1, 3 // 3
	//k, n = 2, 6  // 3
	//k, n = 2, 7  // 4
	k, n = 3, 14 // 4
	drop := superEggDrop(k, n)
	fmt.Println(drop)
}

// leetcode submit region begin(Prohibit modification and deletion)
func superEggDrop(k int, n int) int {
	// 滚动 dp
	if k == 1 { // fast path
		return n
	}
	dp, temp := make([]int, k+1), make([]int, k+1)
	for i := 1; i <= k; i++ {
		dp[i] = 1 // 初始化：操作上限为 1 次
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			temp[j] = dp[j-1] + dp[j] + 1 // 碎 / 不碎（+当前楼层）
		}
		if temp[k] >= n { // 最高楼层已达到 n
			return i
		}
		dp, temp = temp, dp
	}
	return dp[n] // 比如 k=1,n=1

	// 数学：dp
	// f(t,k) 表示操作 t 次，有 k 个鸡蛋，能找到的最高楼层 n，t 的操作上限是 n
	// f(t,1) = t，f(1,k) = 1
	// 碎 / 不碎（+当前楼层）：f(t,k) = f(t-1,k-1) + f(t-1,k)+1
	// 思路：
	// + 1：当前楼层
	// + dp[i-1][j]：对于 dp[i][j]，j 个鸡蛋操作 i-1 次肯定不会碎
	// + dp[i-1][j-1]：如果碎了，可以达到的楼层
	//dp := make([][]int, n+1)
	//for i := 1; i <= n; i++ {
	//	dp[i] = make([]int, k+1)
	//	dp[i][1] = i // 1 个鸡蛋
	//}
	//for i := 2; i <= k; i++ {
	//	dp[1][i] = 1 // k 个鸡蛋也只能操作 1 次
	//}
	//for i := 2; i <= n; i++ {
	//	for j := 1; j <= k; j++ {
	//		dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + 1 // 碎 / 不碎（+当前楼层）
	//	}
	//	if dp[i][k] >= n { // 最高楼层已达到 n
	//		return i
	//	}
	//}
	//return dp[n][k] // 比如 k=1,n=1

	// dp
	// 状态转移方程：二者取大
	// 碎：dp[i][k] = dp[j-1][k-1] + 1
	// 不碎：dp[i][k] = dp[i-j][k] + 1
	// 超时优化：二分查找代替 for j := 1; j <= i; j++
	//maxVal := func(a, b int) int {
	//	if b > a {
	//		return b
	//	}
	//	return a
	//}
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//dp := make([][]int, n+1)
	//dp[0] = make([]int, k+1)
	//for i := 1; i <= n; i++ {
	//	dp[i] = make([]int, k+1)
	//	dp[i][1] = i
	//	for c := 2; c <= k; c++ {
	//		// 迭代：超时
	//		//dp[i][c] = i
	//		//for j := 1; j <= i; j++ {
	//		//	dp[i][c] = minVal(dp[i][c], 1+maxVal(dp[j-1][c-1], dp[i-j][c])) // 碎 / 不碎
	//		//}
	//		// 二分查找：最接近的 l、r
	//		l, r := 1, i
	//		for l+1 < r {
	//			m := (l + r) >> 1
	//			v1, v2 := dp[m-1][c-1], dp[i-m][c] // 尽量让 碎/不碎 接近
	//			switch {
	//			case v1 > v2:
	//				r = m
	//			case v1 < v2:
	//				l = m
	//			default:
	//				l, r = m, m
	//			}
	//		}
	//		dp[i][c] = 1 + minVal( // 取小
	//			maxVal(dp[l-1][c-1], dp[i-l][c]),
	//			maxVal(dp[r-1][c-1], dp[i-r][c]))
	//	}
	//}
	//return dp[n][k]

	// 决策单调性：竞赛中的考点
	// 这里我们不会叙述 何为决策单调性 以及 如何根据决策单调性写出优化的动态规划，而是仅指出决策单调性的存在性
}

//leetcode submit region end(Prohibit modification and deletion)
