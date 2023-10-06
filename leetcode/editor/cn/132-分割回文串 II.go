//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
//
// 返回符合要求的 最少分割次数 。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//
//
// 示例 2：
//
//
//输入：s = "a"
//输出：0
//
//
// 示例 3：
//
//
//输入：s = "ab"
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 2000
// s 仅由小写英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 714 👎 0

package main

import "fmt"

func main() {
	s := "aab"
	cut := minCut(s)
	fmt.Println(cut)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCut(s string) int {
	// dp
	n := len(s)
	memo := make([][]bool, n)
	for i := n - 1; i >= 0; i-- {
		memo[i] = make([]bool, n)
		memo[i][i] = true
		for j := i + 1; j < n; j++ { // 标记回文串
			memo[i][j] = s[i] == s[j] && (i+1 == j || memo[i+1][j-1])
		}
	}
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if memo[0][i] { // [0,i] 为回文串
			continue
		}
		dp[i] = 2000 // 1 <= s.length <= 2000
		for j := 0; j < i; j++ {
			if memo[j+1][i] && dp[j]+1 < dp[i] { // 更新
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
