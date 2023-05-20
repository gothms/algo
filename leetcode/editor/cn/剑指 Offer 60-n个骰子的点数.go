//把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
//
//
//
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。
//
//
//
// 示例 1:
//
// 输入: 1
//输出: [0.16667,0.16667,0.16667,0.16667,0.16667,0.16667]
//
//
// 示例 2:
//
// 输入: 2
//输出: [0.02778,0.05556,0.08333,0.11111,0.13889,0.16667,0.13889,0.11111,0.08333,0
//.05556,0.02778]
//
//
//
// 限制：
//
// 1 <= n <= 11
//
// Related Topics 数学 动态规划 概率与统计 👍 554 👎 0

package main

import "fmt"

func main() {
	n := 5
	probability := dicesProbability(n)
	fmt.Println(probability)
}

//leetcode submit region begin(Prohibit modification and deletion)
func dicesProbability(n int) []float64 {
	// dp：
	dp := make([]float64, 6)
	for i := 0; i < 6; i++ {
		dp[i] = 1.0 / 6.0
	}
	for i := 2; i <= n; i++ {
		ndp := make([]float64, 5*i+1)
		for j, c := 0, 5*(i-1)+1; j < c; j++ {
			for k := 0; k < 6; k++ {
				ndp[j+k] += dp[j] / 6.0
			}
		}
		dp = ndp
	}
	return dp
}

//leetcode submit region end(Prohibit modification and deletion)
