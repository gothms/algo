package main

/*
周赛 347：2023.5.27
	第三题：https://leetcode.cn/problems/minimum-cost-to-make-all-characters-equal/
思路：把 i 左侧都变成和 s[i] 一样（只讨论左半部分）
	1.状态转移方程
		1.1.dp[i]=dp[i-1]：s[i]==s[i-1]
			当 s[i]=s[i-1]：0 到 i-1 全部都是 s[i-1]
			所以 dp[i]=dp[i-1]
		1.2.dp[i]=dp[i-1]+i：s[i]!=s[i-1]
			当 s[i]!=s[i-1]：0 到 i-1 全部都是 s[i-1]
			所以 0 到 i-1 全部反转一次，次数等于 i
			dp[i]=dp[i-1]+i
	2.关键：
		2.1.为什么 s[i]!=s[i-1] 时，dp[i]=dp[i-1]+i 就是最小呢？
			也就是为什么要把 i 左侧都变成和 s[i] 一样
		2.2.举个直观例子：假设当前 i 所对应的字符串是 xxx01，设 x y 互反
			a)如果全变成 0：
				第一次：yyy10 = i+1
				第二次：xxx0 = i
				最后加上：xxx 全变成 0，假设需要最少 k 次
				总数：i+1+i+k
				且：全变 0 的最小次数 = i+1+i+k
			b)如果全变成 1：即为 xxx0 全变成 1
				第一次：yyy1 = i
				最后加上：yyy 全变成 1
					因为 xxx 全变成 0 需要 k 次，全反过来就是 1
					所以 yyy 全变成 1 不会超过 k+i-1 次
				总数：i+k+i-1
				且：全变 1 的最小次数 <= i+k+i-1
	3.总结：这个例子简单粗暴的证明了问题
		3.1.即 i 左侧的数，全部变为 s[i] 需要的次数，总是 小于 全部变为 s[i]的反转 需要的次数
			也就解释了 为什么 s[i]!=s[i-1] 时，dp[i]=dp[i-1]+i 最小
			dp[i]=dp[i-1]+i 中的 +i，就是：第一次：yyy1 = i
		3.2.另外，贴图中代码的写法，其实是把dp的两种状态（全变0/1），综合到一起了
			可以用两个状态来写这个题，就简单易懂了
*/
//leetcode submit region begin(Prohibit modification and deletion)
func minimumCost(s string) int64 {
	// dp：优化
	//minVal := func(a, b int) int {
	//	if a < b {
	//		return a
	//	}
	//	return b
	//}
	//min, n := int64(0), len(s)
	//for i := 1; i < n; i++ {
	//	if s[i] != s[i-1] {
	//		min += int64(minVal(i, n-i))
	//	}
	//}
	//return min

	// dp
	min := func(a, b int64) int64 {
		if a < b {
			return a
		}
		return b
	}
	n := len(s)
	dp0, dp1 := make([]int64, n), make([]int64, n)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			if s[i] == '0' {
				dp0[i] = dp0[i-1]
				dp1[i] = dp1[i-1] + 1
			} else {
				dp0[i] = dp0[i-1] + 1
				dp1[i] = dp1[i-1]
			}
		} else {
			dp0[i] = dp1[i-1] + min(int64(i), int64(n-i))
			dp1[i] = dp0[i-1] + min(int64(i), int64(n-i))
		}
	}
	return min(dp0[n-1], dp1[n-1])
}

//leetcode submit region end(Prohibit modification and deletion)
