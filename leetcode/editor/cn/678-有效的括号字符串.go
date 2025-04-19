package main

import "fmt"

func main() {
	s := "(*())"
	validString := checkValidString(s)
	fmt.Println(validString)
}

// leetcode submit region begin(Prohibit modification and deletion)
func checkValidString(s string) bool {
	// 贪心
	// 如果遇到左括号，则将最小值和最大值分别加 1
	// 如果遇到右括号，则将最小值和最大值分别减 1
	// 如果遇到星号，则将最小值减 1，将最大值加 1
	//var mnl, mxl int
	//for _, c := range s {
	//	switch c {
	//	case '(':
	//		mnl++
	//		mxl++
	//	case ')':
	//		mnl = max(mnl-1, 0)
	//		mxl--
	//		if mxl < 0 { // 在索引 i 处，( 已不够
	//			return false
	//		}
	//	case '*':
	//		mnl = max(mnl-1, 0)
	//		mxl++
	//	}
	//}
	//return mnl == 0 // ( 必须被完全匹配

	// 栈
	// 一个栈记录 (，一个栈记录 *
	// 当 ) 出现，优先匹配 (，其次匹配 *
	// 当遍历完 s，剩下的 ( 需要和 * 匹配完，且 ( 比 * 先出现

	// dp
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dp[n-1][n-1] = s[n-1] == '*' // 预处理
	for i := n - 2; i >= 0; i-- {
		l := s[i]
		if l == '*' { // 长度为 1
			dp[i][i] = true
		}
		r := s[i+1]
		dp[i][i+1] = (l == '(' || l == '*') && (r == ')' || r == '*') // 长度为 2
		for j := i + 2; j < n; j++ {                                  // 长度大于 2
			r = s[j]
			if (l == '(' || l == '*') && (r == ')' || r == '*') {
				dp[i][j] = dp[i+1][j-1]
			}
			for k := i; k < j && !dp[i][j]; k++ {
				dp[i][j] = dp[i][k] && dp[k+1][j]
			}
		}
	}
	return dp[0][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
