package main

import (
	"fmt"
)

func main() {
	n, goal, k := 3, 5, 1 // 42
	n, goal, k = 5, 5, 3  // 120
	n, goal, k = 4, 5, 1  // 144
	n, goal, k = 4, 4, 2  // 24
	n, goal, k = 5, 7, 0  // 16800
	n, goal, k = 3, 4, 2  // 6
	playlists := numMusicPlaylists(n, goal, k)
	fmt.Println(playlists)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numMusicPlaylists(n int, goal int, k int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

func numMusicPlaylists_(n int, goal int, k int) int {
	// 容斥原理：b站左程云 099 节

	// 生成函数：第二类斯特林数
	const mod = 1_000_000_007
	pow := func(x, y, mod int) int {
		ret := 1
		for ; y > 0; y >>= 1 {
			if y&1 != 0 {
				ret = ret * x % mod
			}
			x = x * x % mod
		}
		return ret
	}
	inv := func(x int) int {
		return pow(x, mod-2, mod)
	}
	c := 1
	for x := 1; x < n-k; x++ {
		c = c * -x % mod
		if c < 0 {
			c += mod
		}
	}
	c = inv(c)

	ans := 0
	for i := 1; i <= n-k; i++ {
		ans += pow(i, goal-k-1, mod) * c % mod
		c = c * (i - (n - k)) % mod * inv(i) % mod
		if c < 0 {
			c += mod
		}
	}
	for i := 1; i <= n; i++ {
		ans = ans * i % mod
	}
	return ans

	//const mod = 1_000_000_007
	//dp := make([][]int, goal+1)
	//for i := range dp {
	//	dp[i] = make([]int, n+1)
	//}
	//dp[1][1] = n
	//for i := 2; i <= goal; i++ {
	//	for j := 1; j <= min(i, n); j++ {
	//		dp[i][j] = (dp[i-1][j-1]*(n-j+1)%mod + dp[i-1][j]*max(j-k, 0)%mod) % mod
	//	}
	//}
	//return dp[goal][n]
}

func numMusicPlaylists_my(n int, goal int, k int) int {
	// 个人
	// 设 f(n,g) 为长度为 g 的队列，含 n 个不同的数，总共有 f(n,g) 种不同的排列
	// 计算 n=5，goal=5
	// f(5,5) = 5^5 - C(5,4)*f(4,5) - C(5,3)*f(3,5) - C(5,2)*f(2,5) - C(5,1)*f(1,5)
	// f(1,x) = 1
	const mod = 1_000_000_007
	// 以下计算 k==0 的情况，验证是对的
	memo := make([][]int, goal+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
	}
	for i := 1; i <= goal; i++ {
		for j := 2; j <= min(i, n); j++ {
			// 令 f(g,n)，如 f(7,5) = (f(6,5) + f(6,4)) * 5
			memo[i][j] = (memo[i-1][j] + memo[i-1][j-1]) * j % mod // 状态转移方程
		}
		memo[i][1] = 1
	}
	// error：n, goal, k = 3, 4, 2  // 6
	dp := make([][]int, goal+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	for i := n; i <= goal; i++ {
		dp[i][0] = memo[i][n]
		for j := 1; j <= k; j++ {
			dp[i][j] = (dp[i][j-1] - dp[i-1][j-1]*n%mod + mod) % mod
		}
		fmt.Println(i, dp[i])
	}
	return dp[goal][k]
}
