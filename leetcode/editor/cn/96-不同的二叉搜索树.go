package main

import "fmt"

func main() {
	n := 2
	trees := numTrees(n)
	fmt.Println(trees)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numTrees(n int) int {
	// math：卡特兰数
	// https://baike.baidu.com/item/%E5%8D%A1%E7%89%B9%E5%85%B0%E6%95%B0/6125746?fromtitle=catalan&fromid=7605685
	c := 1
	for i := 1; i <= n; i++ {
		c = c * (i<<2 - 2) / (i + 1)
	}
	return c

	// dp
	//dp := make([]int, n+1)
	//dp[0], dp[1] = 1, 1
	//for i := 2; i <= n; i++ {
	//	if i&1 == 1 {
	//		dp[i] = dp[i>>1] * dp[i>>1]
	//	}
	//	for j := (i - 2) >> 1; j >= 0; j-- {
	//		dp[i] += dp[j] * dp[i-j-1] << 1
	//	}
	//}
	//return dp[n]

	// 记忆化搜索
	//memo := make([]int, n+1)
	//memo[0], memo[1] = 1, 1
	//var dfs func(int) int
	//dfs = func(i int) int {
	//	if memo[i] > 0 {
	//		return memo[i]
	//	}
	//	ret := 0
	//	if i&1 == 1 {
	//		d := dfs(i >> 1)
	//		ret = d * d
	//	}
	//	for j := (i - 2) >> 1; j >= 0; j-- {
	//		ret += dfs(j) * dfs(i-j-1) << 1
	//	}
	//	memo[i] = ret
	//	return ret
	//}
	//return dfs(n)

	// 旧写法
	//memo := make([]int, n+1)
	//memo[0], memo[1] = 1, 1
	//return numTreesRecursion(n, memo)
}

func numTreesRecursion(n int, memo []int) int {
	if memo[n] == 0 {
		for i := 0; i < n; i++ {
			memo[n] += numTreesRecursion(n-1-i, memo) * numTreesRecursion(i, memo)
		}
	}
	return memo[n]
}

//leetcode submit region end(Prohibit modification and deletion)
