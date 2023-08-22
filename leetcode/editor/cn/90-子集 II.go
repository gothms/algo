//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//
// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
//
//
//
//
//
//
//
// 示例 1：
//
//
//输入：nums = [1,2,2]
//输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
//
//
// 示例 2：
//
//
//输入：nums = [0]
//输出：[[],[0]]
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
//
//
// Related Topics 位运算 数组 回溯 👍 1126 👎 0

package main

import "sort"

func main() {

}

/*
思路：迭代
	1.在不考虑元素重复的情况，求子集
		类似 LeetCode-78
		所以我们只需要再加上重复元素的判断即可
	2.在排序的情况下，二进制位标记 nums[i] 是否被选择，原则就是
		当 nums[i] 标记被选择，而 nums[j] == nums[j-1] 且 nums[j-1] 被标记不被选择时
		放弃本轮标记位的子集
*/
// leetcode submit region begin(Prohibit modification and deletion)
func subsetsWithDup(nums []int) [][]int {
	// 迭代
	ret, n, m := make([][]int, 1), len(nums), 1<<len(nums)-1
	sort.Ints(nums)	// 有序则更好去重
out:
	for i := 1; i < m; i++ {
		next := make([]int, 0)
		for j := 0; j < n; j++ {
			if i>>j&1 > 0 {	// nums[j] 被选择
				if j > 0 && nums[j] == nums[j-1] && i>>(j-1)&1 == 0 {
					continue out // 去重，放弃本轮标记位的子集
				}
				next = append(next, nums[j])
			}
		}
		ret = append(ret, next)	// 找到一个子集
	}
	ret = append(ret, append([]int(nil), nums...))	// 加尾
	return ret
}
	// 递归
}

//leetcode submit region end(Prohibit modification and deletion)
