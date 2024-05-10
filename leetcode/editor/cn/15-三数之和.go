//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
//k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
//
// 你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。
//
//
// 示例 2：
//
//
//输入：nums = [0,1,1]
//输出：[]
//解释：唯一可能的三元组和不为 0 。
//
//
// 示例 3：
//
//
//输入：nums = [0,0,0]
//输出：[[0,0,0]]
//解释：唯一可能的三元组和为 0 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics 数组 双指针 排序 👍 5982 👎 0

package main

import (
	"fmt"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4} // [-4 -1 -1 0 1 2]
	nums = []int{0, 0, 0, 0}           // [-4 -1 -1 0 1 2]
	sum := threeSum(nums)
	fmt.Println(sum)
}

/*
思路：排序+双指针
	1.在非降序的数组中，0 <= i < j < k < n，sum = nums[i]+nums[j]+nums[k]
		sum==0：找到三元组，j++ k--
		sum>0：k--
		sum<0：j++
	2.关键点在于不能有重复三元组
		2.1.当 i > 0 && nums[i] == nums[i-1]，跳过
		2.2.当找到三元组时
			对于索引j：j < k && nums[j] == nums[j-1] 跳过
			对于索引k：j < k && nums[k] == nums[k+1] 跳过
*/

// leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {

}

//leetcode submit region end(Prohibit modification and deletion)

//func threeSum(nums []int) [][]int {
//	// 双指针
//	sort.Ints(nums) // 排序
//	n := len(nums)
//	ans := make([][]int, 0)
//	for i := 0; i < n-2; i++ {
//		if nums[i] > 0 { // 剪枝
//			break
//		}
//		if i > 0 && nums[i] == nums[i-1] { // 避免重复三元组
//			continue
//		}
//		for j, k := i+1, n-1; j < k; {
//			if sum := nums[i] + nums[j] + nums[k]; sum > 0 {
//				k--
//			} else if sum < 0 {
//				j++
//			} else {
//				ans = append(ans, []int{nums[i], nums[j], nums[k]})
//				if nums[j] == nums[k] { // j k 互为哨兵
//					break
//				}
//				j++
//				for nums[j] == nums[j-1] { // 避免重复三元组
//					j++
//				}
//				k--
//				for nums[k] == nums[k+1] {
//					k--
//				}
//			}
//		}
//	}
//	return ans
//}
