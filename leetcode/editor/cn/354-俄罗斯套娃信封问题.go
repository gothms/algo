//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。
//
// 当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。
//
// 请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。
//
// 注意：不允许旋转信封。
//
// 示例 1：
//
//
//输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
//输出：3
//解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
//
// 示例 2：
//
//
//输入：envelopes = [[1,1],[1,1],[1,1]]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= envelopes.length <= 10⁵
// envelopes[i].length == 2
// 1 <= wi, hi <= 10⁵
//
//
// Related Topics 数组 二分查找 动态规划 排序 👍 1009 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	envelopes := [][]int{{1, 2},
		{2, 3},
		{3, 4},
		{3, 5},
		{4, 5},
		{5, 6},
		{5, 5},
		{5, 7},
		{6, 7},
		{7, 8}} // 7
	envelopes = [][]int{{46, 89},
		{50, 53},
		{52, 68},
		{72, 45},
		{77, 81}} // 3
	i := maxEnvelopes(envelopes)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxEnvelopes(envelopes [][]int) int {
	// dp：类似 lc-300 lc-1691
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] ||
			envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1] // 重点：envelopes[i][1] > envelopes[j][1]
	})
	dp := make([]int, 0)
	for _, e := range envelopes {
		j := sort.SearchInts(dp, e[1])
		if j == len(dp) {
			dp = append(dp, e[1]) // 套娃
		} else {
			dp[j] = e[1] // 替换为高度更小的信封（依赖于排序 envelopes[i][1] > envelopes[j][1]）
		}
	}
	return len(dp)

	// 无法二分：如数据 [[46 89] [50 53] [52 68] [72 45] [77 81]]
	//sort.Slice(envelopes, func(i, j int) bool {
	//	return envelopes[i][0] < envelopes[j][0] ||
	//		envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1]
	//})
	//var ans int
	//dp := make([]int, len(envelopes)+1)
	//for i, c := range envelopes {
	//	l, r := -1, i-1
	//	for l < r {
	//		m := (l + r + 1) >> 1
	//		if envelopes[m][0] >= c[0] || envelopes[m][1] >= c[1] {
	//			r = m - 1
	//		} else {
	//			l = m
	//		}
	//	}
	//	dp[i+1] = dp[l+1] + 1
	//	ans = max(ans, dp[i+1])
	//}
	//return ans

	// 个人：Time Limit Exceeded
	//sort.Slice(envelopes, func(i, j int) bool {
	//	return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] < envelopes[j][1]
	//})
	//var ans int
	//dp := make([]int, len(envelopes))
	//for i, c := range envelopes {
	//	for j := i - 1; j >= dp[i]; j-- {
	//		if envelopes[j][0] < c[0] && envelopes[j][1] < c[1] {
	//			dp[i] = max(dp[i], dp[j])
	//		}
	//	}
	//	dp[i]++
	//	ans = max(ans, dp[i])
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
