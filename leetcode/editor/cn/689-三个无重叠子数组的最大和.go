//给你一个整数数组 nums 和一个整数 k ，找出三个长度为 k 、互不重叠、且全部数字和（3 * k 项）最大的子数组，并返回这三个子数组。
//
// 以下标的数组形式返回结果，数组中的每一项分别指示每个子数组的起始位置（下标从 0 开始）。如果有多个结果，返回字典序最小的一个。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,1,2,6,7,5,1], k = 2
//输出：[0,3,5]
//解释：子数组 [1, 2], [2, 6], [7, 5] 对应的起始下标为 [0, 3, 5]。
//也可以取 [2, 1], 但是结果 [1, 3, 5] 在字典序上更大。
//
//
// 示例 2：
//
//
//输入：nums = [1,2,1,2,1,2,1,2,1], k = 2
//输出：[0,2,4]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// 1 <= nums[i] < 2¹⁶
// 1 <= k <= floor(nums.length / 3)
//
//
// Related Topics 数组 动态规划 👍 396 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	// lc：从 1-3 个子数组
	var s1, s2, s3, m1, m2, m3, i1, i2, idx1 int
	ret := make([]int, 3)
	for i := k << 1; i < len(nums); i++ {
		s1 += nums[i-k<<1]
		s2 += nums[i-k]
		s3 += nums[i]
		if i >= k*3-1 {
			if s1 > m1 {
				m1 = s1
				i1 = i - k*3 + 1
			}
			if m1+s2 > m2 {
				m2 = m1 + s2
				i2 = i - k<<1 + 1
				idx1 = i1 // 更新第一个索引
			}
			if m2+s3 > m3 {
				m3 = m2 + s3
				ret[0], ret[1], ret[2] = idx1, i2, i-k+1 // 更新
			}
			s1 -= nums[i-k*3+1] // 滑动窗体
			s2 -= nums[i-k<<1+1]
			s3 -= nums[i-k+1]
		}
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
