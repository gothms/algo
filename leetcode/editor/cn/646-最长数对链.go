package main

import (
	"math"
	"sort"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func findLongestChain(pairs [][]int) int {
	// 贪心：lc
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ans, r := 0, math.MinInt32
	for _, p := range pairs {
		if r < p[0] {
			r = p[1]
			ans++
		}
	}
	return ans

	// 最长递增子序列：lc-300
	//sort.Slice(pairs, func(i, j int) bool {
	//	return pairs[i][0] < pairs[j][0]
	//})
	//sub := []int{-1001}
	//for _, p := range pairs {
	//	l, r := p[0], p[1]
	//	if l > sub[len(sub)-1] {
	//		sub = append(sub, r)
	//		continue
	//	}
	//	i := sort.Search(len(sub), func(i int) bool {
	//		return sub[i] >= l
	//	})
	//	sub[i] = min(sub[i], r)
	//}
	//return len(sub) - 1

	// dp：个人
	//sort.Slice(pairs, func(i, j int) bool {
	//	return pairs[i][0] < pairs[j][0]
	//})
	//n := len(pairs)
	//dp := make([]int, n)
	//dp[0] = 1
	//for i := 1; i < n; i++ {
	//	j := i - 1
	//	for j >= 0 && pairs[i][0] <= pairs[j][1] {
	//		j--
	//	}
	//	if j >= 0 {
	//		dp[i] = dp[j] + 1
	//	} else {
	//		dp[i] = 1
	//	}
	//}
	//return dp[n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
