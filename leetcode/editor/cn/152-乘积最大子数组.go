//给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。
//
// 测试用例的答案是一个 32-位 整数。
//
// 子数组 是数组的连续子序列。
//
//
//
// 示例 1:
//
//
//输入: nums = [2,3,-2,4]
//输出: 6
//解释: 子数组 [2,3] 有最大乘积 6。
//
//
// 示例 2:
//
//
//输入: nums = [-2,0,-1]
//输出: 0
//解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
//
//
//
// 提示:
//
//
// 1 <= nums.length <= 2 * 10⁴
// -10 <= nums[i] <= 10
// nums 的任何前缀或后缀的乘积都 保证 是一个 32-位 整数
//
//
// Related Topics 数组 动态规划 👍 2031 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{3, -1, 4}
	nums = []int{-2, 3, -4}
	nums = []int{2, -5, -2, -4, 3}
	nums = []int{-2}
	nums = []int{2, 3, -2, 4} // 6
	product := maxProduct(nums)
	fmt.Println(product)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxProduct(nums []int) int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func maxProduct(nums []int) int {
//	// 子数组为 [v]：重新开始
//	// minDp：记录最小值（备用于两个负数相乘）
//	ans, maxDp, minDp := math.MinInt32, 1, 1
//	for _, v := range nums {
//		maxV, minV := maxDp*v, minDp*v
//		maxDp, minDp = max(v, maxV, minV), min(v, maxV, minV) // 状态转移方程
//		ans = max(ans, maxDp)
//	}
//	return ans
//}
