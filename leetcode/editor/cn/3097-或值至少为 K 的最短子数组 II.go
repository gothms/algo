//给你一个 非负 整数数组 nums 和一个整数 k 。
//
// 如果一个数组中所有元素的按位或运算 OR 的值 至少 为 k ，那么我们称这个数组是 特别的 。
//
// 请你返回 nums 中 最短特别非空 子数组的长度，如果特别子数组不存在，那么返回 -1 。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,3], k = 2
//
//
// 输出：1
//
// 解释：
//
// 子数组 [3] 的按位 OR 值为 3 ，所以我们返回 1 。
//
// 示例 2：
//
//
// 输入：nums = [2,1,8], k = 10
//
//
// 输出：3
//
// 解释：
//
// 子数组 [2,1,8] 的按位 OR 值为 11 ，所以我们返回 3 。
//
// 示例 3：
//
//
// 输入：nums = [1,2], k = 0
//
//
// 输出：1
//
// 解释：
//
// 子数组 [1] 的按位 OR 值为 1 ，所以我们返回 1 。
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁵
// 0 <= nums[i] <= 109
// 0 <= k <= 10⁹
//
//
// Related Topics 位运算 数组 滑动窗口 👍 7 👎 0

package main

import (
    "math"
)

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumSubarrayLength(nums []int, k int) int {
	// OR 按位增量：https://www.bilibili.com/video/BV19t421g7Pd/?vd_source=c71b2932bca679afa5372a83fe1143e6
	ors := make([][2]int, 0)
	ret := math.MaxInt
	for i, v := range nums {
		ors = append(ors, [2]int{0, i}) // OR 值，索引
		j := 0
		for idx := range ors {
			ors[idx][0] |= v      // OR 运算
			if ors[idx][0] >= k { // 更新子数组
				ret = min(ret, i-ors[idx][1]+1)
			}
			if ors[j][0] == ors[idx][0] { // or 值不变，则此区间内的 OR 运算没有导致增量
				ors[j][1] = ors[idx][1] // 更新左区间
			} else { // or 值变化
				j++               // 记录导致变化的索引，如 1e9 < 2^30，则最多有 30 个元素索引
				ors[j] = ors[idx] // 记录变化开始的左区间
			}
		}
		ors = ors[:j+1] // 截取掉“费数据”
	}
	if ret == math.MaxInt {
		return -1
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
