//给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
//
//
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
//
//
// 两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。
//
//
//
// 示例 1：
//
//
//输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace" ，它的长度为 3 。
//
//
// 示例 2：
//
//
//输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc" ，它的长度为 3 。
//
//
// 示例 3：
//
//
//输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= text1.length, text2.length <= 1000
// text1 和 text2 仅由小写英文字符组成。
//
//
//
//
//
// 注意：本题与主站 1143 题相同： https://leetcode-cn.com/problems/longest-common-
//subsequence/
//
// Related Topics 字符串 动态规划 👍 161 👎 0

package main

import "fmt"

func main() {
	text1 := "ezupkr"
	text2 := "ubmrapg" // 2
	subsequence := longestCommonSubsequence(text1, text2)
	fmt.Println(subsequence)
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp：优化
	n := len(text2)
	dp := make([]int, n+1) // 滚动数组
	for _, x := range text1 {
		var pre int // 记录 text2 中前一个字符
		for j, y := range text2 {
			if x == y {
				dp[j+1], pre = pre+1, dp[j+1]
			} else {
				pre = dp[j+1]
				dp[j+1] = max(dp[j], dp[j+1])
			}
		}
	}
	return dp[n]

	// dp
	//m, n := len(text1), len(text2)
	//dp := make([][]int, m+1)
	//for i := range dp {
	//	dp[i] = make([]int, n+1)
	//}
	//for i, x := range text1 {
	//	for j, y := range text2 {
	//		if x == y {
	//			dp[i+1][j+1] = max(dp[i][j]+1, dp[i][j+1], dp[i+1][j])
	//		} else {
	//			dp[i+1][j+1] = max(dp[i][j], dp[i][j+1], dp[i+1][j])
	//		}
	//	}
	//}
	//return dp[m][n]
}

//leetcode submit region end(Prohibit modification and deletion)
