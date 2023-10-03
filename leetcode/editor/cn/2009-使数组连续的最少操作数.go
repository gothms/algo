//给你一个整数数组 nums 。每一次操作中，你可以将 nums 中 任意 一个元素替换成 任意 整数。
//
// 如果 nums 满足以下条件，那么它是 连续的 ：
//
//
// nums 中所有元素都是 互不相同 的。
// nums 中 最大 元素与 最小 元素的差等于 nums.length - 1 。
//
//
// 比方说，nums = [4, 2, 5, 3] 是 连续的 ，但是 nums = [1, 2, 3, 5, 6] 不是连续的 。
//
// 请你返回使 nums 连续 的 最少 操作次数。
//
//
//
// 示例 1：
//
// 输入：nums = [4,2,5,3]
//输出：0
//解释：nums 已经是连续的了。
//
//
// 示例 2：
//
// 输入：nums = [1,2,3,5,6]
//输出：1
//解释：一个可能的解是将最后一个元素变为 4 。
//结果数组为 [1,2,3,5,4] ，是连续数组。
//
//
// 示例 3：
//
// 输入：nums = [1,10,100,1000]
//输出：3
//解释：一个可能的解是：
//- 将第二个元素变为 2 。
//- 将第三个元素变为 3 。
//- 将第四个元素变为 4 。
//结果数组为 [1,2,3,4] ，是连续数组。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁹
//
//
// Related Topics 数组 二分查找 👍 32 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 1, 3, 5, 5}
	nums = []int{1, 10, 100, 1000}
	operations := minOperations(nums)
	fmt.Println(operations)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minOperations(nums []int) int {
	minVal := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	n := len(nums)
	sort.Ints(nums) // 排序
	j := 1
	for i := 1; i < n; i++ { // 剔除重复的元素
		if nums[i] == nums[i-1] {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	nums = nums[:j] // 重复的元素，都必须替换
	cnt := j
	for i := 0; i < cnt; i++ {
		c := j - sort.SearchInts(nums, nums[i]+n) + i // i 为左侧需要替换的数量，j - sort.Search 为右侧需要替换的数量
		cnt = minVal(cnt, c)                          // 缩小查询的次数
	}
	return cnt + n - j // cnt + 重复的元素
}

//leetcode submit region end(Prohibit modification and deletion)
