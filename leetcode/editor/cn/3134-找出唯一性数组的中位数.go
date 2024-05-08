//给你一个整数数组 nums 。数组 nums 的 唯一性数组 是一个按元素从小到大排序的数组，包含了 nums 的所有非空子数组中不同元素的个数。
//
// 换句话说，这是由所有 0 <= i <= j < nums.length 的 distinct(nums[i..j]) 组成的递增数组。
//
// 其中，distinct(nums[i..j]) 表示从下标 i 到下标 j 的子数组中不同元素的数量。
//
// 返回 nums 唯一性数组 的 中位数 。
//
// 注意，数组的 中位数 定义为有序数组的中间元素。如果有两个中间元素，则取值较小的那个。
//
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,3]
//
//
// 输出：1
//
// 解释：
//
// nums 的唯一性数组为 [distinct(nums[0..0]), distinct(nums[1..1]), distinct(nums[2..2]
//), distinct(nums[0..1]), distinct(nums[1..2]), distinct(nums[0..2])]，即 [1, 1, 1,
// 2, 2, 3] 。唯一性数组的中位数为 1 ，因此答案是 1 。
//
// 示例 2：
//
//
// 输入：nums = [3,4,3,4,5]
//
//
// 输出：2
//
// 解释：
//
// nums 的唯一性数组为 [1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3] 。唯一性数组的中位数为 2 ，因此答
//案是 2 。
//
// 示例 3：
//
//
// 输入：nums = [4,3,5,4]
//
//
// 输出：2
//
// 解释：
//
// nums 的唯一性数组为 [1, 1, 1, 1, 2, 2, 2, 3, 3, 3] 。唯一性数组的中位数为 2 ，因此答案是 2 。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
// Related Topics 数组 哈希表 二分查找 滑动窗口 👍 6 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	//nums = []int{3, 4, 3, 4, 5}
	//nums = []int{4, 3, 5, 4}
	array := medianOfUniquenessArray(nums)
	fmt.Println(array)
}

// leetcode submit region begin(Prohibit modification and deletion)
func medianOfUniquenessArray(nums []int) int {
	// 二分 + 滑动窗体
	// https://leetcode.cn/problems/find-the-median-of-the-uniqueness-array/solutions/2759114/er-fen-da-an-hua-dong-chuang-kou-pythonj-ykg9/
	n := len(nums)
	k := (n*(n+1)>>1 + 1) >> 1 // 第 k 大
	return sort.Search(n, func(i int) bool {
		cnt, l := 0, 0
		memo := make(map[int]int)
		for r, v := range nums {
			memo[v]++
			for ; len(memo) > i; l++ {
				if memo[nums[l]] == 1 {
					delete(memo, nums[l])
				} else {
					memo[nums[l]]--
				}
			}
			cnt += r - l + 1
			if cnt >= k {
				return true
			}
		}
		return false
	})
}

//leetcode submit region end(Prohibit modification and deletion)
