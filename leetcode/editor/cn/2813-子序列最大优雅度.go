package main

import (
	"fmt"
	"sort"
)

func main() {
	items := [][]int{
		{3, 2},
		{5, 1},
		{10, 1}}
	k := 2
	items = [][]int{
		{3, 1},
		{3, 1},
		{2, 2},
		{5, 3}}
	k = 3
	items = [][]int{
		{3, 4},
		{8, 4},
		{2, 2},
		{1, 3}}
	k = 2 // 14
	items = [][]int{{10, 1},
		{10, 1},
		{10, 1},
		{10, 1},
		{10, 1},
		{10, 1},
		{10, 1},
		{10, 1},
		{10, 1},
		{10, 1},
		{3, 2},
		{3, 3},
		{3, 4},
		{3, 5},
		{3, 6},
		{3, 7},
		{3, 8},
		{3, 9},
		{3, 10},
		{3, 11}}
	k = 10 // 137
	elegance := findMaximumElegance(items, k)
	fmt.Println(elegance)
}

// leetcode submit region begin(Prohibit modification and deletion)
func findMaximumElegance(items [][]int, k int) int64 {
	// 贪心
	//sort.Slice(items, func(i, j int) bool { // 1.排序
	//	return items[i][0] > items[j][0]
	//})
	//memo, exItem := make(map[int]bool, k), make([]int, 0)
	//ans, cur := 0, 0
	//for _, item := range items[:k] { // 2.贪心，profit 最大的 k 个 item
	//	p, c := item[0], item[1]
	//	if !memo[c] {
	//		memo[c] = true // 记录已出现的 category
	//	} else {
	//		exItem = append(exItem, p) // 记录 category 出现次数超过 1 的 item
	//	}
	//	cur += p
	//}
	//cur += len(memo) * len(memo)
	//ans = cur
	//
	//for _, item := range items[k:] { // 3.寻找边界，即当 category*category > 贡献值超过 profit
	//	if len(exItem) == 0 {
	//		break
	//	}
	//	p, c := item[0], item[1]
	//	if !memo[c] { // 已出现的 category 不会增加贡献值
	//		cur += p - exItem[len(exItem)-1] + len(memo)<<1 + 1 // 关键
	//		memo[c] = true
	//		exItem = exItem[:len(exItem)-1] // 关键
	//		ans = max(ans, cur)
	//	}
	//}
	//return int64(ans)

	// dp
	sort.Slice(items, func(i, j int) bool {
		return items[i][1] < items[j][1]
	})
	dp := make([][3]int, k+1)
	for i, item := range items {
		p, c := item[0], item[1]
		for j := min(k, i+1); j > 0; j-- {
			cur := dp[j-1][0] + p
			has := dp[j-1][1] == c
			if !has {
				cur += dp[j-1][2]<<1 + 1
			}
			if cur > dp[j][0] {
				dp[j] = [3]int{cur, c, dp[j-1][2]}
				if !has {
					dp[j][2]++
				}
			}
		}
	}
	return int64(dp[k][0])
}

//leetcode submit region end(Prohibit modification and deletion)
