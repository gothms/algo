//给你两个下标从 0 开始的数组 nums1 和 nums2 ，两个数组都只包含非负整数。请你求出另外一个数组 nums3 ，包含 nums1 和 nums2
// 中 所有数对 的异或和（nums1 中每个整数都跟 nums2 中每个整数 恰好 匹配一次）。
//
// 请你返回 nums3 中所有整数的 异或和 。
//
//
//
// 示例 1：
//
// 输入：nums1 = [2,1,3], nums2 = [10,2,5,0]
//输出：13
//解释：
//一个可能的 nums3 数组是 [8,0,7,2,11,3,4,1,9,1,6,3] 。
//所有这些数字的异或和是 13 ，所以我们返回 13 。
//
//
// 示例 2：
//
// 输入：nums1 = [1,2], nums2 = [3,4]
//输出：0
//解释：
//所有数对异或和的结果分别为 nums1[0] ^ nums2[0] ，nums1[0] ^ nums2[1] ，nums1[1] ^ nums2[0] 和
//nums1[1] ^ nums2[1] 。
//所以，一个可能的 nums3 数组是 [2,5,1,6] 。
//2 ^ 5 ^ 1 ^ 6 = 0 ，所以我们返回 0 。
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length, nums2.length <= 10⁵
// 0 <= nums1[i], nums2[j] <= 10⁹
//
//
// Related Topics 位运算 脑筋急转弯 数组 👍 15 👎 0

package main

import "fmt"

func main() {
	nums1 := []int{2, 1, 1}
	nums2 := []int{10, 2, 5, 0}
	nums := xorAllNums(nums1, nums2)
	fmt.Println(nums)

	sum := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum ^= nums1[i] ^ nums2[j]
		}
	}
	fmt.Println(sum)
}

/*
思路：异或
	1.异或运算 a ^ b ^ c = a ^ (b ^ c) = (a ^ b) ^ c
	2.根据这个定理，当 n, m := len(nums1), len(nums2)，所求值 xor =
		n 为奇数：异或 nums2 的所有元素
		m 为奇数：异或 nums1 的所有元素
*/
// leetcode submit region begin(Prohibit modification and deletion)
func xorAllNums(nums1 []int, nums2 []int) int {
	xor, n, m := 0, len(nums1), len(nums2)
	if n&1 > 0 {
		for i := 0; i < m; i++ {
			xor ^= nums2[i]
		}
	}
	if m&1 > 0 {
		for i := 0; i < n; i++ {
			xor ^= nums1[i]
		}
	}
	return xor
}

//leetcode submit region end(Prohibit modification and deletion)
