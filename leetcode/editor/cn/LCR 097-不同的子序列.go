//给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
//
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列
//，而 "AEC" 不是）
//
// 题目数据保证答案符合 32 位带符号整数范围。
//
//
//
// 示例 1：
//
//
//输入：s = "rabbbit", t = "rabbit"
//输出：3
//解释：
//如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
//rabbbit
//rabbbit
//rabbbit
//
// 示例 2：
//
//
//输入：s = "babgbag", t = "bag"
//输出：5
//解释：
//如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。
//babgbag
//babgbag
//babgbag
//babgbag
//babgbag
//
//
//
//
// 提示：
//
//
// 0 <= s.length, t.length <= 1000
// s 和 t 由英文字母组成
//
//
//
//
//
// 注意：本题与主站 115 题相同： https://leetcode-cn.com/problems/distinct-subsequences/
//
// Related Topics 字符串 动态规划 👍 67 👎 0

package main

import "fmt"

func main() {
	s := "rabbbit"
	t := "rabbit"
	//s = "babgbag"
	//t = "bag"
	distinct := numDistinct(s, t)
	fmt.Println(distinct)
}

// leetcode submit region begin(Prohibit modification and deletion)
func numDistinct(s string, t string) int {
	// lc：dp
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1 // t 为空字符串
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				// s[i] 匹配 t[j]：dp[i+1][j+1]
				// s[i] 不匹配 t[j]：dp[i+1][j]
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]

	// memo 优化
	//n := len(t)
	//memo := make(map[int32][]int, n+1)
	//for i, c := range s {
	//	memo[c] = append(memo[c], i)
	//}
	//last := memo[int32(t[0])]
	//dpLast := make([]int, len(last)+1)
	//for i := range dpLast {
	//	dpLast[i] = i
	//}
	//for i := 1; i < n; i++ {
	//	cur := memo[int32(t[i])]
	//	dpCur := make([]int, len(cur)+1)
	//	for j, k, m := 0, 0, len(last); j < len(cur); j++ {
	//		for k < m && last[k] < cur[j] {
	//			k++
	//		}
	//		dpCur[j+1] += dpCur[j] + dpLast[k] // 前缀和
	//	}
	//	last, dpLast = cur, dpCur
	//}
	//return dpLast[len(dpLast)-1]

	// memo
	//n := len(t)
	//memo := make(map[int32][]int, n+1) // 统计每个字符出现的位置
	//for i, c := range s {
	//	memo[c] = append(memo[c], i)
	//}
	//dp := make([][]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, len(memo[int32(t[i])])+1)
	//}
	//for i := range dp[0] {
	//	dp[0][i] = i // 初始化第一个字符
	//}
	//last := memo[int32(t[0])] // 上一个字符
	//for i := 1; i < n; i++ {
	//	cur := memo[int32(t[i])] // 当前字符
	//	for j, k, m := 0, 0, len(last); j < len(cur); j++ {
	//		for k < m && last[k] < cur[j] { // 当前字符的位置 cur[j] 之前出现了 k 个“上一个字符”
	//			k++
	//		}
	//		dp[i][j+1] += dp[i-1][k] // 状态转移方程
	//	}
	//	last = cur
	//}
	//ans := 0
	//for _, v := range dp[n-1] { // 统计
	//	ans += v
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
