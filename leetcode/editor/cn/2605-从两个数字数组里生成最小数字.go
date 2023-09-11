//给你两个只包含 1 到 9 之间数字的数组 nums1 和 nums2 ，每个数组中的元素 互不相同 ，请你返回 最小 的数字，两个数组都 至少 包含这个数
//字的某个数位。
//
//
//
// 示例 1：
//
// 输入：nums1 = [4,1,3], nums2 = [5,7]
//输出：15
//解释：数字 15 的数位 1 在 nums1 中出现，数位 5 在 nums2 中出现。15 是我们能得到的最小数字。
//
//
// 示例 2：
//
// 输入：nums1 = [3,5,2,6], nums2 = [3,1,7]
//输出：3
//解释：数字 3 的数位 3 在两个数组中都出现了。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 9
// 1 <= nums1[i], nums2[i] <= 9
// 每个数组中，元素 互不相同 。
//
//
// Related Topics 数组 哈希表 枚举 👍 16 👎 0

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	for i := 0; i < 9; i++ {
		fmt.Print(1<<i, ", ")
	}
	fmt.Println(bits.TrailingZeros(1))
}

/*
思路：位运算
	1.分别求 nums1 和 nums2 的元素集
		1.1.如果两个元素集有交集
			那么交集中最小的元素，就是所求最小数字
			如 nums1 = [3,5,2,6], nums2 = [3,1,7]，交集为 3
		1.2.如果两个元素没有交集
			取出两个交集各自的最小元素，如 nums1 = [4,1,3], nums2 = [5,7]
			最小元素分别为 1、5，那么 15 就是所求
	2.由题意 1 <= nums1[i], nums2[i] <= 9，为了方便为运算的计算，预定义数组：
		bitN := [10]uint{0, 2, 4, 8, 16, 32, 64, 128, 256, 512}
*/
// leetcode submit region begin(Prohibit modification and deletion)
func minNumber(nums1 []int, nums2 []int) int {
	bitN := [10]uint{0, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	var v1, v2 uint
	for _, v := range nums1 {
		v1 |= bitN[v] // nums1 元素集
	}
	for _, v := range nums2 {
		v2 |= bitN[v] // nums2 元素集
	}
	if mn := v1 & v2; mn > 0 { // 交集
		return bits.TrailingZeros(mn)
	}
	mn := (v1 & -v1) | (v2 & -v2) // 各自的最小元素
	return bits.TrailingZeros(mn)*10 + bits.Len(mn) - 1
}

//leetcode submit region end(Prohibit modification and deletion)
