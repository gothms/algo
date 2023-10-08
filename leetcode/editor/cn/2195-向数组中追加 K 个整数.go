//给你一个整数数组 nums 和一个整数 k 。请你向 nums 中追加 k 个 未 出现在 nums 中的、互不相同 的 正 整数，并使结果数组的元素和 最
//小 。
//
// 返回追加到 nums 中的 k 个整数之和。
//
//
//
// 示例 1：
//
// 输入：nums = [1,4,25,10,25], k = 2
//输出：5
//解释：在该解法中，向数组中追加的两个互不相同且未出现的正整数是 2 和 3 。
//nums 最终元素和为 1 + 4 + 25 + 10 + 25 + 2 + 3 = 70 ，这是所有情况中的最小值。
//所以追加到数组中的两个整数之和是 2 + 3 = 5 ，所以返回 5 。
//
// 示例 2：
//
// 输入：nums = [5,6], k = 6
//输出：25
//解释：在该解法中，向数组中追加的两个互不相同且未出现的正整数是 1 、2 、3 、4 、7 和 8 。
//nums 最终元素和为 5 + 6 + 1 + 2 + 3 + 4 + 7 + 8 = 36 ，这是所有情况中的最小值。
//所以追加到数组中的两个整数之和是 1 + 2 + 3 + 4 + 7 + 8 = 25 ，所以返回 25 。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i], k <= 10⁹
//
//
// Related Topics 贪心 数组 数学 排序 👍 50 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 34, 65, 12}
	k := 2
	nums = []int{96, 44, 99, 25, 61, 84, 88, 18, 19, 33, 60, 86, 52, 19, 32, 47, 35, 50, 94, 17, 29, 98, 22, 21, 72, 100, 40, 84}
	k = 35 // 794
	kSum := minimalKSum(nums, k)
	fmt.Println(kSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimalKSum(nums []int, k int) int64 {
	// 排序
	sort.Ints(nums) // 排序
	max, sum := k, 0
	for i, v := range nums {
		if i > 0 && v == nums[i-1] { // 排除重复元素
			continue
		}
		if v > max { // 已找完
			break
		}
		max++ // 被占用，则向后多找一个数
		sum += v
	}
	return int64(max*(max+1)>>1 - sum)
}

//leetcode submit region end(Prohibit modification and deletion)
