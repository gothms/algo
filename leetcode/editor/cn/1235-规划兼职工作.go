//你打算利用空闲时间来做兼职工作赚些零花钱。
//
// 这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。
//
// 给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。
//
// 注意，时间上出现重叠的 2 份工作不能同时进行。
//
// 如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。
//
//
//
// 示例 1：
//
//
//
// 输入：startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
//输出：120
//解释：
//我们选出第 1 份和第 4 份工作，
//时间范围是 [1-3]+[3-6]，共获得报酬 120 = 50 + 70。
//
//
// 示例 2：
//
//
//
// 输入：startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60
//]
//输出：150
//解释：
//我们选择第 1，4，5 份工作。
//共获得报酬 150 = 20 + 70 + 60。
//
//
// 示例 3：
//
//
//
// 输入：startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
//输出：6
//
//
//
//
// 提示：
//
//
// 1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4
// 1 <= startTime[i] < endTime[i] <= 10^9
// 1 <= profit[i] <= 10^4
//
//
// Related Topics 数组 二分查找 动态规划 排序 👍 415 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	startTime := []int{1, 2, 3, 4, 6}
	endTime := []int{3, 5, 10, 6, 9}
	profit := []int{20, 20, 100, 70, 60}
	startTime = []int{4, 2, 4, 8, 2}
	endTime = []int{5, 5, 5, 10, 8}
	profit = []int{1, 2, 8, 10, 4} // 18
	startTime = []int{6, 15, 7, 11, 1, 3, 16, 2}
	endTime = []int{19, 18, 19, 16, 10, 8, 19, 8}
	profit = []int{2, 9, 1, 19, 5, 7, 3, 19} // 41
	scheduling := jobScheduling(startTime, endTime, profit)
	fmt.Println(scheduling)
}

// leetcode submit region begin(Prohibit modification and deletion)
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	// 优化
	ans, n := 0, len(startTime)
	dp, idx := make([]int, n+1), make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return endTime[idx[i]] < endTime[idx[j]] // 按结束时间排序
	})
	for i, v := range idx {
		k := sort.Search(n, func(j int) bool {
			return endTime[idx[j]] > startTime[v] // 重要逻辑
		}) // 第一个大于 startTime[v] 的 idx 的索引，再 -1 就是 <= startTime[v]
		dp[i+1] = max(dp[i], dp[k]+profit[v]) // 引入了哨兵
		ans = max(ans, dp[i+1])
	}
	return ans

	// 排序 + 二分
	//ans, n := 0, len(startTime)
	//dp, idx := make([]int, n), make([]int, n)
	//for i := range idx {
	//	idx[i] = i
	//}
	//sort.Slice(idx, func(i, j int) bool {
	//	return endTime[idx[i]] < endTime[idx[j]] // 按结束时间排序
	//	//return endTime[idx[i]] < endTime[idx[j]] ||
	//	//endTime[idx[i]] == endTime[idx[j]] && startTime[idx[i]] < startTime[idx[j]]
	//})
	//for i, v := range idx {
	//	k := sort.Search(n, func(j int) bool {
	//		return endTime[idx[j]] > startTime[v]	// 重要逻辑
	//	}) - 1 // 第一个大于 startTime[v] 的 idx 的索引，再 -1 就是 <= startTime[v]
	//	if k < 0 {
	//		dp[i] = profit[v]
	//	} else {
	//		dp[i] = dp[k] + profit[v]
	//	}
	//	if i > 0 {
	//		dp[i] = max(dp[i], dp[i-1]) // 有可能 dp[i-1] 更大
	//	}
	//	ans = max(ans, dp[i])
	//}
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
