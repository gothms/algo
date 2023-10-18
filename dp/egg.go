package dp

/*
887
1884
*/

var tedDp []int

func init() {
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	N := 1000
	tedDp = make([]int, N+1)
	tedDp[1] = 1
	for i := 2; i <= N; i++ {
		for j := 1; j < i; j++ {
			tedDp[i] = minVal(tedDp[i], maxVal(j, tedDp[i-j]+1)) // 碎 / 不碎
		}
	}
}

func twoEggDrop(n int) int {
	return tedDp[n]

	// math
	// 让公式 1 + 2 + 3 + 4 + .... + x >= n 成立的最小的x，就是 f 的最小次数
	//return int(math.Ceil((math.Sqrt(float64(n*8+1)) - 1.0) / 2.0))
}

func superEggDrop(k int, n int) int {
	// dp
	// 状态转移方程：二者取大
	// 碎：dp[i][k] = dp[j-1][k-1] + 1
	// 不碎：dp[i][k] = dp[i-j][k] + 1
	// 超时优化：二分查找代替 for j := 1; j <= i; j++
	maxVal := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		dp[i][1] = i
		for c := 2; c <= k; c++ {
			// 迭代：超时
			//dp[i][c] = i
			//for j := 1; j <= i; j++ {
			//	dp[i][c] = minVal(dp[i][c], 1+maxVal(dp[j-1][c-1], dp[i-j][c])) // 碎 / 不碎
			//}
			// 二分查找：最接近的 l、r
			l, r := 1, i
			for l+1 < r {
				m := (l + r) >> 1
				v1, v2 := dp[m-1][c-1], dp[i-m][c] // 尽量让 碎/不碎 接近
				switch {
				case v1 > v2:
					r = m
				case v1 < v2:
					l = m
				default:
					l, r = m, m
				}
			}
			dp[i][c] = 1 + minVal( // 取小
				maxVal(dp[l-1][c-1], dp[i-l][c]),
				maxVal(dp[r-1][c-1], dp[i-r][c]))
		}
	}
	return dp[n][k]
}
func superEggDropMath(k int, n int) int {
	// math
	// 状态转移方程：f(t,k) 操作 t 次，有 k 个鸡蛋，能找到的最高楼层
	// t=[1,n]：f(t,1)=t，f(1,k)=1
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
}

// 决策单调性：竞赛中的考点
