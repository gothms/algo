//给你一个由 不同 整数组成的数组 nums ，和一个目标整数 target 。请你从 nums 中找出并返回总和为 target 的元素组合的个数。
//
// 题目数据保证答案符合 32 位整数范围。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,3], target = 4
//输出：7
//解释：
//所有可能的组合为：
//(1, 1, 1, 1)
//(1, 1, 2)
//(1, 2, 1)
//(1, 3)
//(2, 1, 1)
//(2, 2)
//(3, 1)
//请注意，顺序不同的序列被视作不同的组合。
//
//
// 示例 2：
//
//
//输入：nums = [9], target = 3
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// 1 <= nums[i] <= 1000
// nums 中的所有元素 互不相同
// 1 <= target <= 1000
//
//
//
//
// 进阶：如果给定的数组中含有负数会发生什么？问题会产生何种变化？如果允许负数出现，需要向题目中添加哪些限制条件？
//
// Related Topics 数组 动态规划 👍 945 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 3, 6, 6}
	//nums = []int{9}
	target := 21
	sum4 := combinationSum4(nums, target)
	fmt.Println(sum4)
}

/*
排列示例
    12233 排列数 = A(5,5) - A(2,2) - A(2,2)
*/

// leetcode submit region begin(Prohibit modification and deletion)
func combinationSum4(nums []int, target int) int {
	// dp
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, v := range nums {
			if i >= v {
				dp[i] += dp[i-v]
			}
		}
	}
	return dp[target]

	// 记忆化搜索

	// dfs：超时
	//sort.Ints(nums)
	//ret := 0
	//var dfs func(int, int)
	//dfs = func(s, i int) {
	//	if s == 0 {
	//		ret++
	//		return
	//	}
	//	for i >= 0 && nums[i] > s {
	//		i--
	//	}
	//	if i >= 0 {
	//		for _, v := range nums[:i+1] {
	//			dfs(s-v, i)
	//		}
	//	}
	//}
	//dfs(target, len(nums)-1)
	//return ret
}

//leetcode submit region end(Prohibit modification and deletion)
