//假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
//
//
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//
//
// 示例 2:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
//
//
// 限制：
//
// 0 <= 数组长度 <= 10^5
//
//
//
// 注意：本题与主站 121 题相同：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-
//stock/
//
// Related Topics 数组 动态规划 👍 346 👎 0

package main

func main() {

}

/*
思路：dp
	1.只买卖一次股票的情况下，第 i 天的利润为：今天的股价 - [0,i) 天的最低股票价格
		stock：记录 [0,i) 天的最低股票价格
		cash：记录 [0,i) 天的最大利润
	2.状态转移方程：
		stock[i] = maxVal(stock[0,i), -prices[i])
			前 [0,i) 天的某一天股价更低，或今天的最低
		cash[i] = maxVal(cash[0,i), prices[i]+stock[i])
			前 [0,i) 天的某一天利润更高，或今天的利润更高
	3.此题只是可以买卖股票一次，其他变种有：
		lc 121 122 123 188 714，dp 思路都是相似的
*/
//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	// dp
	if len(prices) == 0 {
		return 0
	}
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	cash, stock, n := 0, -prices[0], len(prices)
	for i := 1; i < n; i++ {
		cash = maxVal(cash, prices[i]+stock)
		stock = maxVal(stock, -prices[i])
	}
	return cash
}

//leetcode submit region end(Prohibit modification and deletion)
