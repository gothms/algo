package main

import "fmt"

func main() {
	s := "327"
	//s = "9999999999999" // 101
	//s = "999999"        // 11
	//s = "135135"        // 7
	combinations := numberOfCombinations(s)
	fmt.Println(combinations)

	//fmt.Println(len(s))
	//fmt.Println(1 << (len(s) - 1) % 1_000_000_007)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfCombinations(num string) int {
	// lc
	const mod int = 1e9 + 7
	if num[0] == '0' {
		return 0
	}
	ans := 0

	n := len(num)
	// 计算 lcp
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if num[i] == num[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}
	// 返回 s[l1:l2] <= s[l2:r2]
	lessEq := func(l1, l2, r2 int) bool {
		l := lcp[l1][l2]
		return l >= r2-l2 || num[l1+l] < num[l2+l]
	}

	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		f[0][j] = 1
	}
	for i := 1; i < n; i++ {
		if num[i] == '0' {
			continue
		}
		// k 和 j 同时向左向右扩展
		for j, k, sum := i, i-1, 0; j < n; j++ {
			f[i][j] = sum // 对应上面所说的长度小于最后一个划分出的整数
			if k < 0 {
				continue
			}
			if num[k] > '0' && lessEq(k, i, j+1) {
				f[i][j] = (f[i][j] + f[k][i-1]) % mod // 对应上面所说的长度等于最后一个划分出的整数
			}
			sum = (sum + f[k][i-1]) % mod
			k--
		}
	}
	for _, row := range f {
		ans = (ans + row[n-1]) % mod
	}
	return ans

	// dp + 前缀和：优化
	//if num[0] == '0' { // fast path
	//	return 0
	//}
	//const mod = 1_000_000_007
	//n := len(num)
	//dp := make([][]int, n)
	//for i := n - 1; i >= 0; i-- {
	//	dp[i] = make([]int, n+1)
	//	if num[i] == '0' { // 正整数不能前导 0
	//		continue
	//	}
	//	k := (n - i) >> 1 // 能取的最大位数，否则就大于下一个整数了
	//	for j := 1; j <= k; j++ {
	//		t := n - i - j                     // 用示例 s = "135135"，i=2 j=2 理解
	//		if num[i:i+j] <= num[i+j:i+j<<1] { // 比下一个相同位数的整数大
	//			dp[i][j] = (dp[i+j][t] - dp[i+j][min(j-1, t>>1)] + dp[i][j-1] + mod) % mod
	//		} else { // 否则下个整数的位数肯定比 num[i:i+j] 多 1
	//			if j == t { // j == n-i-j
	//				dp[i][j] = dp[i][j-1]
	//			} else {
	//				dp[i][j] = (dp[i+j][t] - dp[i+j][min(j, t>>1)] + dp[i][j-1] + mod) % mod
	//			}
	//		}
	//	}
	//	dp[i][n-i] = 1 + dp[i][k] // n-i：num[i:] 看做一个整数
	//	// (k,n-i) 之间的值都为 0，是为了优化前缀和，如果需要用到这个区间的值，则取索引 k 处的值
	//}
	//return dp[0][n] % mod

	// dp + 前缀和：个人
	//if num[0] == '0' {
	//	return 0
	//}
	//const mod = 1_000_000_007
	//n := len(num)
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, n+2)
	//}
	//for i := n - 1; i >= 0; i-- {
	//	if num[i] == '0' {
	//		continue
	//	}
	//	k := (n - i) >> 1
	//	for j := 1; j <= k; j++ {
	//		//t := (n - i - j) >> 1
	//		if num[i:i+j] <= num[i+j:i+j<<1] {
	//			dp[i][j] = (dp[i+j][n-i-j] - dp[i+j][j-1] + dp[i][j-1] + mod) % mod
	//			//dp[i][j] = (dp[i+j][n-i-j] - dp[i+j][min(j-1, t)] + dp[i][j-1] + mod) % mod
	//		} else {
	//			dp[i][j] = (dp[i+j][n-i-j] - dp[i+j][j] + dp[i][j-1] + mod) % mod
	//		}
	//	}
	//	for j := k; j <= n-i; j++ {
	//		dp[i][j] = dp[i][k]
	//	}
	//	dp[i][n-i]++
	//}
	//return dp[0][n]
}

//leetcode submit region end(Prohibit modification and deletion)
