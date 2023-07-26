//给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[
//b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
//
//
// 0 <= a, b, c, d < n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// 你可以按 任意顺序 返回答案 。
//
//
//
// 示例 1：
//
//
//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = [2,2,2,2,2], target = 8
//输出：[[2,2,2,2]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
//
// Related Topics 数组 双指针 排序 👍 1631 👎 0

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 2, 2, 2}
	target := 8
	sum := fourSum(nums, target)
	fmt.Println(sum)
}

/*
思路：排序+双指针
	1.n 数之和的思路和4数之和一样，都可以化简为三数之和
	2.所求为 f(n, target)，化简为求 f(n-1, target-nums[i])
		递归，直到 n == 3
		使用三数之和求解 n==3
*/
// leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	ans, n, m := make([][]int, 0), len(nums)-3, len(nums)-2
	sort.Ints(nums)
	threeSum := func(idx int, tar int) { // 三数之和，参考 lc-15
		for i := idx + 1; i < m; i++ {
			if i > idx+1 && nums[i] == nums[i-1] {
				continue
			}
		out:
			for j, k := i+1, m+1; j < k; {
				sum := nums[i] + nums[j] + nums[k]
				switch {
				case sum == tar:
					ans = append(ans, []int{nums[idx], nums[i], nums[j], nums[k]})
					if nums[j] == nums[k] {
						break out
					}
					k--
					for nums[k] == nums[k+1] {
						k--
					}
					j++
					for nums[j] == nums[j-1] {
						j--
					}
				case sum > tar:
					k--
				case sum < tar:
					j++
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if nums[i]<<2 > target { // 剪枝
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 排除重复的解
			continue
		}
		threeSum(i, target-nums[i])
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
