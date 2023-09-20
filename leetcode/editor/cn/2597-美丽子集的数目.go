//给你一个由正整数组成的数组 nums 和一个 正 整数 k 。
//
// 如果 nums 的子集中，任意两个整数的绝对差均不等于 k ，则认为该子数组是一个 美丽 子集。
//
// 返回数组 nums 中 非空 且 美丽 的子集数目。
//
// nums 的子集定义为：可以经由 nums 删除某些元素（也可能不删除）得到的一个数组。只有在删除元素时选择的索引不同的情况下，两个子集才会被视作是不同的
//子集。
//
//
//
// 示例 1：
//
// 输入：nums = [2,4,6], k = 2
//输出：4
//解释：数组 nums 中的美丽子集有：[2], [4], [6], [2, 6] 。
//可以证明数组 [2,4,6] 中只存在 4 个美丽子集。
//
//
// 示例 2：
//
// 输入：nums = [1], k = 1
//输出：1
//解释：数组 nums 中的美丽数组有：[1] 。
//可以证明数组 [1] 中只存在 1 个美丽子集。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 20
// 1 <= nums[i], k <= 1000
//
//
// Related Topics 数组 动态规划 回溯 👍 36 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 4, 6}
	k := 2
	nums = []int{10, 4, 5, 7, 2, 1} // 23
	k = 3
	//nums = []int{4, 2, 5, 9, 10, 3} // 23
	//k = 1
	i := beautifulSubsets(nums, k)
	fmt.Println(i)
}

// leetcode submit region begin(Prohibit modification and deletion)
func beautifulSubsets(nums []int, k int) int {
	//absVal := func(a int) int {
	//	if a < 0 {
	//		return -a
	//	}
	//	return a
	//}
	//cnt, n := 0, len(nums)
	//dp := make([]int, n)
	////sort.Ints(nums) // 避免 [4,2,5]，k=1 情况下，数据污染
	//for i, v := range nums {
	//	for j := 0; j < i; j++ {
	//		if absVal(v-nums[j]) != k {
	//			dp[i] += dp[j]
	//			//} else {
	//			//	dp[i] -= dp[j]
	//		}
	//	}
	//	dp[i]++
	//	cnt += dp[i]
	//}
	//fmt.Println(dp)
	//return cnt

	// 回溯
	cnt := 0
	var dfs func(int, map[int]int8)
	dfs = func(i int, memo map[int]int8) {
		if i == -1 { // 已越界
			cnt++
			return
		}
		dfs(i-1, memo)
		if memo[nums[i]-k] == 0 && memo[nums[i]+k] == 0 {
			memo[nums[i]]++
			dfs(i-1, memo)
			memo[nums[i]]-- // 回溯
		}
	}
	dfs(len(nums)-1, make(map[int]int8, len(nums)))
	return cnt - 1

	// lc：dp
	//groups := map[int]map[int]int{}
	//for _, x := range nums {
	//	if groups[x%k] == nil {
	//		groups[x%k] = map[int]int{}
	//	}
	//	groups[x%k][x]++ // x%k 分为 k 组，并统计相同数字的次数
	//}
	////fmt.Println(groups)
	//ans := 1
	//for _, cnt := range groups { // 计算每组的子集数
	//	m := len(cnt)
	//	type pair struct{ x, c int }
	//	g := make([]pair, 0, m)
	//	for x, c := range cnt {
	//		g = append(g, pair{x, c}) // 值和次数对
	//	}
	//	sort.Slice(g, func(i, j int) bool { return g[i].x < g[j].x }) // 按值排序
	//	f := make([]int, m+1)
	//	f[0] = 1
	//	f[1] = 1 << g[0].c
	//	for i := 1; i < m; i++ { // 每组的子集数
	//		if g[i].x-g[i-1].x == k { // 不选 g[i] / 不选 g[i-1]
	//			f[i+1] = f[i] + f[i-1]*(1<<g[i].c-1)
	//		} else {
	//			f[i+1] = f[i] << g[i].c // g[i] 和 g[i-1] 都可以选择
	//		}
	//	}
	//	//fmt.Println(f)
	//	ans *= f[m] // 统计总子集数
	//}
	//return ans - 1 // -1 去掉空集
}

//leetcode submit region end(Prohibit modification and deletion)
