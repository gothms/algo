//给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
//
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子
//序列。
//
// 示例 1：
//
//
//输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,0,3,2,3]
//输出：4
//
//
// 示例 3：
//
//
//输入：nums = [7,7,7,7,7,7,7]
//输出：1
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2500
// -10⁴ <= nums[i] <= 10⁴
//
//
//
//
// 进阶：
//
//
// 你能将算法的时间复杂度降低到 O(n log(n)) 吗?
//
//
// Related Topics 数组 二分查找 动态规划 👍 3627 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	lis := lengthOfLIS(nums)
	fmt.Println(lis)
}

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLIS(nums []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func lengthOfLIS(nums []int) int {
//	// 二分查找
//	lis := make([]int, 1, len(nums)+1)
//	lis[0] = math.MinInt32 // 哨兵
//	for _, v := range nums {
//		if v > lis[len(lis)-1] { // 追加子序列
//			lis = append(lis, v)
//		} else {
//			lis[sort.SearchInts(lis, v)] = v // 调整子序列
//		}
//	}
//	return len(lis) - 1
//
//	// dp
//	//var ans int
//	//dp := make([]int, len(nums))
//	//for i, v := range nums {
//	//	for j := i - 1; j >= max(0, dp[i]-1); j-- { // max(0, dp[i]-1)：缩小遍历范围
//	//		if v > nums[j] {
//	//			dp[i] = max(dp[i], dp[j]) // 状态转移方程
//	//		}
//	//	}
//	//	dp[i]++ // 自身
//	//	ans = max(ans, dp[i])
//	//}
//	//return ans
//}
