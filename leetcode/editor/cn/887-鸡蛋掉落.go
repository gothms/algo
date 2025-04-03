package main

import (
	"fmt"
)

func main() {
	k, n := 1, 3    // 3
	k, n = 2, 6     // 3
	k, n = 2, 7     // 4
	k, n = 3, 14    // 4
	k, n = 6, 10000 // 16
	//k, n = 9, 10000 // 14
	drop := superEggDrop(k, n)
	fmt.Println(drop)
}

// leetcode submit region begin(Prohibit modification and deletion)
func superEggDrop(k int, n int) int {
	dp := make([][]int, k+1)
	for j := range dp {
		dp[j] = make([]int, n+1)
	}
	for i := 0; i <= n; i++ {
		dp[1][i] = i
	}
	for j := 2; j <= k; j++ {
		for i := 1; i <= n; i++ {
			//v := math.MaxInt32
			//for d := 1; d <= i; d++ {
			//	v = min(v, 1+max(dp[j-1][d-1], dp[j][i-d]))
			//}
			//dp[j][i] = v
			l, r := 1, i
			for l < r {
				m := (l + r) >> 1
				v1, v2 := dp[j-1][m-1], dp[j][i-m]
				switch {
				case v1 < v2:
					l = m + 1
				case v1 > v2:
					r = m - 1
				default:
					l, r = m, m
				}
			}
			dp[j][i] = 1 + min(
				max(dp[j-1][l-1], dp[j][i-l]),
				max(dp[j-1][r-1], dp[j][i-r]))
		}
	}
	return dp[k][n]

	//memo := make([][]int, n+1)
	//for i := range memo {
	//	memo[i] = make([]int, k+1)
	//	memo[i][1] = i
	//}
	//for j := 1; j <= k; j++ {
	//	memo[1][j] = 1
	//}
	//var dfs func(int, int) int
	//dfs = func(i int, k int) int {
	//	v := &memo[i][k]
	//	if *v != 0 {
	//		return *v
	//	}
	//	ans := i
	//	for j := 1; j < i; j++ {
	//		ans = min(ans, max(dfs(j-1, k-1), dfs(i-j, k))+1)
	//	}
	//	*v = ans
	//	return ans
	//}
	//return dfs(n, k)
}

//leetcode submit region end(Prohibit modification and deletion)

func superEggDrop_(k int, n int) int {
	// 滚动 dp
	//if k == 1 { // fast path
	//	return n
	//}
	//dp, temp := make([]int, k+1), make([]int, k+1)
	//for i := 1; i <= k; i++ {
	//	dp[i] = 1 // 初始化：操作上限为 1 次
	//}
	//for i := 2; i <= n; i++ {
	//	for j := 1; j <= k; j++ {
	//		temp[j] = dp[j-1] + dp[j] + 1 // 碎 / 不碎（+当前楼层）
	//	}
	//	if temp[k] >= n { // 最高楼层已达到 n
	//		return i
	//	}
	//	dp, temp = temp, dp
	//}
	//return dp[n] // 比如 k=1,n=1

	// 数学：dp
	// f(t,k) 表示操作 t 次，有 k 个鸡蛋，能找到的最高楼层 n，t 的操作上限是 n
	// f(t,1) = t，f(1,k) = 1
	// 碎 / 不碎（+当前楼层）：f(t,k) = f(t-1,k-1) + f(t-1,k)+1
	// 思路：
	// + 1：当前楼层
	// + dp[i-1][j]：对于 dp[i][j]，j 个鸡蛋操作 i-1 次肯定不会碎
	// + dp[i-1][j-1]：如果碎了，可以达到的楼层
	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i // 1 个鸡蛋
	}
	for i := 2; i <= k; i++ {
		dp[1][i] = 1 // k 个鸡蛋也只能操作 1 次
	}
	for i := 2; i <= n; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + 1 // 碎 / 不碎（+当前楼层）
		}
		if dp[i][k] >= n { // 最高楼层已达到 n
			return i
		}
	}
	return dp[n][k] // 比如 k=1,n=1

	// dp
	// 状态转移方程：二者取大
	// 碎：dp[i][k] = dp[j-1][k-1] + 1
	// 不碎：dp[i][k] = dp[i-j][k] + 1
	// 超时优化：二分查找代替 for j := 1; j <= i; j++
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
	//		dp[i][c] = 1 + min( // 取小
	//			max(dp[l-1][c-1], dp[i-l][c]),
	//			max(dp[r-1][c-1], dp[i-r][c]))
	//	}
	//}
	//return dp[n][k]

	// 决策单调性：竞赛中的考点
}
