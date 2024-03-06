//给你两个数字字符串 num1 和 num2 ，以及两个整数 max_sum 和 min_sum 。如果一个整数 x 满足以下条件，我们称它是一个好整数：
//
//
// num1 <= x <= num2
// min_sum <= digit_sum(x) <= max_sum.
//
//
// 请你返回好整数的数目。答案可能很大，请返回答案对 10⁹ + 7 取余后的结果。
//
// 注意，digit_sum(x) 表示 x 各位数字之和。
//
//
//
// 示例 1：
//
//
//输入：num1 = "1", num2 = "12", min_num = 1, max_num = 8
//输出：11
//解释：总共有 11 个整数的数位和在 1 到 8 之间，分别是 1,2,3,4,5,6,7,8,10,11 和 12 。所以我们返回 11 。
//
//
// 示例 2：
//
//
//输入：num1 = "1", num2 = "5", min_num = 1, max_num = 5
//输出：5
//解释：数位和在 1 到 5 之间的 5 个整数分别为 1,2,3,4 和 5 。所以我们返回 5 。
//
//
//
//
// 提示：
//
//
// 1 <= num1 <= num2 <= 10²²
// 1 <= min_sum <= max_sum <= 400
//
//
// Related Topics 数学 字符串 动态规划 👍 34 👎 0

package main

import (
	"fmt"
	"strings"
)

func main() {
	num1, num2 := "4179205230", "7748704426"
	min, max := 8, 46 // 883045899
	i := count(num1, num2, min, max)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func count(num1 string, num2 string, min_sum int, max_sum int) int {
	const mod = 1_000_000_007
	n := len(num2)
	num1 = strings.Repeat("0", n-len(num1)) + num1 // 补前导零，和 num2 对齐

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, min(9*n, max_sum)+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(int, int, bool, bool) int
	dfs = func(i, sum int, limitLow, limitHigh bool) (res int) {
		if sum > max_sum { // 非法
			return
		}
		if i == n {
			if sum >= min_sum { // 合法
				res = 1
			}
			return
		}
		if !limitLow && !limitHigh {
			p := &memo[i][sum]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		lo := 0
		if limitLow {
			lo = int(num1[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(num2[i] - '0')
		}
		for d := lo; d <= hi; d++ { // 枚举当前数位填 d
			res = (res + dfs(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)) % mod
		}
		return
	}
	return dfs(0, 0, true, true)
}

//leetcode submit region end(Prohibit modification and deletion)
