//给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
//
//
// '.' 匹配任意单个字符
// '*' 匹配零个或多个前面的那一个元素
//
//
// 所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。
//
// 示例 1：
//
//
//输入：s = "aa", p = "a"
//输出：false
//解释："a" 无法匹配 "aa" 整个字符串。
//
//
// 示例 2:
//
//
//输入：s = "aa", p = "a*"
//输出：true
//解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
//
//
// 示例 3：
//
//
//输入：s = "ab", p = ".*"
//输出：true
//解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 20
// 1 <= p.length <= 20
// s 只包含从 a-z 的小写字母。
// p 只包含从 a-z 的小写字母，以及字符 . 和 *。
// 保证每次出现字符 * 时，前面都匹配到有效的字符
//
//
// Related Topics 递归 字符串 动态规划 👍 3894 👎 0

package main

import "fmt"

func main() {
	s := "aa"
	p := "a*" // true
	s = "aab"
	p = "c*a*b" // true
	s = "mississippi"
	p = "mis*is*p*." // false
	s = "a"
	p = ".*" // true
	//s = "ab"
	//p = ".*" // true
	match := isMatch(s, p)
	fmt.Println(match)
}

// leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
	// 递归：个人 AC
	var dfs func(int, int) bool
	dfs = func(i, j int) bool {
		if i == -1 && j == -1 { // 两个空字符串
			return true
		}
		if j < 0 { // s 非空，p 为空
			return false
		}
		switch p[j] {
		case '*': // 匹配 0 个字符 / 1 个字符
			return dfs(i, j-2) || i >= 0 && (s[i] == p[j-1] || p[j-1] == '.') && dfs(i-1, j)
		case '.':
			return i >= 0 && dfs(i-1, j-1)
		default:
			return i >= 0 && s[i] == p[j] && dfs(i-1, j-1)
		}
	}
	return dfs(len(s)-1, len(p)-1)

	//return isMatchRecursion(s, p, 0, 0)
}

func isMatchRecursion(s string, p string, i int, j int) bool {
	if j == len(p) {
		return i == len(s)
	}
	firstMatch := i < len(s) && s[i] == p[j] || p[j] == '.'
	if j+1 < len(p) && p[j+1] == '*' {
		return isMatchRecursion(s, p, i, j+2) || firstMatch && isMatchRecursion(s, p, i+1, j)
	}
	return firstMatch && isMatchRecursion(s, p, i+1, j+1)
}

//leetcode submit region end(Prohibit modification and deletion)

//func isMatch(s string, p string) bool {
//	// 优化
//	m, n := len(s), len(p)
//	dp := make([][]bool, m+1)
//	for i := range dp {
//		dp[i] = make([]bool, n+1)
//	}
//	dp[0][0] = true
//	for i := 0; i <= m; i++ { // i=0：空字符串和 p 匹配
//		for j := 1; j <= n; j++ {
//			switch p[j-1] {
//			case '*':
//				dp[i][j] = dp[i][j-2]                             // 匹配 0 个。j 肯定 >=2
//				if i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') { // 匹配 <0 个。但只考虑匹配 1 个，因为匹配 2 个已经被计算
//					dp[i][j] = dp[i][j] || dp[i-1][j]	// dp[i-1][j-2] 已被包含在 dp[i-1][j]
//				}
//			case '.':
//				dp[i][j] = i > 0 && dp[i-1][j-1]
//			default:
//				dp[i][j] = i > 0 && s[i-1] == p[j-1] && dp[i-1][j-1]
//			}
//		}
//	}
//	return dp[m][n]
//
//	// dp：个人 AC
//	//m, n := len(s), len(p)
//	//dp := make([][]bool, m+1)
//	//for i := range dp {
//	//	dp[i] = make([]bool, n+1)
//	//}
//	//dp[0][0] = true // 1.两个空字符串
//	//for j := 0; j < n; j++ {
//	//	dp[0][j+1] = p[j] == '*' && dp[0][j-1] // 2.空字符串匹配 p，比如 p = "c*a*b"
//	//}
//	//check := func(i, j int, v uint8, f func(k int) bool) {
//	//	for k := i; ; k-- {
//	//		if dp[k+1][j-1] {
//	//			dp[i+1][j+1] = true
//	//			break
//	//		}
//	//		f(k)
//	//	}
//	//}
//	//for i := 0; i < m; i++ {
//	//	for j := 0; j < n; j++ {
//	//		switch p[j] {
//	//		case '.':
//	//			dp[i+1][j+1] = dp[i][j]
//	//		case '*':
//	//			if v := p[j-1]; v == '.' {
//	//				//for k := i; ; k-- { // 3.k=-1 表示 s 是空字符串，k=i 表示 .* 匹配 0 个字符
//	//				//	if dp[k+1][j-1] {
//	//				//		dp[i+1][j+1] = true
//	//				//		break
//	//				//	}
//	//				//	if k < 0 {
//	//				//		break
//	//				//	}
//	//				//}
//	//				check(i, j, v, func(k int) bool {
//	//					return k < 0
//	//				})
//	//			} else {
//	//				//for k := i; ; k-- { // 4.同 3.
//	//				//	if dp[k+1][j-1] {
//	//				//		dp[i+1][j+1] = true
//	//				//		break
//	//				//	}
//	//				//	if k < 0 || s[k] != v { // 先匹配再检查 s[k] 是否为同字符 v
//	//				//		break
//	//				//	}
//	//				//}
//	//				check(i, j, v, func(k int) bool {
//	//					return k < 0 || s[k] != v
//	//				})
//	//			}
//	//		default:
//	//			dp[i+1][j+1] = s[i] == p[j] && dp[i][j]
//	//		}
//	//	}
//	//}
//	//return dp[m][n]
//}
