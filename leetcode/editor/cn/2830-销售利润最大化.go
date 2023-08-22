//给你一个整数 n 表示数轴上的房屋数量，编号从 0 到 n - 1 。
//
// 另给你一个二维整数数组 offers ，其中 offers[i] = [starti, endi, goldi] 表示第 i 个买家想要以 goldi 枚
//金币的价格购买从 starti 到 endi 的所有房屋。
//
// 作为一名销售，你需要有策略地选择并销售房屋使自己的收入最大化。
//
// 返回你可以赚取的金币的最大数目。
//
// 注意 同一所房屋不能卖给不同的买家，并且允许保留一些房屋不进行出售。
//
//
//
// 示例 1：
//
//
//输入：n = 5, offers = [[0,0,1],[0,2,2],[1,3,2]]
//输出：3
//解释：
//有 5 所房屋，编号从 0 到 4 ，共有 3 个购买要约。
//将位于 [0,0] 范围内的房屋以 1 金币的价格出售给第 1 位买家，并将位于 [1,3] 范围内的房屋以 2 金币的价格出售给第 3 位买家。
//可以证明我们最多只能获得 3 枚金币。
//
// 示例 2：
//
//
//输入：n = 5, offers = [[0,0,1],[0,2,10],[1,3,2]]
//输出：10
//解释：有 5 所房屋，编号从 0 到 4 ，共有 3 个购买要约。
//将位于 [0,2] 范围内的房屋以 10 金币的价格出售给第 2 位买家。
//可以证明我们最多只能获得 10 枚金币。
//
//
//
// 提示：
//
//
// 1 <= n <= 10⁵
// 1 <= offers.length <= 10⁵
// offers[i].length == 3
// 0 <= starti <= endi <= n - 1
// 1 <= goldi <= 10³
//
//
// 👍 23 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 5
	offers := [][]int{{0, 2, 2}, {1, 3, 2}, {0, 0, 1}}
	n = 10 // 12
	offers = [][]int{{0, 6, 5}, {2, 9, 4}, {0, 9, 2}, {3, 9, 3}, {1, 6, 10}, {0, 1, 3}, {3, 8, 9}, {4, 8, 3}, {2, 6, 5}, {0, 4, 6}}
	profit := maximizeTheProfit(n, offers)
	fmt.Println(profit)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximizeTheProfit(n int, offers [][]int) int {
	// dp
	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Slice(offers, func(i, j int) bool {
		return offers[i][1] < offers[j][1]
	})
	ret, m := 0, len(offers)
	dp := make([]int, n+1)
	for i, j := 1, 0; i <= n; i++ {
		dp[i] = dp[i-1] // 每个值都需要记录
		for j < m && i > offers[j][1] {
			dp[i] = maxVal(dp[i], dp[offers[j][0]]+offers[j][2]) // offers[j] 不卖/卖
			j++
		}
		ret = maxVal(ret, dp[i])
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
