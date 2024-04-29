//给你两个整数数组 nums1 和 nums2 ，它们长度都为 n 。
//
// 两个数组的 异或值之和 为 (nums1[0] XOR nums2[0]) + (nums1[1] XOR nums2[1]) + ... + (
//nums1[n - 1] XOR nums2[n - 1]) （下标从 0 开始）。
//
//
// 比方说，[1,2,3] 和 [3,2,1] 的 异或值之和 等于 (1 XOR 3) + (2 XOR 2) + (3 XOR 1) = 2 + 0 +
//2 = 4 。
//
//
// 请你将 nums2 中的元素重新排列，使得 异或值之和 最小 。
//
// 请你返回重新排列之后的 异或值之和 。
//
//
//
// 示例 1：
//
// 输入：nums1 = [1,2], nums2 = [2,3]
//输出：2
//解释：将 nums2 重新排列得到 [3,2] 。
//异或值之和为 (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2 。
//
// 示例 2：
//
// 输入：nums1 = [1,0,3], nums2 = [5,3,4]
//输出：8
//解释：将 nums2 重新排列得到 [5,4,3] 。
//异或值之和为 (1 XOR 5) + (0 XOR 4) + (3 XOR 3) = 4 + 4 + 0 = 8 。
//
//
//
//
// 提示：
//
//
// n == nums1.length
// n == nums2.length
// 1 <= n <= 14
// 0 <= nums1[i], nums2[i] <= 10⁷
//
//
// Related Topics 位运算 数组 动态规划 状态压缩 👍 44 👎 0

package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	nums1 := []int{1, 0, 3}
	nums2 := []int{5, 3, 4}
	xorSum := minimumXORSum(nums1, nums2)
	fmt.Println(xorSum)
}

// leetcode submit region begin(Prohibit modification and deletion)
func minimumXORSum(nums1 []int, nums2 []int) int {
	// 状态压缩
	n := len(nums1)
	m := 1 << n
	dp := make([]int, m)
	for mask := 1; mask < m; mask++ {
		dp[mask] = math.MaxInt32
		cnt := bits.OnesCount(uint(mask)) - 1
		// 写法二
		for i := mask; i > 0; i &= i - 1 { // 从低到高，移除 1
			k := i & -i                                                                 // 取出低位的 1
			dp[mask] = min(dp[mask], nums1[cnt]^nums2[bits.Len(uint(k))-1]+dp[mask&^k]) // ^ 运算优先级等于 +
		}
		// 写法一
		//for i := 0; i < n; i++ {
		//	if mask&(1<<i) > 0 {
		//		dp[mask] = min(dp[mask], nums1[cnt]^nums2[i]+dp[mask^(1<<i)])
		//	}
		//}
	}
	return dp[m-1]
}

//leetcode submit region end(Prohibit modification and deletion)
