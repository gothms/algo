//给你一个长度为 n 、下标从 0 开始的整数数组 nums ，表示收集不同巧克力的成本。每个巧克力都对应一个不同的类型，最初，位于下标 i 的巧克力就对应第
// i 个类型。
//
// 在一步操作中，你可以用成本 x 执行下述行为：
//
//
// 同时修改所有巧克力的类型，将巧克力的类型 iᵗʰ 修改为类型 ((i + 1) mod n)ᵗʰ。
//
//
// 假设你可以执行任意次操作，请返回收集所有类型巧克力所需的最小成本。
//
//
//
// 示例 1：
//
//
//输入：nums = [20,1,15], x = 5
//输出：13
//解释：最开始，巧克力的类型分别是 [0,1,2] 。我们可以用成本 1 购买第 1 个类型的巧克力。
//接着，我们用成本 5 执行一次操作，巧克力的类型变更为 [1,2,0] 。我们可以用成本 1 购买第 2 个类型的巧克力。
//然后，我们用成本 5 执行一次操作，巧克力的类型变更为 [2,0,1] 。我们可以用成本 1 购买第 0 个类型的巧克力。
//因此，收集所有类型的巧克力需要的总成本是 (1 + 5 + 1 + 5 + 1) = 13 。可以证明这是一种最优方案。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,3], x = 4
//输出：6
//解释：我们将会按最初的成本收集全部三个类型的巧克力，而不需执行任何操作。因此，收集所有类型的巧克力需要的总成本是 1 + 2 + 3 = 6 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 10⁹
// 1 <= x <= 10⁹
//
//
// Related Topics 数组 枚举 👍 22 👎 0

package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{15, 150, 56, 69, 214, 203}
	x := 42
	cost := minCost(arr, x)
	fmt.Println(cost)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minCost(nums []int, x int) int64 {
	//minVal := func(a, b int) int {
	//	if b < a {
	//		return b
	//	}
	//	return a
	//}
	//n := len(nums)
	//memo := make([]bool, n)
	//min, cnt := math.MaxInt32, 0
	//var cost int64
	//for _, v := range nums {
	//	min = minVal(min, v)
	//}
	//for i, v := range nums {
	//	if v < min+x {
	//		cost += int64(v)
	//		memo[i] = true
	//		cnt++
	//	}
	//}
	//for _, v := range nums {
	//	cost += int64(minVal(x+min, v))
	//}
	//return cost

	// lc：右移 i 次，则是选择任意区间内 i 个数中最小的数
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	n := len(nums)
	s := make([]int64, n) // s[k] 统计操作 k 次的总成本
	for i := range s {
		s[i] = int64(i) * int64(x)
	}
	for i, mn := range nums { // 子数组左端点
		for j := i; j < n+i; j++ { // 子数组右端点（把数组视作环形的）
			mn = minVal(mn, nums[j%n]) // 维护从 nums[i] 到 nums[j] 的最小值
			s[j-i] += int64(mn)        // 累加操作 j-i 次的花费
			//if j-i == 3 {
			//	fmt.Println(i, j)
			//}
		}
	}
	//ret := int64(math.MaxInt64)
	//for _, v := range s {
	//	if v < ret {
	//		ret = v
	//	}
	//}
	return slices.Min(s)
}

//leetcode submit region end(Prohibit modification and deletion)
