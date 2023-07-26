//给你一个长度为 n 下标从 0 开始的整数数组 nums ，它包含 1 到 n 的所有数字，请你返回上升四元组的数目。
//
// 如果一个四元组 (i, j, k, l) 满足以下条件，我们称它是上升的：
//
//
// 0 <= i < j < k < l < n 且
// nums[i] < nums[k] < nums[j] < nums[l] 。
//
//
//
//
// 示例 1：
//
// 输入：nums = [1,3,2,4,5]
//输出：2
//解释：
//- 当 i = 0 ，j = 1 ，k = 2 且 l = 3 时，有 nums[i] < nums[k] < nums[j] < nums[l] 。
//- 当 i = 0 ，j = 1 ，k = 2 且 l = 4 时，有 nums[i] < nums[k] < nums[j] < nums[l] 。
//没有其他的四元组，所以我们返回 2 。
//
//
// 示例 2：
//
// 输入：nums = [1,2,3,4]
//输出：0
//解释：只存在一个四元组 i = 0 ，j = 1 ，k = 2 ，l = 3 ，但是 nums[j] < nums[k] ，所以我们返回 0 。
//
//
//
//
// 提示：
//
//
// 4 <= nums.length <= 4000
// 1 <= nums[i] <= nums.length
// nums 中所有数字 互不相同 ，nums 是一个排列。
//
//
// Related Topics 树状数组 数组 动态规划 枚举 前缀和 👍 21 👎 0

package main

func main() {

}

// leetcode submit region begin(Prohibit modification and deletion)
func countQuadruplets(nums []int) int64 {
	cnt, n := int64(0), len(nums)
	temp := make([][]int, n)
	temp[n-1] = make([]int, n+1)
	for k := n - 2; k > 0; k-- {
		temp[k] = append([]int(nil), temp[k+1]...) // 累计
		for v := nums[k+1] - 1; v > 0; v-- {
			temp[k][v]++ // k 的右边，比 v 大的元素个数
		}
	}
	for j := 1; j < n-2; j++ {
		for k := j + 1; k < n-1; k++ {
			if nums[j] > nums[k] {
				cnt += int64(temp[k][nums[j]] * (nums[k] - n + j + 1 + temp[j][nums[k]]))
				// nums[i] < nums[k] < nums[j] < nums[l]
				// k 右边比 nums[j] 大的数的个数 * j 左边比 nums[k] 小的数的个数
				// nums[k] - (n-j-1 - temp[j][nums[k]])
			}
		}
	}
	//fmt.Println(temp)
	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
