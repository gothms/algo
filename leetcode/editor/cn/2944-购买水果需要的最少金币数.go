package main

import (
	"fmt"
	"slices"
)

func main() {
	prices := []int{3, 1, 2}                                                             //4
	prices = []int{1, 10, 1, 1}                                                          // 2
	prices = []int{26, 18, 6, 12, 49, 7, 45, 45}                                         // 39
	prices = []int{1, 37, 19, 38, 11, 42, 18, 33, 6, 37, 15, 48, 23, 12, 41, 18, 27, 32} // 37
	prices = []int{14, 37, 37, 38, 24, 15, 12}                                           // 63
	coins := minimumCoins(prices)
	fmt.Println(coins)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumCoins(prices []int) int {
	// 栈

	// dp

	// 记忆化搜索
}

//leetcode submit region end(Prohibit modification and deletion)

func minimumCoins_(prices []int) int {
	// 栈
	//n := len(prices)
	//q := [][2]int{{n + 1, 0}} // 哨兵
	//for i := n; i > 0; i-- {
	//	for q[0][0] > i<<1+1 { // > i<<1+1
	//		q = q[1:]
	//	}
	//	f := prices[i-1] + q[0][1] // 最小号金币数
	//	for f <= q[len(q)-1][1] {  // 递增栈：prices[i-1] + q[0][1] 必然大于栈底元素
	//		q = q[:len(q)-1]
	//	}
	//	q = append(q, [2]int{i, f})
	//}
	//return q[len(q)-1][1]

	// 递推 2
	n := len(prices)
	dp := make([]int, n+1) // +1 则 (i+1)<<1+1 不越界
	copy(dp[n>>1:], prices[n>>1:])
	for i := n>>1 - 1; i >= 0; i-- {
		dp[i] = prices[i] + slices.Min(dp[i+1:(i+1)<<1+1])
	}
	return dp[0]
	// 递推 1
	//n := len(prices)
	//for i := (n+1)>>1 - 1; i > 0; i-- {
	//	prices[i-1] += slices.Min(prices[i : i<<1+1])
	//}
	//return prices[0]

	// 记忆化
	//n := len(prices)
	//memo := make([]int, (n+1)>>1)
	//var dfs func(int) int
	//dfs = func(i int) int {
	//	if i<<1 >= n { // 可以跳过后面所有
	//		return prices[i-1]
	//	}
	//	v := &memo[i]
	//	if *v > 0 {
	//		return *v
	//	}
	//	ret := math.MaxInt32
	//  // 枚举下一个要买的水果 [i+1,i<<1+1]
	//	for j := i + 1; j <= i<<1+1; j++ { // j <= i<<1+1：可以跳过 i 个，但 i<<1+1 不能跳过
	//		ret = min(ret, dfs(j))
	//	}
	//	*v = ret + prices[i-1]
	//	return *v
	//}
	//return dfs(1)

	// dp：个人
	//n := len(prices)
	//dpBuy, dpFree := make([]int, n+1), make([]int, n+1)
	//dpBuy[0], dpBuy[1], dpFree[1] = prices[0], prices[0], prices[0]
	//buy, free := []int{1}, []int{1}
	//for i := 2; i <= n; i++ { // dp 的索引
	//	if buy[0]<<1 < i { // 不满足 i*2 >= j
	//		buy = buy[1:]
	//	}
	//	if free[0]<<1 < i {
	//		free = free[1:]
	//	}
	//	dpBuy[i] = min(dpBuy[(i)>>1], dpBuy[buy[0]]) + prices[i-1]
	//	for len(buy) > 0 && dpBuy[i] <= dpBuy[buy[len(buy)-1]] {
	//		buy = buy[:len(buy)-1] // 单调递增栈：保存 i*2 范围内的最少金币
	//	}
	//	buy = append(buy, i)
	//	dpFree[i] = min(dpBuy[(i+1)>>1], dpBuy[free[0]])
	//	for len(free) > 0 && dpBuy[i] <= dpBuy[free[len(free)-1]] {
	//		free = free[:len(free)-1]
	//	}
	//	free = append(free, i)
	//}
	//return min(dpBuy[n], dpFree[n])
}
