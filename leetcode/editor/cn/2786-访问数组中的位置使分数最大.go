//给你一个下标从 0 开始的整数数组 nums 和一个正整数 x 。
//
// 你 一开始 在数组的位置 0 处，你可以按照下述规则访问数组中的其他位置：
//
//
// 如果你当前在位置 i ，那么你可以移动到满足 i < j 的 任意 位置 j 。
// 对于你访问的位置 i ，你可以获得分数 nums[i] 。
// 如果你从位置 i 移动到位置 j 且 nums[i] 和 nums[j] 的 奇偶性 不同，那么你将失去分数 x 。
//
//
// 请你返回你能得到的 最大 得分之和。
//
// 注意 ，你一开始的分数为 nums[0] 。
//
//
//
// 示例 1：
//
// 输入：nums = [2,3,6,1,9,2], x = 5
//输出：13
//解释：我们可以按顺序访问数组中的位置：0 -> 2 -> 3 -> 4 。
//对应位置的值为 2 ，6 ，1 和 9 。因为 6 和 1 的奇偶性不同，所以下标从 2 -> 3 让你失去 x = 5 分。
//总得分为：2 + 6 + 1 + 9 - 5 = 13 。
//
//
// 示例 2：
//
// 输入：nums = [2,4,6,8], x = 3
//输出：20
//解释：数组中的所有元素奇偶性都一样，所以我们可以将每个元素都访问一次，而且不会失去任何分数。
//总得分为：2 + 4 + 6 + 8 = 20 。
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// 1 <= nums[i], x <= 10⁶
//
//
// Related Topics 数组 动态规划 👍 6 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{38, 92, 23, 30, 25, 96, 6, 71, 78, 77, 33, 23, 71, 48, 87, 77, 53, 28, 6, 20, 90, 83, 42, 21, 64, 95, 84, 29, 22, 21, 33, 36, 53, 51, 85, 25, 80, 56, 71, 69, 5, 21, 4, 84, 28, 16, 65, 7}
	x := 52 // 1545
	nums = []int{8, 50, 65, 85, 8, 73, 55, 50, 29, 95, 5, 68, 52, 79}
	x = 74 // 470
	nums = []int{18, 13, 60, 61, 57, 21, 10, 98, 51, 3, 13, 36, 72, 70, 68, 62, 52, 83, 63, 63, 53, 42, 59, 98, 95, 48, 22, 64, 94, 80, 14, 14}
	x = 2 // 1433
	score := maxScore(nums, x)
	fmt.Println(score)
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxScore(nums []int, x int) int64 {
	maxVal := func(a, b int64) int64 {
		if b > a {
			return b
		}
		return a
	}
	var a, b int64
	if nums[0]&1 == 0 { // 初始分数为 nums[0]
		a, b = int64(nums[0]), int64(nums[0]-x)
	} else {
		b, a = int64(nums[0]), int64(nums[0]-x)
	}
	for i := 1; i < len(nums); i++ {
		if v := nums[i]; v&1 == 0 {
			a = maxVal(a+int64(v), b+int64(v-x))
		} else {
			b = maxVal(b+int64(v), a+int64(v-x))
		}
	}
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
