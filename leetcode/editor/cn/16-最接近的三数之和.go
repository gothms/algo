//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
// 示例 2：
//
//
//输入：nums = [0,0,0], target = 1
//输出：0
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -10⁴ <= target <= 10⁴
//
//
// Related Topics 数组 双指针 排序 👍 1419 👎 0

package main

import (
	"math"
	"sort"
)

func main() {

}

/*
思路：排序+双指针
	1.思路和 lc-15：三数之和 一样
	2.closest 和 min 分别记录当前最接近 target 的数
		和 closest与target 的差的绝对值
*/
// leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	// 双指针
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	closest, min, n := 0, math.MaxInt32, len(nums)-2
	sort.Ints(nums) // 排序
	for i := 0; i < n; i++ {
		if nums[i]*3-target > min { // 剪枝
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 重复计算
			continue
		}
		for j, k := i+1, n+1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			switch {
			case sum == target: // 没有比相等更接近的
				return target
			case sum > target: // 更新索引
				k--
			default:
				j++
			}
			if curr := abs(sum - target); curr < min {
				closest, min = sum, curr // 更新最接近数和最小绝对差
			}
		}
	}
	return closest
}

//leetcode submit region end(Prohibit modification and deletion)
