//给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。 
//
// 返回 你可以获得的最大乘积 。 
//
// 
//
// 示例 1: 
//
// 
//输入: n = 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1。 
//
// 示例 2: 
//
// 
//输入: n = 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。 
//
// 
//
// 提示: 
//
// 
// 2 <= n <= 58 
// 
//
// Related Topics 数学 动态规划 👍 1241 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func integerBreak(n int) int {
    maxVal := func(a, b int) int {
        if a > b {
            return a
        }
        return b
    }
    dp3, dp2, dp1 := 0, 0, 1
    for i := 3; i <= n; i++ {
        dp1, dp2, dp3 = maxVal(maxVal((i-2)<<1, dp2<<1), maxVal((i-3)*3, dp3*3)), dp1, dp2
    }
    return dp1
}
//leetcode submit region end(Prohibit modification and deletion)
