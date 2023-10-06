//某个程序本来应该输出一个整数数组。但是这个程序忘记输出空格了以致输出了一个数字字符串，我们所知道的信息只有：数组中所有整数都在 [1, k] 之间，且数组中
//的数字都没有前导 0 。
//
// 给你字符串 s 和整数 k 。可能会有多种不同的数组恢复结果。
//
// 按照上述程序，请你返回所有可能输出字符串 s 的数组方案数。
//
// 由于数组方案数可能会很大，请你返回它对 10^9 + 7 取余 后的结果。
//
//
//
// 示例 1：
//
// 输入：s = "1000", k = 10000
//输出：1
//解释：唯一一种可能的数组方案是 [1000]
//
//
// 示例 2：
//
// 输入：s = "1000", k = 10
//输出：0
//解释：不存在任何数组方案满足所有整数都 >= 1 且 <= 10 同时输出结果为 s 。
//
//
// 示例 3：
//
// 输入：s = "1317", k = 2000
//输出：8
//解释：可行的数组方案为 [1317]，[131,7]，[13,17]，[1,317]，[13,1,7]，[1,31,7]，[1,3,17]，[1,3,1,7
//]
//
//
// 示例 4：
//
// 输入：s = "2020", k = 30
//输出：1
//解释：唯一可能的数组方案是 [20,20] 。 [2020] 不是可行的数组方案，原因是 2020 > 30 。 [2,020] 也不是可行的数组方案，因为
// 020 含有前导 0 。
//
//
// 示例 5：
//
// 输入：s = "1234567890", k = 90
//输出：34
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10^5.
// s 只包含数字且不包含前导 0 。
// 1 <= k <= 10^9.
//
//
// Related Topics 字符串 动态规划 👍 61 👎 0

package main

import (
	"fmt"
	"math"
)

func main() {
	s := "1234567890"
	k := 90
	s = "1317"
	k = 2000
	//s = "600342244431311113256628376226052681059918526204"
	//k = 703 // 411743991
	ofWays := numberOfArrays(s, k)
	fmt.Println(ofWays)
}

/*
dp
	c = 0：
		> k：返回 0
		<= k：dp[i] = dp[i-1]
	c > 0：
		遍历，直到 > k
*/

// leetcode submit region begin(Prohibit modification and deletion)
func numberOfArrays(s string, k int) int {
	// dp
	const mod = 1_000_000_007
	last, n := 0, len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	for i, c := range s {
		v, t := int(c-'0'), 10
		if v != 0 {
			last = i
		} else { // fast path
			t = int(math.Pow(10, float64(i-last))) // 10^x 也变了
			v = t * int(s[last]-'0')               // v * 100...
			if v > k {
				return 0 // fast path
			}
			t *= 10 // 补一位
		}
		for j := last; v <= k; t *= 10 {
			if s[j]-'0' > 0 { // 不能是 0 前导
				dp[i+1] = (dp[i+1] + dp[j]) % mod
			}
			if j == 0 { // 已遍历完 s
				break
			}
			v += int(s[j-1]-'0') * t
			j--
		}
	}
	return dp[n]

	// lc

	// 回溯：超时
	//cnt, n := 0, len(s)
	//var dfs func(int)
	//dfs = func(i int) {
	//	if i == n {
	//		cnt++
	//	}
	//	v := 0
	//	for j := i; j < n; j++ {
	//		v = v*10 + int(s[j]-'0')
	//		if v > k {
	//			break
	//		}
	//		if j+1 == n || s[j+1] != '0' {
	//			dfs(j + 1)
	//		}
	//	}
	//}
	//dfs(0)
	//return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
