package main

import "math"

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func palindromePartition(s string, k int) int {
	n := len(s)
	memo := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		memo[i] = make([]int, n)
		for j := i + 1; j < n; j++ {
			memo[i][j] = memo[i+1][j-1] // s[i:j+1] 是回文串，需要修改的字符数
			if s[i] != s[j] {
				memo[i][j]++
			}
		}
	}
	dp := make([][]int, k)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0] = memo[0] // k=1
	for i := 1; i < k; i++ {
		for r := i; r < n; r++ { // 优化：r<=n-k+i
			dp[i][r] = 100
			for l := i; l <= r; l++ { // s[0:l] 分割为 i-1 个回文串，s[l:r+1] 为第 i 个回文串
				dp[i][r] = min(dp[i][r], dp[i-1][l-1]+memo[l][r])
			}
		}
	}
	return dp[k-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)

func palindromePartition_(s string, k int) int {
	n := len(s)
	memoChange := make([][]int, n)
	for i := range memoChange {
		memoChange[i] = make([]int, n)
		for j := range memoChange[i] {
			memoChange[i][j] = -1 // -1 表示没有计算过
		}
	}
	// 把 s[i:j+1] 改成回文串的最小修改次数
	var minChange func(int, int) int
	minChange = func(i, j int) int {
		if i >= j { // 子串只有一个字母，或者子串是空串
			return 0 // 无需修改
		}
		p := &memoChange[i][j]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := minChange(i+1, j-1)
		if s[i] != s[j] {
			res++
		}
		*p = res // 记忆化
		return res
	}

	memoDfs := make([][]int, k)
	for i := range memoDfs {
		memoDfs[i] = make([]int, n)
		for j := range memoDfs[i] {
			memoDfs[i][j] = -1 // -1 表示没有计算过
		}
	}
	// 把 s[:r+1] 切 i 刀，分成 i+1 个子串，每个子串改成回文串的最小总修改次数
	var dfs func(int, int) int
	dfs = func(i, r int) int {
		if i == 0 { // 只有一个子串
			return minChange(0, r)
		}
		p := &memoDfs[i][r]
		if *p != -1 { // 之前计算过
			return *p
		}
		res := math.MaxInt
		// 枚举子串左端点 l
		for l := i; l <= r; l++ {
			res = min(res, dfs(i-1, l-1)+minChange(l, r))
		}
		*p = res // 记忆化
		return res
	}

	return dfs(k-1, n-1)
}
